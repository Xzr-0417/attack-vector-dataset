package main

import (
	"fmt"
	"os"
	"strings"
	
	"github.com/Xzr-0417/sstigen/engineconfig"
)

func main() {
	const enginesDir = "engines"

	configs, err := engineconfig.ParseAllEngines(enginesDir)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse engines: %v", err))
	}

	var payloadList []string
	for _, config := range configs {
		for _, workflow := range config.Workflows {
			for _, attack := range workflow.Attacks {
				payloadList = append(payloadList, attack.BasicPayload.Payload)
				for _, extra := range attack.ExtraPayloads {
					payloadList = append(payloadList, extra.Payload)
				}
			}
		}
	}

	if err := os.WriteFile("payload_list.txt", []byte(strings.Join(payloadList, "\n")), 0644); err != nil {
		panic(fmt.Sprintf("Failed to write file: %v", err))
	}

	fmt.Println("Generated", len(payloadList), "payloads")
}
