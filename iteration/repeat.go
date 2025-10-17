package iteration

import "strings"

const repeatCount = 5

func Repeat(a string) string {

	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(a)
	}
	return repeated.String()
}
