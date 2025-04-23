package engineconfig

import (
	"fmt"
	"math/rand"
	"strings"
	"text/template"
)

func generateRenderPayload(raw string, ctx Context) (string, error) {
	a := rand.Intn(1000) + 5
	b := rand.Intn(1000) + 5

	tpl, err := template.New("").Parse(raw)
	if err != nil {
		return "", err
	}

	var buf strings.Builder
	err = tpl.Execute(&buf, map[string]interface{}{
		"FirstNum":  a,
		"SecondNum": b,
	})
	if err != nil {
		return "", err
	}

	return ctx.Begin + buf.String() + ctx.End, nil
}

func generateExecPayload(raw string, ctx Context, commands []string) (string, error) {
	if len(commands) == 0 {
		return "", fmt.Errorf("no commands provided")
	}

	selectedCmd := commands[rand.Intn(len(commands))]
	
	tpl, err := template.New("").Parse(raw)
	if err != nil {
		return "", err
	}

	var buf strings.Builder
	err = tpl.Execute(&buf, map[string]interface{}{
		"Command": selectedCmd,
	})
	if err != nil {
		return "", err
	}

	return ctx.Begin + buf.String() + ctx.End, nil
}

func createRenderPayloads(tech Technique, contexts []Context) (SSTIWorkflow, error) {
	var attacks []Attack

	for _, payload := range tech.Payloads {
		for _, ctx := range contexts {
			basic, err := generateRenderPayload(payload, ctx)
			if err != nil {
				return SSTIWorkflow{}, err
			}

			var extras []Payload
			for i := 0; i < 3; i++ {
				p, err := generateRenderPayload(payload, ctx)
				if err != nil {
					return SSTIWorkflow{}, err
				}
				extras = append(extras, Payload{Payload: p})
			}

			attacks = append(attacks, Attack{
				BasicPayload:   Payload{Payload: basic},
				ExtraPayloads: extras,
			})
		}
	}

	return SSTIWorkflow{
		Attacks:   attacks,
		Technique: tech.Name,
	}, nil
}
