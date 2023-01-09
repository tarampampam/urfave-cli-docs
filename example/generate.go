//go:build docs
// +build docs

package example

import _ "github.com/tarampampam/urfave-cli-docs/markdown"

// Run using `go generate -tags docs ./...`

// Generate CLI usage documentation and write it to the README.md file (between special tags).
//go:generate go run generate_readme.go
