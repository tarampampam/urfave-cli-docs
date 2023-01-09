# CLI docs generator

![go_version][badge_go_version]
[![tests][badge_tests]][actions]
[![coverage][badge_coverage]][coverage]
[![docs][badge_docs]][docs]

For [urfave](https://github.com/urfave/cli)-based applications.

```bash
$ go get github.com/tarampampam/urfave-cli-docs/markdown@latest
```

## Usage example

Add to your `README.md` file the following lines:

```markdown
<!--GENERATED:CLI_DOCS-->
<!--/GENERATED:CLI_DOCS-->
```

Next, create a file `generate_readme.go`:

```go
//go:build docs
// +build docs

package main

import (
	"os"

	"github.com/tarampampam/urfave-cli-docs/markdown"

	"example"
)

func main() {
	var app = example.NewApp() // <-- your app here

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
```

Then, create a file `generate.go` for generating the documentation:

```go
package example

// Run using `go generate -tags docs ./...`

// Generate CLI usage documentation and write it to the README.md file (between special tags).
//go:generate go run generate_readme.go
```

And run code generation:

```bash
$ go generate -tags docs ./...
```

Viola! Now, open your readme file and watch the result. Example can be found [here](example/readme.md).

[badge_tests]:https://img.shields.io/github/actions/workflow/status/tarampampam/urfave-cli-docs/tests.yml?branch=master
[badge_coverage]:https://img.shields.io/codecov/c/github/tarampampam/urfave-cli-docs/master.svg?maxAge=30
[badge_docs]:https://pkg.go.dev/badge/mod/github.com/tarampampam/urfave-cli-docs
[badge_go_version]:https://img.shields.io/badge/go%20version-%3E=1.16-61CFDD.svg
[actions]:https://github.com/tarampampam/urfave-cli-docs/actions
[coverage]:https://codecov.io/gh/tarampampam/urfave-cli-docs
[docs]:https://pkg.go.dev/github.com/tarampampam/urfave-cli-docs
