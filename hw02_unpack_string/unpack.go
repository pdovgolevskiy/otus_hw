package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

// Проверить, что в строке нет повторяющихся цифр, цифра не первая и строка не пустая.
func isValid(runeStr []rune) bool {
	if len(runeStr) == 0 {
		return false
	}
	if unicode.IsDigit(runeStr[0]) {
		return false
	}
	for i := 1; i < len(runeStr)-1; i++ {
		if unicode.IsDigit(runeStr[i]) && unicode.IsDigit(runeStr[i+1]) {
			return false
		}
	}
	return true
}

func Unpack(inStr string) (string, error) {
	if inStr == "" {
		return "", nil
	}
	runeStr := []rune(inStr)
	if !isValid(runeStr) {
		return "", ErrInvalidString
	}
	var sb strings.Builder
	for i := 0; i < len(runeStr); i++ {
		if i == (len(runeStr) - 1) {
			sb.WriteRune(runeStr[i])
			break
		}
		// Если след. цифра, то добавить повторяющуюся строку и пропустить эту цифру в цикле.
		if unicode.IsDigit(runeStr[i+1]) {
			sb.WriteString(strings.Repeat(string(runeStr[i]), int(runeStr[i+1])-'0'))
			i++
			continue
		}
		sb.WriteRune(runeStr[i])
	}
	return sb.String(), nil
}
