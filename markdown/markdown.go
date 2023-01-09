package markdown

import (
	"bytes"
	_ "embed"
	"regexp"
	"strings"
	"text/template"

	"github.com/Kunde21/markdownfmt/v3/markdown"
	"github.com/urfave/cli/v2"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"

	cliDocs "gh.tarampamp.am/urfave-cli-docs"
)

// Markdown is a markdown formatter.
type Markdown struct {
	appPath string
	tmpl    string
}

// Option is a markdown formatter option.
type Option func(*Markdown)

// WithTemplate allows to override the default template.
func WithTemplate(tmpl string) Option { return func(m *Markdown) { m.tmpl = tmpl } }

// WithAppPath allows to override the default app path.
func WithAppPath(path string) Option { return func(m *Markdown) { m.appPath = path } }

//go:embed markdown.tmpl
var defaultMarkdownTemplate string

// Render renders the markdown documentation for the given cli app.
func Render(app *cli.App, opt ...Option) (string, error) {
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

	formatted, err := FormatMarkdown(w.Bytes())
	if err != nil {
		return "", err
	}

	return strings.TrimRight(string(formatted), "\n"), nil
}

// ReplaceBetween replaces the text between the given start and end markers (tags).
func ReplaceBetween(start, end string, where, with string) (string, error) {
	re, err := regexp.Compile("(?s)" + regexp.QuoteMeta(start) + "(.*?)" + regexp.QuoteMeta(end))
	if err != nil {
		return "", err
	}

	return re.ReplaceAllString(where, strings.Join([]string{start, with, end}, "\n")), nil
}

// FormatMarkdown formats given Markdown content (like a `go fmt`, but for markdown).
func FormatMarkdown(in []byte) ([]byte, error) {
	var (
		out = bytes.NewBuffer(nil)
		gm  = goldmark.New(
			goldmark.WithExtensions(extension.GFM),
			goldmark.WithParserOptions(parser.WithAttribute()), // We need this to enable # headers {#custom-ids}.
			goldmark.WithRenderer(markdown.NewRenderer()),
		)
	)

	if err := gm.Convert(in, out); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
