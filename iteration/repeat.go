package iteration

import (
	"strings"
)

func Repeat(letter string, repeatCount int) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(letter)
	}
	return repeated.String()
}
