package util

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"regexp"
	"strings"
)

func ToPascalCase(s string) string {
	str := strings.ToLower(s)

	re := regexp.MustCompile(`([[:alnum:]]+)`)

	words := re.FindAllString(str, -1)

	for i, word := range words {
		words[i] = cases.Title(language.English).String(word)
	}

	return strings.Join(words, "")
}

func NullableString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
