package urfave_cli_docs

import (
	"regexp"
	"strings"
)

// ReplaceBetween replaces the text between the given start and end markers (tags).
func ReplaceBetween(start, end string, where, with string) (string, error) {
	re, err := regexp.Compile("(?s)" + regexp.QuoteMeta(start) + "(.*?)" + regexp.QuoteMeta(end))
	if err != nil {
		return "", err
	}

	return re.ReplaceAllString(where, strings.Join([]string{start, with, end}, "\n")), nil
}
