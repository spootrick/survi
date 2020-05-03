package util

import (
	"html"
	"strings"
)

func EscapeHTMLAndTrimString(s string) string {
	return html.EscapeString(strings.TrimSpace(s))
}

func EscapeHTMLAndTrimStringPtr(s *string) *string {
	str := html.EscapeString(strings.TrimSpace(*s))

	return &str
}
