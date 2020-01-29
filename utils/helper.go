package utils

import (
	"fmt"
	"strconv"
	"strings"
)

//ConvertIntToStringArray Convert array of Integer to array of String
func ConvertIntToStringArray(values []int) []string {
	valuesText := []string{}
	for i := range values {
		number := values[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	return valuesText
}

//Matrix2String Convert 2D Array to comma seperated string
func Matrix2String(matrixArray [][]int) string {
	var response string
	for _, row := range matrixArray {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(ConvertIntToStringArray(row), ","))
	}

	return response
}

//ConvertCSVToMatrix Convert 2D Array of String to 2D array of Int's
func ConvertCSVToMatrix(csvData [][]string) ([][]int, error) {
	xl := len(csvData[0])
	yl := len(csvData)

	csvMatrix := make([][]int, xl)
	for i := range csvMatrix {
		csvMatrix[i] = make([]int, yl)
	}

	for i, line := range csvData {
		for j, val := range line {
			var valInt int
			var err error
			if val == "" {
				valInt = 0
			} else {
				valInt, err = strconv.Atoi(val)
			}
			if err != nil {
				return nil, fmt.Errorf("Given value is not a number, %s", val)
			}
			csvMatrix[i][j] = valInt
		}
	}

	return csvMatrix, nil
}
