package advanced

import (
	"fmt"
	"testing"
)

// filename for testing should needs to be _test.go
// subtests allow grouping related tests and provide test better test organisation tests and reporting
// with t.Run subtests failures willbe written clearly
func Add(a, b int) int {
	return a + b
}

func TestAddSubtests(t *testing.T) {
	tests := []struct{ a, b, expected int }{
		{2, 3, 5},
		{0, 1, 0},
		{-1, 1, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Add(%d,%d)", test.a, test.b), func(t *testing.T) {
			result := Add(test.a, test.b)
			if result != test.expected {
				t.Errorf("result = %d; want %d", result, test.expected)
			}
		})
	}
}

// func TestAddTableDriven(t *testing.T) {
// 	test := []struct{ a, b, expected int }{
// 		{2, 3, 5},
// 		{0, 1, 0},
// 		{-1, 1, 0},
// 	}

// 	for _, test := range test {
// 		result := Add(test.a, test.b)
// 		if result != test.expected {
// 			t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
// 		}
// 	}
// }

//Functions for testing needs to start with Test and we also need to use a pointer to our testing function
// func TestAdd(t *testing.T) {
// 	result := Add(2, 4)
// 	expected := 5
// 	if result != expected {
// 		t.Errorf("Add(2,3) = %d; want %d", result, expected)
// 	}
// }
