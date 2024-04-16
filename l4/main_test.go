package main

import "testing"

func TestCeilToPowerOf2(t *testing.T) {
	var data = []struct {
		input    int
		expected int
	}{
		{1, 2},
		{2, 4},
		{3, 4},
		{4, 8},
		{7, 8},
		{8, 16},
		{15, 16},
		{16, 32},
		{64, 128},
	}

	for _, tt := range data {
		t.Run("t", func(t *testing.T) {
			res, _ := CeilToPowerOf2(tt.input)
			if res != tt.expected {
				t.Errorf("CeilToPowerOf2 (%d) got %d, expected %d", tt.input, res, tt.expected)
			}
		})
	}

}

func TestCalcKthTerm_E(t *testing.T) {
	var data = []struct {
		input    int
		expected float64
	}{
		{1, calcNthTerm_E(1)},
		{2, calcNthTerm_E(2)},
		{3, calcNthTerm_E(3)},
		{4, calcNthTerm_E(4)},
		{5, calcNthTerm_E(5)},
		{6, calcNthTerm_E(6)},
		{7, calcNthTerm_E(7)},
		{8, calcNthTerm_E(8)},
	}

	for _, tt := range data {
		t.Run("t", func(t *testing.T) {
			n := tt.input
			kmin, kmax := intPow(2, n-1), intPow(2, n)-1
			res := 1.0
			for i := kmin; i < kmax+1; i++ {
				res *= calcKthTerm_E(i)
			}
			if res != tt.expected {
				t.Errorf("TestCalcKthTerm_E (%d) got %g, expected %g", tt.input, res, tt.expected)
			}
		})
	}

}
