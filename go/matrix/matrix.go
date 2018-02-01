package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	rows [][]int
	cols [][]int
}

func validate(strMatrix string) (int, int, error) {
	strRows := strings.Split(strMatrix, "\n")
	numRows := len(strRows)
	numCols := len(strings.Split(strings.TrimSpace(strRows[0]), " "))
	for i := 1; i < numRows; i++ {
		if len(strings.Split(strings.TrimSpace(strRows[i]), " ")) != numCols {
			return -1, -1, errors.New("Invalid matrix, rows have different number of columns")
		}
	}
	return numRows, numCols, nil
}

// New
func New(strMatrix string) (*Matrix, error) {
	var rows [][]int
	var cols [][]int
	var row []int

	_, numCols, invalidMatrix := validate(strMatrix)
	if invalidMatrix != nil {
		return nil, invalidMatrix
	}

	strRows := strings.Split(strMatrix, "\n")
	for i := 0; i < numCols; i++ {
		cols = append(cols, []int{})
	}
	for _, strRow := range strRows {
		row = []int{}
		strRow = strings.TrimSpace(strRow)
		for i, strNum := range strings.Split(strRow, " ") {
			n, err := strconv.Atoi(strNum)
			if err != nil {
				return nil, err
			}
			row = append(row, n)
			cols[i] = append(cols[i], n)
		}
		rows = append(rows, row)
	}
	m := &Matrix{rows, cols}
	return m, nil
}

// Cols returns a copy of matrix's columns
func (m *Matrix) Cols() [][]int {
	copiedCols := make([][]int, len(m.cols))
	for i, colSlice := range m.cols {
		for _, colValue := range colSlice {
			copiedCols[i] = append(copiedCols[i], colValue)
		}
	}
	return copiedCols
}

// Rows returns a copy of matrix's rows
func (m *Matrix) Rows() [][]int {
	copiedRows := make([][]int, len(m.rows))
	for i, origRow := range m.rows {
		for _, rowValue := range origRow {
			copiedRows[i] = append(copiedRows[i], rowValue)
		}
	}
	return copiedRows
}

// Set updates matrix value, return true if row index and col index are valid, false otherwise
func (m *Matrix) Set(r int, c int, val int) bool {
	if r >= 0 && r < len(m.rows) && c >= 0 && c < len(m.cols) {
		m.rows[r][c] = val
		m.cols[c][r] = val
		return true
	}
	return false
}
