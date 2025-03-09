package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CapitalizePhrase(phrase string) string {
    caser := cases.Title(language.Und)
    return caser.String(strings.TrimSpace(phrase))
}
