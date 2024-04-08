package matrix

import (
	"fmt"
	"math/rand"
	"sync"
)

type Matrix struct {
	rows [][]float64
	n    int
	m    int
}

// Create new matrix from double slice
func NewMatrix(sliceOfRows [][]float64) *Matrix {
	n, m := len(sliceOfRows), len(sliceOfRows[0])
	var rows [][]float64 = make([][]float64, n)

	for i := 0; i < n; i++ {
		row := sliceOfRows[i]
		rows[i] = row[:]
	}
	return &Matrix{
		rows: rows,
		n:    n,
		m:    m,
	}
}

// Return new random n x m matrix with float64 values ranging in [a, b)
func RandomMatrix(n, m int, a, b float64) *Matrix {
	var arr [][]float64 = make([][]float64, n)
	for i := 0; i < n; i++ {
		var row []float64 = make([]float64, m)
		for j := 0; j < m; j++ {
			row[j] = rand.Float64()*(b-a) + a
		}
		arr[i] = row
	}
	return NewMatrix(arr)
}

// Returns verbose error if multiplying is not possible
func (m1 *Matrix) CanMultiply(m2 *Matrix) error {
	if m1.m != m2.n {
		return fmt.Errorf("matricies can not be multiplied. M1.m != M2.n (%d != %d)", m1.m, m2.n)
	}
	return nil
}

// Return the [i][j] cell from product matrix of m1 and m2
func (m1 *Matrix) ProdRowAndCol(m2 *Matrix, i, j int) (float64, error) {
	if err := m1.CanMultiply(m2); err != nil {
		return 0, err
	}
	if i >= m1.n || i < 0 {
		return 0, fmt.Errorf("invalid row index (%d): Matrix 1 has only %d rows", i, m1.n)
	}
	if j >= m2.m || j < 0 {
		return 0, fmt.Errorf("invalid column index (%d): Matrix 2 has only %d columns", i, m2.m)
	}
	var res float64
	for k := 0; k < m2.n; k++ {
		res += m1.rows[i][k] * m2.rows[k][j]
	}
	return res, nil
}

// Multiply matrix m by matrix m2
func (m1 *Matrix) Times(m2 *Matrix) (*Matrix, error) {
	if err := m1.CanMultiply(m2); err != nil {
		return nil, err
	}

	var res [][]float64 = make([][]float64, m1.m)
	for i := 0; i < m1.n; i++ {
		res[i] = make([]float64, m2.n)
		for j := 0; j < m2.m; j++ {
			res[i][j], _ = m1.ProdRowAndCol(m2, i, j)
		}
	}
	return NewMatrix(res), nil
}

// Multiply matrix m by matrix m2 in with *threads* goroutines
func (m1 *Matrix) TimesPar(m2 *Matrix, threads int) (*Matrix, error) {
	if err := m1.CanMultiply(m2); err != nil {
		return nil, err
	}
	if threads <= 0 {
		threads = 12
	}

	var wg sync.WaitGroup
	wg.Add(threads)
	var res [][]float64 = make([][]float64, m1.m)
	for i := 0; i < m1.n; i++ {
		res[i] = make([]float64, m2.m)
	}

	for t := 0; t < threads; t++ {
		go func(num int) {
			defer wg.Done()

			for block := num; block < m1.n*m2.m; block += threads {
				var i, j int = block % m1.n, block / m2.m
				rp, _ := m1.ProdRowAndCol(m2, i, j)
				res[i][j] = rp
			}
		}(t)
	}
	wg.Wait()
	return NewMatrix(res), nil
}

// Print matrix to stdout ugly
func PrintMatrix(m *Matrix) {
	for _, row := range m.rows {
		fmt.Print("|")
		for _, v := range row {
			fmt.Printf("%.3f ", v)
		}
		fmt.Print("|\n")
	}
	fmt.Println()
}
