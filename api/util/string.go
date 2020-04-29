package util

import (
	"html"
	"strings"
)

func EscapeHTMLAndTrimString(s string) string {
	return html.EscapeString(strings.TrimSpace(s))
}
