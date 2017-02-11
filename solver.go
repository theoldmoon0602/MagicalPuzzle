package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
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
			return nil, fmt.Errorf("%d: invalid format", i)
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

type Puzzle struct {
	x, y   int
	values []int
	size   int
}

func GetZeroPos(matrix []int) (int, int, error) {
	zero := LinearSearch(matrix, 0)
	if zero == -1 {
		return 0, 0, errors.New("no zero found")
	}

	// calculation matrix size
	size, err := MatrixSize(matrix)
	if err != nil {
		return 0, 0, err
	}

	// calculation zero position in 2D space
	y := zero / size
	x := zero % size

	return x, y, nil
}

func NewPuzzle(matrix []int) (*Puzzle, error) {
	puzzle := Puzzle{}
	puzzle.values = matrix
	x, y, err := GetZeroPos(matrix)
	if err != nil {
		return nil, err
	}
	puzzle.x = x
	puzzle.y = y
	size, err := MatrixSize(matrix)
	if err != nil {
		return nil, err
	}
	puzzle.size = size
	return &puzzle, nil
}

func (p *Puzzle) DoOperation(o byte) error {
	dx, dy := 0, 0
	switch o {
	case 'h':
		if p.x-1 < 0 {
			return fmt.Errorf("invalid operation->%c", o)
		}
		dx--
	case 'j':
		if p.y+1 >= p.size {
			return fmt.Errorf("invalid operation->%c", o)
		}
		dy++
	case 'k':
		if p.y-1 < 0 {
			return fmt.Errorf("invalid operation->%c", o)
		}
		dy--
	case 'l':
		if p.x+1 >= p.size {
			return fmt.Errorf("invalid operation->%c", o)
		}
		dx++
	default:
		return fmt.Errorf("unknown operation->%c", o)
	}
	// swap two values
	a := (p.y+dy)*p.size + (p.x + dx)
	b := p.y*p.size + p.x
	p.values[a], p.values[b] = p.values[b], p.values[a]
	p.x += dx
	p.y += dy

	return nil
}

func (p *Puzzle) DoOperations(r io.Reader) error {
	for i := 0; ; i++ {
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
			return err
		}

		err = p.DoOperation(op[0])
		if err != nil {
			return err
		}
	}
	return nil
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

func Solve(p *Puzzle) (float64, error) {
	rand.Seed(time.Now().UnixNano())
	ops := []byte{'h', 'j', 'l', 'k'}

	curScore, _ := CalcScore(p.values)

	for i := 0; i < 1000; {
		o := rand.Intn(4)

		err := p.DoOperation(ops[o])
		if err != nil {
			continue
		}

		score, err := CalcScore(p.values)
		if err != nil {
			return 0.0, err
		}

		if (score > curScore) && rand.Intn(5) == 0 {
			pi := (o + 2) % 4
			p.DoOperation(ops[pi])
		} else {
			if score < 5 {
				break
			}
			curScore = score
			fmt.Printf("%c", ops[o])
			i++
		}
	}
	return curScore, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("HogeEEEE")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	values, err := ReadInput(f)
	if err != nil {
		log.Fatal(err)
	}

	p, err := NewPuzzle(values)
	if err != nil {
		log.Fatal(err)
	}
	score, err := CalcScore(p.values)
	v := score
	if err != nil {
		log.Fatal(err)
	}

	score, err = Solve(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nfrom %e(%.2f) -> ", v, v)
	fmt.Printf("%e(%.2f)\n", score, score)
	// DumpMatrix(p.values)
}
