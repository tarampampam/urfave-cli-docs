//go:build docs
// +build docs

package main

import (
	"os"

	"github.com/tarampampam/urfave-cli-docs/markdown"

	"example"
)

func main() {
	var app = example.NewApp()

	// generate markdown documentation for the app
	docs, err := markdown.Render(app)
	if err != nil {
		panic(err)
	}

	const readmeFilePath = "./readme.md"

	// read readme file
	readme, err := os.ReadFile(readmeFilePath)
	if err != nil {
		panic(err)
	}

	const start, end = "<!--GENERATED:CLI_DOCS-->", "<!--/GENERATED:CLI_DOCS-->"

	// replace the documentation section in the readme file
	updated, err := markdown.ReplaceBetween(start, end, string(readme), docs)
	if err != nil {
		panic(err)
	}

	// write the updated readme file
	if err = os.WriteFile(readmeFilePath, []byte(updated), 0664); err != nil {
		panic(err)
	}
}
