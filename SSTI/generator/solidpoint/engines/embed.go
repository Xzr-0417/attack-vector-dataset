package engines

import (
	"embed"
	"io/fs"
)

//go:embed *.yml
var rawConfigFiles embed.FS

func RawConfigFiles() fs.FS {
	return rawConfigFiles
}
