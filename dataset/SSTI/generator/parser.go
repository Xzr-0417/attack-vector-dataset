package engineconfig

import (
	"fmt"
	"os"
	"path/filepath"
	
	"gopkg.in/yaml.v3"
)

type EngineConfig struct {
	Name            string      `yaml:"name"`
	Contexts        []Context   `yaml:"contexts"`
	AttackTechniques []Technique `yaml:"attackTechniques"`
	Workflows       []SSTIWorkflow
}

type Context struct {
	Begin string `yaml:"begin"`
	End   string `yaml:"end"`
}

type Technique struct {
	Name     string   `yaml:"name"`
	Payloads []string `yaml:"payloads"`
	Commands []string `yaml:"commands,omitempty"`
}

type SSTIWorkflow struct {
	Attacks   []Attack
	Technique string
}

type Attack struct {
	BasicPayload   Payload
	ExtraPayloads []Payload
}

type Payload struct {
	Payload string
}

func ParseAllEngines(dir string) ([]*EngineConfig, error) {
	var configs []*EngineConfig

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || filepath.Ext(path) != ".yaml" {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", path, err)
		}

		var cfg EngineConfig
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			return fmt.Errorf("failed to parse %s: %w", path, err)
		}

		rendered, err := renderConfig(cfg)
		if err != nil {
			return fmt.Errorf("failed to render %s: %w", path, err)
		}

		configs = append(configs, rendered)
		return nil
	})

	return configs, err
}

func renderConfig(raw EngineConfig) (*EngineConfig, error) {
	rendered := &EngineConfig{
		Name:     raw.Name,
		Contexts: raw.Contexts,
	}

	for _, tech := range raw.AttackTechniques {
		var workflow SSTIWorkflow
		var err error

		switch tech.Name {
		case "Render":
			workflow, err = createRenderPayloads(tech, raw.Contexts)
		case "Exec":
			workflow, err = createExecPayloads(tech, raw.Contexts)
		default:
			continue
		}

		if err != nil {
			return nil, err
		}
		rendered.Workflows = append(rendered.Workflows, workflow)
	}

	return rendered, nil
}

func getPayloadsFromWorkflows(workflows []SSTIWorkflow) []string {
	var payloads []string
	for _, wf := range workflows {
		for _, attack := range wf.Attacks {
			payloads = append(payloads, attack.BasicPayload.Payload)
			for _, extra := range attack.ExtraPayloads {
				payloads = append(payloads, extra.Payload)
			}
		}
	}
	return payloads
}

func createExecPayloads(tech Technique, contexts []Context) (SSTIWorkflow, error) {
	var attacks []Attack

	for _, payload := range tech.Payloads {
		for _, ctx := range contexts {
			basic, err := generateExecPayload(payload, ctx, tech.Commands)
			if err != nil {
				return SSTIWorkflow{}, err
			}

			var extras []Payload
			for i := 0; i < 3; i++ {
				p, err := generateExecPayload(payload, ctx, tech.Commands)
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
