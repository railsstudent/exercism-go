package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type teamRecord struct {
	name  string
	score int
	win   int
	draw  int
	loss  int
}

type byScoreTeamName []teamRecord

func (t byScoreTeamName) Len() int {
	return len(t)
}

func (t byScoreTeamName) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t byScoreTeamName) Less(i, j int) bool {
	// sort by score in descending order and then team name in ascending order
	if t[i].score < t[j].score {
		return false
	}

	if t[i].score > t[j].score {
		return true
	}

	// same score, sort by name
	return t[i].name < t[j].name
}

func Tally(reader io.Reader, writer io.Writer) error {
	standing := make(map[string]*teamRecord)

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	rows := strings.Split(fmt.Sprintf("%s", b), "\n")
	for _, row := range rows {
		if row == "" || row[0] == byte('#') {
			continue
		}
		record := strings.Split(row, ";")
		if len(record) != 3 {
			return errors.New("Incorrect number of arguments")
		}

		firstTeamName := record[0]
		secondTeamName := record[1]
		outcome := record[2]
		if outcome != "win" && outcome != "loss" && outcome != "draw" {
			return errors.New("Bad input")
		}

		if _, ok := standing[firstTeamName]; !ok {
			standing[firstTeamName] = &teamRecord{name: firstTeamName, score: 0, win: 0, draw: 0, loss: 0}
		}
		if _, ok2 := standing[secondTeamName]; !ok2 {
			standing[secondTeamName] = &teamRecord{name: secondTeamName, score: 0, win: 0, draw: 0, loss: 0}
		}

		firstTeam := standing[firstTeamName]
		secondTeam := standing[secondTeamName]
		switch outcome {
		case "win":
			firstTeam.win += 1
			firstTeam.score += 3
			secondTeam.loss += 1
		case "draw":
			firstTeam.draw++
			firstTeam.score += 1
			secondTeam.draw++
			secondTeam.score += 1
		case "loss":
			firstTeam.loss++
			secondTeam.win++
			secondTeam.score += 3
		}
	}

	var teamArray byScoreTeamName
	for _, v := range standing {
		teamArray = append(teamArray, *v)
	}
	sort.Sort(byScoreTeamName(teamArray))

	output := fmt.Sprintf("%-31s|%4s|%4s|%4s|%4s|%3s\n", "Team", " MP ", "  W ", "  D ", "  L ", "  P")
	for _, v := range teamArray {
		output += fmt.Sprintf("%-31s|  %d |  %d |  %d |  %d |  %d\n", v.name, (v.win + v.loss + v.draw), v.win, v.draw, v.loss, v.score)
	}
	fmt.Fprint(writer, output)
	return nil
}
