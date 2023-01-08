package markdown

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/russross/blackfriday"
	"github.com/shurcooL/markdownfmt/markdown"
	"github.com/urfave/cli/v2"

	cliDocs "github.com/tarampampam/urfave-cli-docs"
)

type Markdown struct {
	appPath string
	tmpl    string
}

type MarkdownOption func(*Markdown)

func WithTemplate(tmpl string) MarkdownOption { return func(m *Markdown) { m.tmpl = tmpl } }
func WithAppPath(path string) MarkdownOption  { return func(m *Markdown) { m.appPath = path } }

var (
	//go:embed markdown.tmpl
	defaultMarkdownTemplate string
)

func ToMarkdown(app *cli.App, opt ...MarkdownOption) (string, error) {
	var md = Markdown{ // defaults
		appPath: "./app",
		tmpl:    defaultMarkdownTemplate,
	}

	for _, option := range opt {
		option(&md)
	}

	t, err := template.New("cli").Funcs(template.FuncMap{
		"join": strings.Join,
	}).Parse(md.tmpl)
	if err != nil {
		return "", err
	}

	var w bytes.Buffer

	if err = t.ExecuteTemplate(&w, "cli", cliDocs.NewApp(app, md.appPath)); err != nil {
		return "", err
	}

	// extensions for GitHub Flavored Markdown-like parsing.
	const extensions = blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_NO_EMPTY_LINE_BEFORE_BLOCK

	return string(blackfriday.Markdown(w.Bytes(), markdown.NewRenderer(&markdown.Options{}), extensions)), nil
}
