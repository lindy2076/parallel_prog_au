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

var threads int = 2 * 12

func main() {
	var sel string
	fmt.Println("WELCOME TO PI OR E ESTIMATION. WRITE 'PI' FOR PI AND EVERYTING ELSE FOR E:")
	_, err := fmt.Scanf("%s", &sel)
	if err != nil {
		panic(err)
	}
	if sel == "PI" {
		mainPi()
	} else {
		mainE()
	}
}

// Estimate number of iterations and estimate using *threads* goroutines
func mainPi() {
	var userErr float64
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
		prevRes *= calcNthTerm_Pi(n)
	}
	return n - 1
}

// Calculate Pi = 2 * prod_{i=1}^n (4i^2/(4i^2 - 1))
func kindaPi(n int) float64 {
	res := 2.0
	for i := 1; i < n+1; i++ {
		res *= calcNthTerm_Pi(i)
	}
	return res
}

// Get nth term for kindaPi formula
func calcNthTerm_Pi(n int) float64 {
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
				threadRes *= calcNthTerm_Pi(ns)
			}
			piRes.mu.Lock()
			defer piRes.mu.Unlock()
			piRes.res *= threadRes
		}(thread)
	}
	wg.Wait()
	return piRes.res
}

func mainE() {
	var userErr float64
	fmt.Println("HELLO THERE. PLEASE ENTER THE DESIRED ERROR FOR E ESTIMATION:")
	_, err := fmt.Scanf("%g", &userErr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ERROR IS: %g\n", userErr)

	nE := nForGiverErr_E(userErr)
	fmt.Printf("%d ITERATIONS ARE ENOUGH FOR THE DESIRED ERROR\n", nE)
	fmt.Printf("%g IS THE RESULT\n", kindaE(nE))

}

func nForGiverErr_E(uErr float64) int {
	n := 1
	prevRes := kindaE(n)
	for n = 2; math.Abs(prevRes-math.E) > uErr; n++ {
		fmt.Printf("%g\n", prevRes)
		prevRes *= calcNthTerm_E(n)
	}
	return n - 1
}

func kindaE(n int) float64 {
	res := 2.0
	for i := 1; i < n+1; i++ {
		res *= calcNthTerm_E(i)
	}
	return res
}

func calcNthTerm_E(n int) float64 {
	// fn := float64(n)
	// t1 := math.Pow(fn*2+3, fn+0.5)
	// t2 := math.Pow(fn*2-1, fn-0.5)
	// return t1 * t2 / math.Pow(fn*2+1, fn*2)
	top := 1
	bottom := 1
	for i := 0; i < intPow(2, n-1); i++ {
		top *= intPow(2, n+1) - i*2
		bottom *= intPow(2, n+1) - 1 - i*2
	}
	return math.Pow(float64(top)/float64(bottom), 1/float64(intPow(2, n)))
}

func intPow(x, n int) int {
	if n == 1 {
		return x
	}
	if n == 0 {
		return 1
	}
	return x * intPow(x, n-1)
}
