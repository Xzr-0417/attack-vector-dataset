package engineconfig

import (
	"fmt"
	"math/rand"
	"strings"
	"text/template"
	"time"
)

const (
	maxNumValue = 10000
)

var (
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func generateRenderPayload(raw string, ctx Context) (string, error) {
	// 生成随机数并保留种子
	a := rnd.Intn(maxNumValue) + 1
	b := rnd.Intn(maxNumValue) + 1

	tpl, err := template.New("render").Parse(raw)
	if err != nil {
		return "", fmt.Errorf("模板解析失败: %w", err)
	}

	var buf strings.Builder
	data := map[string]interface{}{
		"FirstNum":   a,
		"SecondNum":  b,
		"RandomNum":  a + b,
		"Formatted":  fmt.Sprintf("%dX%d", a, b),
	}

	if err := tpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("模板执行失败: %w", err)
	}

	payload := wrapWithContext(buf.String(), ctx)
	fmt.Printf("[生成器] Render payload: %s\n", payload)
	return payload, nil
}

func generateExecPayload(raw string, ctx Context, commands []string) (string, error) {
	if len(commands) == 0 {
		return "", fmt.Errorf("未提供命令列表")
	}

	selectedCmd := commands[rnd.Intn(len(commands))]
	fmt.Printf("[生成器] 选择命令: %s\n", selectedCmd)

	tpl, err := template.New("exec").Parse(raw)
	if err != nil {
		return "", fmt.Errorf("模板解析失败: %w", err)
	}

	var buf strings.Builder
	data := map[string]interface{}{
		"Cmd":         selectedCmd,
		"Formatted":   fmt.Sprintf("0x%x", time.Now().Unix()),
		"PingDomain":  "your-collaborator.com",
		"WaitTime":    rnd.Intn(15) + 5,
	}

	if err := tpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("模板执行失败: %w", err)
	}

	payload := wrapWithContext(buf.String(), ctx)
	fmt.Printf("[生成器] Exec payload: %s\n", payload)
	return payload, nil
}

func generateRenderWorkflow(tech Technique, contexts []Context) (SSTIWorkflow, error) {
	var attacks []Attack

	for _, payload := range tech.Payloads {
		for _, ctx := range contexts {
			basePayload, err := generateRenderPayload(payload, ctx)
			if err != nil {
				return SSTIWorkflow{}, err
			}

			var extras []Payload
			for i := 0; i < extraChecks; i++ {
				p, err := generateRenderPayload(payload, ctx)
				if err != nil {
					return SSTIWorkflow{}, err
				}
				extras = append(extras, Payload{Payload: p})
			}

			attacks = append(attacks, Attack{
				BasicPayload:   Payload{Payload: basePayload},
				ExtraPayloads: extras,
			})
		}
	}

	return SSTIWorkflow{
		Attacks:   attacks,
		Technique: tech.Name,
	}, nil
}

func generateExecWorkflow(tech Technique, contexts []Context) (SSTIWorkflow, error) {
	var attacks []Attack

	for _, payload := range tech.Payloads {
		for _, ctx := range contexts {
			basePayload, err := generateExecPayload(payload, ctx, tech.Commands)
			if err != nil {
				return SSTIWorkflow{}, err
			}

			var extras []Payload
			for i := 0; i < extraChecks; i++ {
				p, err := generateExecPayload(payload, ctx, tech.Commands)
				if err != nil {
					return SSTIWorkflow{}, err
				}
				extras = append(extras, Payload{Payload: p})
			}

			attacks = append(attacks, Attack{
				BasicPayload:   Payload{Payload: basePayload},
				ExtraPayloads: extras,
			})
		}
	}

	return SSTIWorkflow{
		Attacks:   attacks,
		Technique: tech.Name,
	}, nil
}

func wrapWithContext(content string, ctx Context) string {
	return ctx.Begin + content + ctx.End
}
