package web

import "embed"

var embeddedFiles *embed.FS

// SetEmbeddedFiles enables dependency injection of embedded files.
func SetEmbeddedFiles(efs *embed.FS) {
	embeddedFiles = efs
}
