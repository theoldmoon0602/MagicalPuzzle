package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func MatrixSize(matrix []int) (int, error) {
	l := len(matrix)
	if l == 0 {
		return 0, nil
	}

	for i := 0; i*i <= l; i++ {
		if i*i == l {
			return i, nil
		}
	}
	return -1, errors.New("matrix is not square-size")
}

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
	size, err := MatrixSize(values)
	if err != nil {
		return 0, err
	}
	if i < 0 || i >= size {
		return 0, errors.New("invalid index")
	}

	sum := 0
	for j := 0; j < size; j++ {
		sum += values[size*i+j]
	}

	return sum, nil
}

// Calculate sum of col[i]
func ColSum(values []int, i int) (int, error) {
	size, err := MatrixSize(values)
	if err != nil {
		return 0, err
	}
	if i < 0 || i >= size {
		return 0, errors.New("invalid index")
	}

	sum := 0
	for j := 0; j < size; j++ {
		sum += values[size*j+i]
	}

	return sum, nil
}

// Calculate diagonal
func DiagonalSum(values []int) (int, error) {
	size, err := MatrixSize(values)
	if err != nil {
		return 0, err
	}

	sum := 0
	for i := 0; i < size; i++ {
		sum += values[size*i+i]
	}
	return sum, nil
}

// Calculation matrix evaluated score
// it is variance of sum of each row column and diagonal
func CalcScore(values []int) (float64, error) {
	size, err := MatrixSize(values)
	if err != nil {
		return 0.0, err
	}
	sums := make([]int, size*2+1)

	// calculate sum
	for i := 0; i < size; i++ {
		sums[i*2], err = RowSum(values, i)
		if err != nil {
			return 0, err
		}

		sums[i*2+1], err = ColSum(values, i)
		if err != nil {
			return 0, err
		}
	}
	sums[size*2], err = DiagonalSum(values)
	if err != nil {
		return 0, err
	}

	score, err := CalcVariance(sums)
	if err != nil {
		return 0.0, nil
	}

	return score, nil
}

func ReadInput(reader io.Reader) ([]int, error) {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords) // split by nl or space

	// read size
	if !sc.Scan() {
		return nil, errors.New("invalid format")
	}
	size, err := strconv.Atoi(sc.Text())
	if err != nil {
		return nil, err
	}

	// read values
	values := make([]int, 0, size*size)

	for i := 0; i < size*size; i++ {
		if !sc.Scan() {
			return nil, errors.New("invalid format")
		}
		v, err := strconv.Atoi(sc.Text())
		if err != nil {
			return nil, err
		}
		values = append(values, v)
	}

	return values, nil
}

func LinearSearch(values []int, needle int) int {
	for i, v := range values {
		if v == needle {
			return i
		}
	}

	return -1
}

func DoOperations(matrix []int, r io.Reader) ([]int, error) {
	// search zero position
	zero := LinearSearch(matrix, 0)
	if zero == -1 {
		return nil, errors.New("no zero found")
	}

	// calculation matrix size
	size, err := MatrixSize(matrix)
	if err != nil {
		return nil, err
	}

	// calculation zero position in 2D space
	var x, y int
	if zero == 0 {
		x, y = 0, 0
	} else {
		y = zero / size
		x = zero % size
	}

	for {
		// read next operator
		op := make([]byte, 1)
		_, err := r.Read(op)
		if op[0] == '\r' || op[0] == '\n' {
			break
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		o := op[0]
		dx, dy := 0, 0
		switch o {
		case 'h':
			if x-1 < 0 {
				return nil, errors.New("invalid operation")
			}
			dx--
		case 'j':
			if y+1 >= size {
				return nil, errors.New("invalid operation")
			}
			dy++
		case 'k':
			if y-1 < 0 {
				return nil, errors.New("invalid operation")
			}
			dy--
		case 'l':
			if x+1 >= size {
				return nil, errors.New("invalid operation")
			}
			dx++
		default:
			return nil, errors.New("invalid operation")
		}
		// swap two values
		a := (y+dy)*size + (x + dx)
		b := y*size + x
		matrix[a], matrix[b] = matrix[b], matrix[a]
		x += dx
		y += dy
	}
	return matrix, nil

}

func DumpMatrix(matrix []int) {
	size, err := MatrixSize(matrix)
	if err != nil {
		return
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%3d", matrix[size*i+j])
		}
		fmt.Print("\n")
	}
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("%s: <input file> <operation file>\n", os.Args[0])
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	opfile, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer opfile.Close()

	values, err := ReadInput(file)
	if err != nil {
		log.Fatal(err)
	}

	res, err := DoOperations(values, opfile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(CalcScore(res))
}
