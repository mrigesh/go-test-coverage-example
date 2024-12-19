package sample

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b   int
		result int
	}{
		{1, 1, 2},
		{2, 3, 5},
		{-1, -1, -2},
		{-1, 1, 0},
		{0, 0, 0},
	}

	for _, test := range tests {
		if output := Add(test.a, test.b); output != test.result {
			t.Errorf("Expected %d + %d to equal %d, but got %d", test.a, test.b, test.result, output)
		}
	}
}
