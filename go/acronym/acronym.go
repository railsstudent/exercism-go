package acronym

import "strings"

func Abbreviate(s string) string {
	var acm string
	for _, word := range strings.Split(s, " ") {
		for _, token := range strings.Split(word, "-") {
			acm += strings.ToUpper(string(token[0]))
		}
	}
	return acm
}
