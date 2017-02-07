package main

import (
	"errors"
	"math"
)

// CalcVariance calculates Variance of values
func CalcVariance(values []int) (float64, error) {
	if len(values) <= 0 {
		return 0.0, errors.New("invalid length")
	}

	// calculate average
	avg := 0.0
	for _, v := range values {
		avg += float64(v)
	}
	avg /= float64(len(values))

	variance := 0.0
	for _, v := range values {
		variance += math.Pow(float64(v)-avg, 2)
	}
	variance /= float64(len(values))

	return variance, nil
}

// Calculate sum of row[i]
func RowSum(values []int, i int) (int, error) {
	size := len(values)
	if i < 0 || i >= size {
		return -1, errors.New("invalid index")
	}

	sum := 0
	for j := 0; j < size; j++ {
		sum += values[size*i+j]
	}

	return sum, nil
}

// Calculate sum of col[i]
func ColSum(values []int, i int) (int, error) {
	size := len(values)
	if i < 0 || i >= size {
		return -1, errors.New("invalid index")
	}

	sum := 0
	for j := 0; j < size; j++ {
		sum += values[size*j+i]
	}

	return sum, nil
}

// Calculate diagonal
func DiagonalSum(values []int) (sum int) {
	size := len(values)
	for i := 0; i < size; i++ {
		sum += values[i*i]
	}
	return sum
}

// Calculation matrix evaluated score
// it is variance of sum of each row column and diagonal
func CalcScore(values []int) (score float64, err error) {
	size := len(values)
	sums := make([]int, size*2+1)

	// calculate sum
	for i := 0; i < size; i++ {
		sums[i*2], err = RowSum(values, i)
		if err != nil {
			return -1, err
		}

		sums[i*2+1], err = ColSum(values, i)
		if err != nil {
			return -1, err
		}
	}
	sums[size*size] = DiagonalSum(values)

	score, err = CalcVariance(sums)
	if err != nil {
		return 0.0, nil
	}

	return score, nil
}

func main() {
}
