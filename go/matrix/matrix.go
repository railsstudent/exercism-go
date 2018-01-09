package matrix

import "fmt"
import "strings"

type Matrix struct {
	rows [][]int
	cols [][]int
}

func New(strMatrix string) (*Matrix, error) {
	strRows := strings.Split(strMatrix, "\n")
	rows := [][]int{}
	for _, strRow := range strRows {
		elements := strings.Split(strRow, " ")
		for _, e := range elements {
			fmt.Println(int(e))
		}
	}

	fmt.Println(strMatrix)
	fmt.Println()

	m := &Matrix{rows: [][]int{[]int{1}}, cols: [][]int{}}
	return m, nil
}

func (m *Matrix) Cols() [][]int {
	return [][]int{}
}

func (m *Matrix) Rows() [][]int {
	return [][]int{}
}

func (m *Matrix) Set(r int, col int, val int) bool {
	return false
}
