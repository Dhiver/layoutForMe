package main

import (
	"flag"
	"os"
)

var (
	Path string
)

func main() {
	InitLogger()

	path := flag.String("path", ".", "Path of the document folder to layout")
	flag.Parse()
	Path = *path + "/"

	err := os.Chdir(Path)
	if err != nil {
		Logger.Fatalf("Can't get into %s : %s", Path, err)
	}
	LoadConf()

	// Load document metadata
	var metadata Metadata
	metadata.Read()

	// Load language translation
	LoadLanguage(metadata.Lang)
	metadata.Interpret()
	metadata.Sanitize()

	// Templating templates
	for _, output := range metadata.Outputs {
		templatingTemplates(output.Template, metadata)
	}

	// Assembling
	content := assembling(metadata)

	// Rendering
	for _, output := range metadata.Outputs {
		rendering(output, content)
	}
}
