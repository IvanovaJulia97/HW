package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoDetectAndConvert(input string) string {

	if strings.Contains(input, ".") || strings.Contains(input, "-") ||
		strings.Contains(input, " ") {
		return morse.ToText(input)
	}
	return morse.ToMorse(input)
}
