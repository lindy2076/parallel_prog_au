package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type ResWithLock struct {
	res float64
	mu  sync.Mutex
}

const Pi = math.Pi

func main() {
	var userErr float64
	threads := 2 * 12
	fmt.Println("HELLO THERE. PLEASE ENTER THE DESIRED ERROR FOR PI ESTIMATION:")
	_, err := fmt.Scanf("%g", &userErr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ERROR IS: %g\n", userErr)

	nPi := nForGivenErr_Pi(userErr)
	fmt.Printf("%d ITERATIONS ARE ENOUGH FOR THE DESIRED ERROR\n", nPi)
	fmt.Printf("%g IS THE RESULT\n", kindaPi(nPi))

	fmt.Printf("%d GOROUTINES ARE READY TO CALCULATE\n", threads)
	for k := 1; k < threads+1; k++ {
		t1 := time.Now().UnixMicro()
		parallelPi(nPi, k)
		t2 := time.Now().UnixMicro()
		fmt.Printf("%d THREADS CALCULATED IN %d MICROS\n", k, t2-t1)
	}
}

// Get iterations number for kindaPi formula to achieve uErr error
func nForGivenErr_Pi(uErr float64) int {
	n := 1
	prevRes := kindaPi(n)
	for n = 2; math.Abs(prevRes-math.Pi) > uErr; n++ {
		prevRes *= calcNthterm(n)
	}
	return n - 1
}

// Calculate Pi = 2 * prod_{i=1}^n (4i^2/(4i^2 - 1))
func kindaPi(n int) float64 {
	res := 2.0
	for i := 1; i < n+1; i++ {
		res *= calcNthterm(i)
	}
	return res
}

// Get nth term for kindaPi formula
func calcNthterm(n int) float64 {
	nn4 := float64(n * n * 4)
	return nn4 / (nn4 - 1)
}

// Estimate Pi with *threads* goroutines and n iterations of kindaPi
func parallelPi(n, threads int) float64 {
	var wg sync.WaitGroup
	wg.Add(threads)

	piRes := &ResWithLock{res: 2.0}

	for thread := 1; thread < threads+1; thread++ {
		go func(tNum int) {
			defer wg.Done()

			threadRes := 1.0
			for ns := tNum; ns < n; ns += threads {
				threadRes *= calcNthterm(ns)
			}
			piRes.mu.Lock()
			defer piRes.mu.Unlock()
			piRes.res *= threadRes
		}(thread)
	}
	wg.Wait()
	return piRes.res
}
