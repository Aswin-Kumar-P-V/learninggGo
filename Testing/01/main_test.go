package main

import "testing"

type test struct {
	data   []int
	answer int
}

func TestSum(t *testing.T) {

	tests := []test{
		test{[]int{10, 12}, 22},
		test{[]int{2, 3}, 5},
		test{[]int{10, 15}, 25},
		test{[]int{4, -3, 15}, 16},
	}

	for _, value := range tests {
		gotSum := sum(value.data...)
		if gotSum != value.answer {
			t.Error("Expected", value.answer, "Got", gotSum)
		}
	}
}
