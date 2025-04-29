package engineconfig

import (
	"fmt"
	"io/fs"
	"math/rand"
	"strings"
	"text/template"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

const (
	// ExtraChecksAmount is the number of extra checks which must be passed
	// for a vulnerability to be counted.
	// Used for the Render and Blind Exec Sleep techniques
	// for which a single check can result in a false positive.
	ExtraChecksAmount       = 3
	embedFile               = "embed.go"
	firstNumPlaceholder     = "FirstNum"
	secondNumPlaceholder    = "SecondNum"
	formattedNumPlaceholder = "FormattedNum"
	waitTimePlaceholder     = "WaitTime"
	pingDomainPlaceholder   = "PingDomain"
	cmdPlaceholder          = "Cmd"
)

func Parse(rawConfigFiles fs.FS, logger *zap.Logger, baseRespTime int) ([]*ReadyConfig, error) {
	var parsedConfigs []*ReadyConfig
	if err := fs.WalkDir(rawConfigFiles, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if d.Name() == embedFile {
			return nil
		}
		data, err := fs.ReadFile(rawConfigFiles, path)
		if err != nil {
			logger.Error(
				"failed to read file",
				zap.String("filename", path),
				zap.Error(err),
			)
			return nil
		}
		var cfg *EngineConfig
		err = yaml.Unmarshal(data, &cfg)
		if err != nil {
			logger.Error(
				"failed to parse file",
				zap.String("filename", path),
				zap.Error(err),
			)
			return nil
		}

		renderedConfig, err := renderConfig(*cfg, logger, baseRespTime)
		if err != nil {
			return err
		}
		parsedConfigs = append(parsedConfigs, renderedConfig)

		return nil
	}); err != nil {
		return nil, err
	}

	return parsedConfigs, nil
}

// renderConfig inserts all ready values into configs with templates
func renderConfig(rawConfig EngineConfig, logger *zap.Logger, baseRespTime int) (*ReadyConfig, error) {
	readyConfig := &ReadyConfig{rawConfig.Name, []*SSTIWorkflow{}}

	for _, currentCheck := range rawConfig.AttackTechniques {
		switch currentCheck.Name {
		case "Render":
			workflows, err := createRenderPayloads(currentCheck, rawConfig.Contexts)
			if err != nil {
				return nil, err
			}
			readyConfig.Workflows = append(readyConfig.Workflows, workflows...)

		case "Blind Eval Pingback":
			workflows, err := createBlindEvalPayloads(currentCheck, rawConfig.Contexts)
			if err != nil {
				return nil, err
			}
			readyConfig.Workflows = append(readyConfig.Workflows, workflows...)

		case "Exec":
			workflows, err := createExecPayloads(currentCheck, rawConfig.Contexts)
			if err != nil {
				return nil, err
			}
			readyConfig.Workflows = append(readyConfig.Workflows, workflows...)

		case "Blind Exec Sleep":
			workflows, err := createBlindSleepPayloads(currentCheck, rawConfig.Contexts, baseRespTime)
			if err != nil {
				return nil, err
			}
			readyConfig.Workflows = append(readyConfig.Workflows, workflows...)

		case "Blind Exec Pingback":
			workflows, err := createCollaboratorPayloads(currentCheck, rawConfig.Contexts)
			if err != nil {
				return nil, err
			}
			readyConfig.Workflows = append(readyConfig.Workflows, workflows...)
		default:
			logger.Warn(
				"unknown SSTI technique",
				zap.String("technique", currentCheck.Name),
			)
		}
	}

	return readyConfig, nil
}

func createRenderPayloads(technique *Technique, contexts []*Context) ([]*SSTIWorkflow, error) {
	var workflows []*SSTIWorkflow
	var attacks []Attack
	for _, payload := range technique.Payloads {
		for _, ctx := range contexts {

			firstPayload, firstMathResult, err := generateRenderPayload(payload)
			if err != nil {
				return nil, err
			}
			firstPayloadWithCtx := fmt.Sprintf("%s%s%s", ctx.Begin, firstPayload, ctx.End)
			basicPayload := &Payload{
				Payload: firstPayloadWithCtx,
				ExpectedResponse: &ExpectedResponse{
					Body: firstMathResult,
				},
			}

			var extraPayloads []*Payload

			for i := 0; i < ExtraChecksAmount; i++ {
				extraPld, extraMathResult, err := generateRenderPayload(payload)
				if err != nil {
					return nil, err
				}
				extraPldWithCtx := fmt.Sprintf("%s%s%s", ctx.Begin, extraPld, ctx.End)
				extraPayloads = append(extraPayloads, &Payload{
					Payload:          extraPldWithCtx,
					ExpectedResponse: &ExpectedResponse{Body: extraMathResult}})
			}

			attacks = append(attacks, Attack{BasicPayload: basicPayload, ExtraPayloads: extraPayloads})
		}
	}
	workflows = append(workflows, &SSTIWorkflow{
		Attacks:   attacks,
		Technique: technique.Name,
	})

	return workflows, nil
}

