package main

import (
	"embed"
	"os"

	"github.com/egustafson/websb-go/cmd"
	"github.com/egustafson/websb-go/web"
)

//go:embed static templates
var embeddedFiles embed.FS

var (
	// GitSummary = git describe --tags --dirty --always
	GitSummary = "v0.0.0-dirty"
	// BuildDate = date -u +"%Y-%m-%dT%H:%M:%SZ"
	BuildDate = "1970-01-01T00:00:00Z"
)

func init() {
	web.SetEmbeddedFiles(&embeddedFiles)
}

func main() {
	err := cmd.Execute(GitSummary, BuildDate)
	if err != nil {
		// cobra will print an error to stdout
		os.Exit(1)
	}
}
