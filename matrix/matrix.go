package matrix

import (
	"errors"
)

//matrix represents well matrix :)
type matrix [][]int

//New returns matrix from a given string.
func New(input string) (matrix, error) {
	var result matrix = make([][]int, 0)
	var numbers = make([]int, 0)
	var num, rowLenght int
	//check if we have last row empty
	if input[len(input)-1] == '\n' {
		return result, errors.New("empty last row")
	}
	for i, char := range input {
		//ignore extra unecessary spaces
		if i > 0 && !(input[i-1] >= '0' && input[i-1] <= '9') && char == ' ' {
			continue
		}
		//don't accept any character excpet from digits, spaces and new line
		if !(char >= '0' && char <= '9') && char != ' ' && char != '\n' {
			return nil, errors.New("wrong character in string")
		}
		//if we find digit we make it a part of number
		if char >= '0' && char <= '9' {
			num = num*10 + (int(char) - 48)
			if num < 0 {
				return nil, errors.New("number overflow")
			}
		}
		//finish reading number when spacer or new line. Append it to row slice
		if char == ' ' || char == '\n' || i == len(input)-1 {
			numbers = append(numbers, num)
			num = 0
		}

		//if we find or new line or end of string add row slice to slice of rows.
		if char == '\n' || i == len(input)-1 {
			//check if we have the same row length
			if len(result) > 0 && len(numbers) != rowLenght {
				return nil, errors.New("wrong matrix")
			}
			rowLenght = len(numbers)
			result = append(result, numbers)
			numbers = []int{}
		}
	}
	return result, nil
}

func (in *matrix) Rows() [][]int {
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

func (in *matrix) Cols() [][]int {
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

func (in *matrix) Set(row, col, val int) bool {
	if (row < 0 || row >= len(*in)) || (col < 0 || col >= len((*in)[0])) {
		return false
	}
	(*in)[row][col] = val
	return true
}
