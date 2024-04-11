package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// Chosen function: x*sin(x)
// 2.80992881892144680466389747281457179113 on [1, 3]
func f(x float64) float64 {
	return x * math.Sin(x)
}

func main() {
	threads := 2 * 12

	var userErr, a, b float64
	fmt.Println("Please enter the error:")
	_, err := fmt.Scanf("%g", &userErr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Scanned error is %g\n", userErr)
	fmt.Println("Please enter integration borders (a b):")
	_, err = fmt.Scanf("%g %g", &a, &b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Scanned borders are [%g, %g]\n", a, b)

	res, n := CalculateIntegralWithErr(a, b, userErr)
	fmt.Printf("Estimated partition count is %d segments\n", n)
	fmt.Printf("Integral is approx. %g\n", res)

	fmt.Printf("Parallel section starts here. There are %d threads\n", threads)
	ns := []int{n, 10 * n, 100 * n, 1000 * n}
	for _, partitions := range ns {
		for k := 1; k < threads+1; k++ {
			someParallelIntegration(a, b, partitions, k)
		}
	}
}

// Calculate integral numerically via Left Rectangle integration method
// on [a, b] with n segments of equal length.
func LRec(a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	res := f(a)
	for i := 1; i < n; i++ {
		res += f(a + float64(i)*h)
	}
	return res * h
}

// Calculate integral numerically via Left Rectangle integration method
// on [a, b] with estimated error. Returns the integral value and
// partition count to achieve *err* precision.
func CalculateIntegralWithErr(a, b, err float64) (float64, int) {
	prev_result := 0.0
	var n int = 1
	res := LRec(a, b, n)
	for n = 2; math.Abs(res-prev_result) > err; n++ {
		prev_result, res = res, LRec(a, b, n)
	}
	return res, n
}

// Calculate integral numerically via bla bla with *partitions* partitions
// and *threads* goroutines. [a, b] is divided into *threads* segments
// each processed by a goroutine. Each segment is divided into partitions/threads
// fragments which are then calculated by LRec method.
func someParallelIntegration(a, b float64, partitions, threads int) {
	t1 := time.Now().UnixMicro()
	block_len := (b - a) / float64(threads)
	sums_chan := make(chan float64, threads)

	var wg sync.WaitGroup
	wg.Add(threads)
	for thread := 1; thread < threads+1; thread++ {
		go func(tNum int) {
			defer wg.Done()
			left_border := a + block_len*float64(tNum-1)
			right_border := a + block_len*float64(tNum)
			block_sum := LRec(left_border, right_border, partitions/threads)
			sums_chan <- block_sum
		}(thread)
	}
	wg.Wait()
	close(sums_chan)

	finalSum := 0.0
	for v := range sums_chan {
		finalSum += v
	}
	t2 := time.Now().UnixMicro()
	fmt.Printf("Threads: %d, partitions: %d, integral: %g, time (micros): %d\n", threads, partitions, finalSum, t2-t1)
}
