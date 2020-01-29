package utils

import "testing"

func TestConvertIntToStringArray(t *testing.T) {
	var matrixArray = []int{1, 2, 3}
	result := ConvertIntToStringArray(matrixArray)
	if result[0] != "1" || result[1] != "2" || result[2] != "3" {
		t.Errorf("Int to String conversion was incorrect, got: %s, want: %s.", result, "1,2,3")
	}
}

func TestMatrix2String(t *testing.T) {
	var matrixArray = [][]int{{1, 2, 3}}
	result := Matrix2String(matrixArray)
	if result != "1,2,3\n" {
		t.Errorf("2D Array to String conversion was incorrect, got: %s, want: %s.", result, "1,2,3")
	}
}
