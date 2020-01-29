package operations

import "testing"

func TestMultiply(t *testing.T) {
	var matrixArray = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	result := Multiply(matrixArray)
	if result != 362880 {
		t.Errorf("Mutiply Result was incorrect, got: %d, want: %d.", result, 362880)
	}
}

func TestSum(t *testing.T) {
	var matrixArray = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	result := Sum(matrixArray)
	if result != 45 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", result, 45)
	}
}

func TestFlatten(t *testing.T) {
	var matrixArray = [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	result := Flatten(matrixArray)
	if result != "1,2,3,4,5,6,7,8,9" {
		t.Errorf("Flattened String was incorrect, got: %s, want: %s.", result, "1,2,3,4,5,6,7,8,9")
	}
}

func TestInvert(t *testing.T) {
	var matrixArray = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	result := Invert(matrixArray)
	if result[0][0] != 1 || result[0][1] != 4 || result[0][2] != 7 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", result[0], []int{1, 4, 7})
	}
}
