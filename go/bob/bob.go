package bob

import "strings"
import "unicode"

func isYell(remark string) bool {
	testLowercase := func(c rune) bool {
		return unicode.IsLower(c)
	}
	testUppercase := func(c rune) bool {
		return unicode.IsUpper(c)
	}
	if strings.IndexFunc(remark, testLowercase) >= 0 {
		return false
	}

	if strings.IndexFunc(remark, testUppercase) >= 0 {
		return true
	}
	return false
}

func Hey(remark string) string {
	response := "Whatever."
	trimmedRemark := strings.TrimSpace(remark)

	if isYell(trimmedRemark) {
		response = "Whoa, chill out!"
	} else if len(trimmedRemark) == 0 {
		response = "Fine. Be that way!"
	} else if strings.HasSuffix(trimmedRemark, "?") {
		response = "Sure."
	}
	return response
}
