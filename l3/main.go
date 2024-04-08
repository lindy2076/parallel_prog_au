package main

import (
	"fmt"
	"parprog/l3/matrix"
	"time"
)

func main() {
	mSizes := [3]int{10, 100, 500}
	threadCount := 2 * 12
	fmt.Printf("Microseconds for each thread count (1-%d) to compute product of 2 square matricies with n=%v:\n", threadCount, mSizes)

	speedsTable := make([][]float64, threadCount)
	for i := range speedsTable {
		speedsTable[i] = make([]float64, len(mSizes))
	}
	for idx, mSize := range mSizes {
		m1, m2 := matrix.RandomMatrix(mSize, mSize, 1, 5), matrix.RandomMatrix(mSize, mSize, 1, 5)

		for threads := 1; threads < threadCount+1; threads++ {
			t1 := time.Now().UnixMicro()
			m1.TimesPar(m2, threads)
			t2 := time.Now().UnixMicro()
			speedsTable[threads-1][idx] = float64(t2 - t1)
		}
	}
	results := matrix.NewMatrix(speedsTable)
	matrix.PrintMatrix(results)
}
