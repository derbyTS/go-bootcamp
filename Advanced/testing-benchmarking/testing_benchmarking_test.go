package main

import (
	// "fmt"
	"math/rand"
	"testing"
)

func GeneratingRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(100)
	}

	return slice
}

func SumSlice(slice []int) int {
	sum := 0

	for _, v := range slice {
		sum += v
	}

	return sum
}

func TestGenerateRandomeSlice(t *testing.T) {
	size := 100
	slice := GeneratingRandomSlice(size)
	if len(slice) != size {
		t.Errorf("Exptected slice size %d, received %d", size, len(slice))
	}
}

func BenchmarkGenerateRandomSlice(b *testing.B) {
	for range b.N {
		GeneratingRandomSlice(1000)
	}
}

func BenchmarkSumOfSlice(b *testing.B) {
	slice := GeneratingRandomSlice(1000)
	b.ResetTimer() // Setting time during setup before the actual test can skew the result. GeneratingRandomSlice can take a long time
	for range b.N {
		SumSlice(slice)
	}
}

// ==============This is for Testing and Benchmarking
// func Add(a, b int) int {
// 	return a + b
// }

// ==============Benchmarking
// func BenchmarkAddSmallInput(b *testing.B) {
// 	for range b.N {
// 		Add(2, 3)
// 	}
// }
//
// func BenchmarkAddMediumInput(b *testing.B) {
// 	for range b.N {
// 		Add(200, 300)
// 	}
// }
//
// func BenchmarkAddLargeInput(b *testing.B) {
// 	for range b.N {
// 		Add(2000, 3000)
// 	}
// }

// ==============Testing

// func TestAddSubTest(t *testing.T) {
// 	tests := []struct{ a, b, expected int }{
// 		{2, 3, 5},
// 		{0, 0, 0},
// 		{-1, 2, 1},
// 	}
//
// 	for _, test := range tests {
// 		t.Run(fmt.Sprintf("Add(%d %d)", test.a, test.b), func(t *testing.T) {
// 			result := Add(test.a, test.b)
// 			if result != test.expected {
// 				t.Errorf("result = %d; expected = %d", result, test.expected)
// 			}
// 		})
// 	}
// }

// func TestAddTableDriven(t *testing.T) {
// 	tests := []struct{ a, b, expected int }{
// 		{2, 3, 5},
// 		{0, 0, 0},
// 		{-1, 2, 1},
// 	}
//
// 	for _, test := range tests {
// 		result := Add(test.a, test.b)
// 		if result != test.expected {
// 			t.Errorf("Add(%d, %d) = %d, expected %d\n", test.a, test.b, result, test.expected)
// 		}
// 	}
// }

// func TestAdd(t *testing.T) {
// 	result := Add(2, 3)
// 	expected := 5
// 	if result != expected {
// 		t.Errorf("Add(2, 3) = %d; expected: %d\n", result, expected)
// 	}
// }