func createBlindEvalPayloads(technique *Technique, contexts []*Context) ([]*SSTIWorkflow, error) {
	var workflows []*SSTIWorkflow
	var attacks []Attack
	for _, pld := range technique.Payloads {
		for _, ctx := range contexts {
			pldWithCtx := fmt.Sprintf("%s%s%s", ctx.Begin, pld, ctx.End)
			basicPayload := &Payload{Payload: pldWithCtx, ExpectedResponse: nil}
			attacks = append(attacks, Attack{BasicPayload: basicPayload})
		}
	}
	workflows = append(workflows, &SSTIWorkflow{
		Attacks:   attacks,
		Technique: technique.Name,
	})

	return workflows, nil
}

func createExecPayloads(technique *Technique, contexts []*Context) ([]*SSTIWorkflow, error) {
	var workflows []*SSTIWorkflow
	var attacks []Attack
	for _, payload := range technique.Payloads {
		payloadsWithCmd, err := createCommands(payload, technique.Commands)
		if err != nil {
			return nil, err
		}
		for _, pld := range payloadsWithCmd {
			for _, ctx := range contexts {
				firstPayload, firstResult, err := generateExecPayload(pld)
				if err != nil {
					return nil, err
				}
				pldWithCtx := fmt.Sprintf("%s%s%s", ctx.Begin, firstPayload, ctx.End)
				basicPayload := &Payload{
					Payload: pldWithCtx,
					ExpectedResponse: &ExpectedResponse{
						Body: firstResult,
					},
				}

				var extraPayloads []*Payload

				for i := 0; i < ExtraChecksAmount; i++ {
					extraPld, extraResult, err := generateExecPayload(pld)
					if err != nil {
						return nil, err
					}
					extraPldWithCtx := fmt.Sprintf("%s%s%s", ctx.Begin, extraPld, ctx.End)
					extraPayloads = append(extraPayloads, &Payload{
						Payload: extraPldWithCtx,
						ExpectedResponse: &ExpectedResponse{
							Body: extraResult,
						},
					})
				}

				attacks = append(attacks, Attack{BasicPayload: basicPayload, ExtraPayloads: extraPayloads})
			}
		}
	}
	workflows = append(workflows, &SSTIWorkflow{
		Attacks:   attacks,
		Technique: technique.Name,
	})

	return workflows, nil
}

func createBlindSleepPayloads(technique *Technique, contexts []*Context, baseRespTime int) ([]*SSTIWorkflow, error) {
	var workflows []*SSTIWorkflow
	var attacks []Attack

	waitTime := baseRespTime + 10 + rand.Intn(10)
	if technique.ExpectedResponse != nil && technique.ExpectedResponse.Time != 0 {
		waitTime = technique.ExpectedResponse.Time
	}

	for _, pld := range technique.Payloads {
		for _, ctx := range contexts {
			payload, err := generateBlindSleepPayload(pld, waitTime)
			if err != nil {
				return nil, err
			}
			payloadWithCtx := fmt.Sprintf("%s%s%s", ctx.Begin, payload, ctx.End)
			basicPayload := &Payload{
				Payload: payloadWithCtx,
				ExpectedResponse: &ExpectedResponse{
					Time: waitTime,
				},
			}

			attacks = append(attacks, Attack{BasicPayload: basicPayload})
		}
	}
	workflows = append(workflows, &SSTIWorkflow{
		Attacks:   attacks,
		Technique: technique.Name,
	})

	return workflows, nil
}

func createCollaboratorPayloads(technique *Technique, contexts []*Context) ([]*SSTIWorkflow, error) {
	var workflows []*SSTIWorkflow
	var attacks []Attack
	for _, payload := range technique.Payloads {
		payloadsWithCmd, err := createCommands(payload, technique.Commands)
		if err != nil {
			return nil, err
		}

		for _, pld := range payloadsWithCmd {
			for _, ctx := range contexts {
				pldWithCtx := fmt.Sprintf("%s%s%s", ctx.Begin, pld, ctx.End)
				basicPayload := &Payload{Payload: pldWithCtx, ExpectedResponse: nil}
				attacks = append(attacks, Attack{BasicPayload: basicPayload})
			}
		}
	}
	workflows = append(workflows, &SSTIWorkflow{
		Attacks:   attacks,
		Technique: technique.Name,
	})

	return workflows, nil
}

// createCommands inserts commands from the given slice into the corresponding placeholder inside rawPayload.
func createCommands(rawPayload string, commands []string) ([]string, error) {
	readyPayloads := []string{}
	for _, cmd := range commands {
		data := map[string]string{
			cmdPlaceholder: cmd,
		}

		payload, err := executeTemplate(rawPayload, data)
		if err != nil {
			return nil, err
		}

		readyPayloads = append(readyPayloads, payload)
	}

	return readyPayloads, nil
}

func executeTemplate(rawPayload string, data map[string]string) (string, error) {
	var b strings.Builder
	t, err := template.New("").Parse(rawPayload)
	if err != nil {
		return "", err
	}
	if err := t.Execute(&b, data); err != nil {
		return "", err
	}
	return b.String(), nil
}
