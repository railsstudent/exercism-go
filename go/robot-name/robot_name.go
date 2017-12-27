package robotname

import "math/rand"
import "time"

type Robot struct {
	name string
}

var seed *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

const ALPHABETS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const DIGITS = "0123456789"
const lenAlphabets = len(ALPHABETS)
const lenDigits = len(DIGITS)

var used = make(map[string]bool)

func (robot *Robot) Name() string {
	if robot.name == "" {
		unique := false
		for !unique {
			tmpName := ""
			for i := 0; i < 2; i++ {
				tmpName += string(ALPHABETS[seed.Int()%lenAlphabets])
			}

			for i := 0; i < 3; i++ {
				tmpName += string(DIGITS[seed.Int()%lenDigits])
			}
			if _, ok := used[tmpName]; !ok {
				used[tmpName] = true
				unique = true
				robot.name = tmpName
			}
		}
	}
	return robot.name
}

func (robot *Robot) Reset() {
	robot.name = ""
}
