package tournament

import "io"
import "errors"

var OUTCOME = []string { "win", "draw", "draw" }
var POINTS = []int { 3, 1, 0 }

func Tally(reader io.Reader, writer io.Writer) error {
  gamePlayed := 0
  return errors.New("ddd")
}
