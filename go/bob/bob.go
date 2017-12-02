package bob

import "strings"
import "unicode"

func isYell(remark string) bool {
	if strings.ToUpper(remark) == remark {
		for _, l := range remark {
			if unicode.IsLetter(l) {
				return true
			}
		}
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
