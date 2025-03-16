package utils

import (
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CapitalizePhrase(phrase string) string {
    caser := cases.Title(language.Und)
    return caser.String(strings.TrimSpace(phrase))
}
func WriteFile(fileName string, content string) error {
    file, err := os.Create(fileName)
    if err != nil {
        return err
    }
    defer file.Close()
    _, err = file.WriteString(content)
    if err != nil {
        return err
    }
    return nil
}
func ReadFile(fileName string) (string, error) {
    content, err := os.ReadFile(fileName)
    if err != nil {
        return "", err
    }
    return string(content), nil
}
