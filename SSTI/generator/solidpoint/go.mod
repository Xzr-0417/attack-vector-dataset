module payload-generator

go 1.20

require (
	go.uber.org/zap v1.24.0
	gopkg.in/yaml.v3 v3.0.1
)

// 添加以下replace指令指向本地路径
replace dev.solidwall.io/fuchsia/ssti-franziscanner/pkg/engineconfig => /home/xzr/Desktop/ssti-franziscanner/pkg/engineconfig
