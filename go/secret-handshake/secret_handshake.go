package secret

import (
	"fmt"
	"strconv"
	"strings"
)

// Handshake takes an int, converts it to its binary representation and returns a sequence of action strings
func Handshake(code uint) []string {
	m := make(map[uint]string)
	m[0] = ""
	m[1] = "wink"
	m[2] = "double wink"
	m[4] = "close your eyes"
	m[8] = "jump"
	m[16] = ""

	reverseOrder := false
	if code >= 16 {
		reverseOrder = true
	}

	strBin := strconv.FormatUint(uint64(code), 2)
	numPrepadZeroes := 5 - len(strBin)
	// Prepend 0 to make it 5 digit long
	if numPrepadZeroes > 0 {
		strBin = strings.Repeat("0", numPrepadZeroes) + strBin
	}

	fmt.Println("---", strBin)
	if reverseOrder {

	} else {

	}
	return []string{""}
}

//31 = 11111
