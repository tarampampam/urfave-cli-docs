package urfave_cli_docs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	urfaveCliDocs "gh.tarampamp.am/urfave-cli-docs"
)

func TestPrepareMultilineString(t *testing.T) {
	assert.Equal(t, "foo", urfaveCliDocs.PrepareMultilineString("foo"))
	assert.Equal(t, "foo", urfaveCliDocs.PrepareMultilineString(" \nfoo\n\t"))
	assert.Equal(t, "foo \t bar", urfaveCliDocs.PrepareMultilineString(" \nfoo\n\t bar.\n\n"))
}

func TestPrepareFlags(t *testing.T) {
	t.Skip("TODO")
}

func TestPrepareCommands(t *testing.T) {
	t.Skip("TODO")
}

func TestNewApp(t *testing.T) {
	t.Skip("TODO")
}
