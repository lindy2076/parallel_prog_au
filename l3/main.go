package main

import (
	"fmt"
	"parprog/l3/matrix"
	"time"
)

func main() {
	mSizes := [3]int{10, 100, 500}
	threadCount := 2 * 12

	fmt.Printf("Microseconds for each thread count (1-%d) to compute product of 2 square matricies with n=%v:\n\n", threadCount, mSizes)
	fmt.Printf("threads:")
	for i := 0; i < threadCount; i++ {
		fmt.Printf("\t%d", i+1)
	}
	fmt.Println()

	for _, mSize := range mSizes {
		m1, m2 := matrix.RandomMatrix(mSize, mSize, 1, 5), matrix.RandomMatrix(mSize, mSize, 1, 5)
		fmt.Printf("n=%d:", mSize)
		for threads := 1; threads < threadCount+1; threads++ {
			t1 := time.Now().UnixMicro()
			m1.TimesPar(m2, threads)
			t2 := time.Now().UnixMicro()

			fmt.Printf("\t%d", t2-t1)
		}
		fmt.Println()
	}
}
