package main

import (
	"math"
	"testing"
)

func testCalcVariance(t *testing.T) {
	values := []int{10, 10, 10, 10, 10}
	variance, err := CalcVariance(values)
	if err != nil {
		t.Error(err)
	}
	if variance > 1e-10 {
		t.Errorf("variance is wrong %f", variance)
	}

	values = []int{}
	variance, err = CalcVariance(values)
	if err == nil {
		t.Error("variance is not return nil")
	}

	values = []int{1, 2, 3, 4, 5, 6, 7}
	variance, err = CalcVariance(values)
	if err != nil {
		t.Error(err)
	}
	if math.Abs(variance-4.0) > 1e-10 {
		t.Errorf("variance is wrong: %f", variance)
	}
}
