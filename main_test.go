package main

import (
	"math"
	"testing"
)

func TestMatrixSIze(t *testing.T) {
	values := make([]int, 100, 100)
	size, err := MatrixSize(values)
	expected := 10
	if err != nil {
		t.Error(err)
	}
	if size != expected {
		t.Errorf("invalid size. expected %d, actually %d\n", expected, size)
	}

	values = make([]int, 10, 10)
	size, err = MatrixSize(values)
	if err == nil {
		t.Error("no error when should happen")
	}
}

func TestCalcVariance(t *testing.T) {
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

func TestRowSum(t *testing.T) {
	matrix := []int{
		0, 1, 2, 3, 4,
		5, 6, 7, 8, 9,
		2, 3, 4, 5, 6,
		7, 8, 9, 0, 1,
		3, 4, 5, 6, 7,
	}
	expected := []int{10, 35, 20, 25, 25}
	for i, v := range expected {
		sum, err := RowSum(matrix, i)
		if err != nil {
			t.Error(err)
		}
		if sum != v {
			t.Errorf("rowsum %d is wrong. expected %d - actually %d.\n", i, v, sum)
		}
	}
}

func TestColSum(t *testing.T) {
	matrix := []int{
		0, 1, 2, 3, 4,
		5, 6, 7, 8, 9,
		2, 3, 4, 5, 6,
		7, 8, 9, 0, 1,
		3, 4, 5, 6, 7,
	}
	expected := []int{17, 22, 27, 22, 27}
	for i, v := range expected {
		sum, err := ColSum(matrix, i)
		if err != nil {
			t.Error(err)
		}
		if sum != v {
			t.Errorf("ColSum %d is wrong. expected %d - actually %d.\n", i, v, sum)
		}
	}
}

func TestDiagonalSum(t *testing.T) {
	matrix := []int{
		0, 1, 2, 3, 4,
		5, 6, 7, 8, 9,
		2, 3, 4, 5, 6,
		7, 8, 9, 0, 1,
		3, 4, 5, 6, 7,
	}
	sum, err := DiagonalSum(matrix)
	if err != nil {
		t.Error(err)
	}
	if sum != 17 {
		t.Errorf("DiagonalSum is wrong. expected 17, returned %d\n", sum)
	}
}
