module example

go 1.19

require (
	github.com/tarampampam/urfave-cli-docs v0.0.0-00010101000000-000000000000
	github.com/tarampampam/urfave-cli-docs/markdown v0.0.0-00010101000000-000000000000
	github.com/urfave/cli/v2 v2.23.7
)

require (
	github.com/Kunde21/markdownfmt/v3 v3.1.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	github.com/yuin/goldmark v1.3.5 // indirect
)

replace (
	github.com/tarampampam/urfave-cli-docs => ../
	github.com/tarampampam/urfave-cli-docs/markdown => ../markdown
)
