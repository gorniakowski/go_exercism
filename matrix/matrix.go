package matrix

import (
	"errors"
	"strconv"
	"strings"
)

//Matrix represents well matrix :)
type Matrix [][]int

//New returns matrix from a given string.
func New(input string) (Matrix, error) {
	result := make([][]int, 0)
	//split input into rows
	rows := strings.Split(input, "\n")
	//loop over each row
	for i, row := range rows {
		row = strings.TrimSpace(row)
		rowSplit := strings.Split(row, " ")
		rowInts := make([]int, 0, len(rowSplit))
		//check if each row has the same size
		if i > 0 && len(rowSplit) != len(result[i-1]) {
			return nil, errors.New("wrong matrix")
		}
		//parse input to ints, and and to row made of ints
		for _, elm := range rowSplit {
			number, err := strconv.Atoi(elm)
			if err != nil {
				return nil, err
			}
			rowInts = append(rowInts, number)
		}
		//add parsed row to a matrix
		result = append(result, rowInts)
	}
	return result, nil

}

//Rows returns deep copy of matrix, row by row
func (in *Matrix) Rows() [][]int {
	out := make([][]int, 0, len(*in))
	for _, row := range *in {
		newRow := make([]int, 0, len(row))
		for _, num := range row {
			newRow = append(newRow, num)
		}
		out = append(out, newRow)
	}
	return out
}

//Cols returns deep copy of Matrix, trasposed columns into rows
func (in *Matrix) Cols() [][]int {
	out := make([][]int, 0, len((*in)[0]))
	for i := 0; i < len((*in)[0]); i++ {
		newRow := make([]int, 0, len((*in)[0]))
		for _, row := range *in {
			newRow = append(newRow, row[i])
		}
		out = append(out, newRow)
	}

	return out
}

//Set allows to change desired matrix element, returns true is success
func (in *Matrix) Set(row, col, val int) bool {
	if (row < 0 || row >= len(*in)) || (col < 0 || col >= len((*in)[0])) {
		return false
	}
	(*in)[row][col] = val
	return true
}
