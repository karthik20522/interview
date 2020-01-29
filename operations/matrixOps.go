package operations

import (
	"fmt"
	"strings"
)

//Multiply Return the product of the integers in the matrix
func Multiply(matrixArray [][]int) int {
	result := 1
	xl := len(matrixArray[0])
	yl := len(matrixArray)

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result = matrixArray[i][j] * result
		}
	}

	return result
}

//Sum Return the sum of the integers in the matrix
func Sum(matrixArray [][]int) int {
	result := 0
	xl := len(matrixArray[0])
	yl := len(matrixArray)

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result += matrixArray[i][j]
		}
	}

	return result
}

//Flatten Return the matrix as a 1 line string, with values separated by commas.
func Flatten(records [][]string) string {
	var response string
	for _, row := range records {
		if len(response) > 0 {
			response = fmt.Sprintf("%s,%s", response, strings.Join(row, ","))
		} else {
			response = fmt.Sprintf("%s%s", response, strings.Join(row, ","))
		}
	}

	return response
}

//Invert Return the matrix as a string in matrix format where the columns and rows are inverted
func Invert(matrixArray [][]int) [][]int {
	xl := len(matrixArray[0])
	yl := len(matrixArray)
	result := make([][]int, xl)
	for i := range result {
		result[i] = make([]int, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = matrixArray[j][i]
		}
	}
	return result
}
