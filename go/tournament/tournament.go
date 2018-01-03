package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type teamRecord struct {
	score int
	win   int
  draw  int
	loss  int
}

// https://medium.com/@matryer/golang-advent-calendar-day-seventeen-io-reader-in-depth-6f744bb4320b

func Tally(reader io.Reader, writer io.Writer) error {
	standing := make(map[string]*teamRecord)

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	rows := strings.Split(fmt.Sprintf("%s", b), "\n")
	for _, row := range rows {
		fmt.Printf("%s\n", row)
		record := strings.Split(row, ";")
		//fmt.Printf("%d\n", len(record))
		if len(record) == 3 {
			firstTeamName := record[0]
			secondTeamName := record[1]
			outcome := record[2]
			if outcome != "win" && outcome != "loss" && outcome != "draw" {
				return errors.New(fmt.Sprintf("bad input: %s", outcome))
			}

			if _, ok := standing[firstTeamName]; !ok {
				standing[firstTeamName] = &teamRecord{score: 0, win: 0, draw: 0, loss: 0}
			}
			if _, ok2 := standing[secondTeamName]; !ok2 {
				standing[secondTeamName] = &teamRecord{score: 0, win: 0, draw: 0, loss: 0}
			}

      firstTeam := standing[firstTeamName]
      secondTeam := standing[secondTeamName]
      switch outcome {
			case "win":
	      firstTeam.win += 1
        firstTeam.score += 3
        secondTeam.loss += 1
        fmt.Println(firstTeamName, firstTeam)
        fmt.Println(secondTeamName, secondTeam)
			case "draw":
        firstTeam.draw++
        firstTeam.score += 1
        secondTeam.draw++
        secondTeam.score += 1
        fmt.Println(firstTeamName, firstTeam)
        fmt.Println(secondTeamName, secondTeam)

			case "loss":
        firstTeam.loss++
        secondTeam.win++
        secondTeam.score += 3
        fmt.Println(firstTeamName, firstTeam)
        fmt.Println(secondTeamName, secondTeam)
			}
		}
	}
	output := fmt.Sprintf("%-31s|%4s|%4s|%4s|%4s|%4s\n", "TEAM", " MP ", "  W " , "  D ", "  L ", "  P ")
	for k, v := range standing {
    output += fmt.Sprintf("%-31s|  %d |  %d |  %d |  %d |  %d\n", k, (v.win+v.loss+v.draw), v.win, v.draw, v.loss, v.score)
  }
	fmt.Printf("%s", output)
	fmt.Fprint(writer, output)

	//return errors.New("ddd")
	return nil
}
