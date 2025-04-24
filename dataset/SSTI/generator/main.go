package main

import (
	"fmt"
	"os"
	"strings"
	
	"github.com/Xzr-0417/attack-vector-dataset/dataset/SSTI/generator/engineconfig"
)

func main() {
	const enginesDir = "engines"
	fmt.Println("[Main] 初始化SSTI Payload生成器")
	fmt.Printf("[Main] 扫描引擎目录: %s\n", enginesDir)

	configs, err := engineconfig.ParseAllEngines(enginesDir)
	if err != nil {
		panic(fmt.Sprintf("[Main] 配置解析失败: %v", err))
	}
	fmt.Printf("[Main] 发现 %d 个模板引擎配置\n", len(configs))

	var payloadList []string
	totalTechniques := 0
	
	for _, config := range configs {
		fmt.Printf("[Engine] 处理引擎: %s\n", config.Name)
		for _, workflow := range config.Workflows {
			totalTechniques++
			fmt.Printf("  |- 技术类型: %-18s 发现 %d 个攻击向量\n", 
				workflow.Technique, 
				len(workflow.Attacks)*4) // 基础payload + 3个额外payload
			
			for _, attack := range workflow.Attacks {
				payloadList = append(payloadList, attack.BasicPayload.Payload)
				for _, extra := range attack.ExtraPayloads {
					payloadList = append(payloadList, extra.Payload)
				}
			}
		}
	}

	if len(payloadList) == 0 {
		fmt.Println("[Warning] 没有生成任何payload，请检查：")
		fmt.Println("1. engines目录是否存在YAML文件")
		fmt.Println("2. YAML文件语法是否正确")
		fmt.Println("3. 是否包含支持的攻击技术（Render/Exec）")
	}

	if err := os.WriteFile("payload_list.txt", []byte(strings.Join(payloadList, "\n")), 0644); err != nil {
		panic(fmt.Sprintf("[Main] 文件写入失败: %v", err))
	}

	fmt.Printf("\n[Result] 成功生成 %d 个payload（来自 %d 种技术）\n", 
		len(payloadList), 
		totalTechniques)
}
