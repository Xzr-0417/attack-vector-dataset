package engineconfig

import (
	"fmt"
	"os"
	"path/filepath"
	"gopkg.in/yaml.v3"
)

const (
	debugMode   = true
	extraChecks = 3
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

	log("开始扫描引擎目录: %s", dir)
	
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("目录遍历错误: %w", err)
		}

		// 跳过非YAML文件和目录
		if info.IsDir() || !isYamlFile(path) {
			return nil
		}

		log("找到配置文件: %s", path)

		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("文件读取失败: %w", err)
		}

		var cfg EngineConfig
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			return fmt.Errorf("YAML解析失败: %w", err)
		}
		log("成功解析 %s 配置", cfg.Name)

		// 校验必要字段
		if cfg.Name == "" {
			return fmt.Errorf("缺少引擎名称(name字段)")
		}
		if len(cfg.Contexts) == 0 {
			return fmt.Errorf("%s 缺少上下文(contexts)", cfg.Name)
		}

		rendered, err := renderEngineConfig(cfg)
		if err != nil {
			return fmt.Errorf("配置渲染失败: %w", err)
		}

		configs = append(configs, rendered)
		return nil
	})

	return configs, err
}

func renderEngineConfig(raw EngineConfig) (*EngineConfig, error) {
	engine := &EngineConfig{
		Name:     raw.Name,
		Contexts: raw.Contexts,
	}

	if len(raw.AttackTechniques) == 0 {
		log("警告: %s 没有配置攻击技术", raw.Name)
		return engine, nil
	}

	for _, tech := range raw.AttackTechniques {
		log("处理攻击技术: %s", tech.Name)
		
		var workflow SSTIWorkflow
		var err error

		switch tech.Name {
		case "Render":
			workflow, err = generateRenderWorkflow(tech, raw.Contexts)
		case "Exec":
			workflow, err = generateExecWorkflow(tech, raw.Contexts)
		default:
			log("跳过未支持的技术类型: %s", tech.Name)
			continue
		}

		if err != nil {
			return nil, err
		}
		engine.Workflows = append(engine.Workflows, workflow)
	}

	log("完成 %s 的配置渲染，生成 %d 个工作流", 
		raw.Name, 
		len(engine.Workflows))
	
	return engine, nil
}

// 辅助函数
func isYamlFile(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".yaml" || ext == ".yml"
}

func log(format string, args ...interface{}) {
	if debugMode {
		fmt.Printf("[Parser] "+format+"\n", args...)
	}
}
