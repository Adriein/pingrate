package utils

import (
	"github.com/rotisserie/eris"
	"regexp"
	"unicode"
)

func CamelToSnake(str string) string {
	var result []rune

	for i, char := range str {
		if unicode.IsUpper(char) && i > 0 {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(char))
	}

	return string(result)
}

func SnakeToCamel(str string) string {
	var result []rune
	capitalizeNext := false
	for _, char := range str {
		if char == '_' {
			capitalizeNext = true
			continue
		}
		if capitalizeNext {
			result = append(result, unicode.ToUpper(char))
			capitalizeNext = false
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}

func ExtractJSON(input string) (string, error) {
	re := regexp.MustCompile(`(\{.*\})`)

	matches := re.FindStringSubmatch(input)

	if len(matches) < 2 {
		return "", eris.New("no JSON found")
	}

	return matches[1], nil
}
