package main

import (
	"fmt"
	"os"
	"path/filepath"
	"embed"
	"log"
	"go.uber.org/zap"
	"dev.solidwall.io/fuchsia/ssti-franziscanner/pkg/engineconfig"
)

//go:embed engines/*.yml
var engineFiles embed.FS

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	// 1. 解析所有引擎配置
	configs, err := engineconfig.Parse(engineFiles, logger, 2000)
	if err != nil {
		log.Fatal("解析配置失败:", err)
	}

	// 2. 生成Payload集合
	payloadSet := make(map[string]struct{})
	for _, cfg := range configs {
		for _, workflow := range cfg.Workflows {
			for _, attack := range workflow.Attacks {
				// 基础Payload
				if attack.BasicPayload != nil {
					payloadSet[attack.BasicPayload.Payload] = struct{}{}
				}
				// 额外验证Payload
				for _, extra := range attack.ExtraPayloads {
					payloadSet[extra.Payload] = struct{}{}
				}
			}
		}
	}

	// 3. 写入文件
	outputFile := "payloadlist.txt"
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("创建文件失败:", err)
	}
	defer f.Close()

	for p := range payloadSet {
		if _, err := f.WriteString(p + "\n"); err != nil {
			log.Println("写入失败:", err)
		}
	}

	fmt.Printf("生成完成！共生成 %d 个Payload\n", len(payloadSet))
}
