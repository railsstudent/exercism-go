package twelve

import (
	"fmt"
	"strings"
)

// Song returns the verse of twelfth days of Christmas
func Song() string {
	song := ""
	for i := 1; i <= 12; i++ {
		song += fmt.Sprintln(Verse(i))
	}
	return song
}

// Verse returns the verse of day days of Christmas
func Verse(day int) string {
	strDayMap := make(map[int]string)
	strDayMap[1] = "first"
	strDayMap[2] = "second"
	strDayMap[3] = "third"
	strDayMap[4] = "fourth"
	strDayMap[5] = "fifth"
	strDayMap[6] = "sixth"
	strDayMap[7] = "seventh"
	strDayMap[8] = "eighth"
	strDayMap[9] = "ninth"
	strDayMap[10] = "tenth"
	strDayMap[11] = "eleventh"
	strDayMap[12] = "twelfth"

	presents := [12]string{"twelve Drummers Drumming", "eleven Pipers Piping", "ten Lords-a-Leaping",
		"nine Ladies Dancing", "eight Maids-a-Milking", "seven Swans-a-Swimming",
		"six Geese-a-Laying", "five Gold Rings", "four Calling Birds",
		"three French Hens", "two Turtle Doves", "a Partridge in a Pear Tree"}

	var present string
	if day > 1 {
		present = strings.Join(presents[12-day:11], ", ") + ", and " + presents[11]
	} else {
		present = presents[12-day]
	}
	template := "On the %s day of Christmas my true love gave to me, %s."
	return fmt.Sprintf(template, strDayMap[day], present)
}
