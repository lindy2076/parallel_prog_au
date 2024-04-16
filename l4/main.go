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

	fmt.Printf("%d GOROUTINES ARE READY TO CALCULATE\n", threads)
	for k := 1; k < threads+1; k++ {
		t1 := time.Now().UnixMicro()
		parallelE(nE, k)
		t2 := time.Now().UnixMicro()
		fmt.Printf("%d THREADS CALCULATED IN %d MICROS\n", k, t2-t1)
	}
}

func nForGiverErr_E(uErr float64) int {
	n := 1
	prevRes := kindaE(n)
	for n = 2; math.Abs(prevRes-math.E) > uErr; n++ {
		// prevRes *= calcNthTerm_E(n)
		prevRes *= calcKthTerm_E(n)
	}
	return n - 1
}

func kindaE(n int) float64 {
	res := 2.0
	for i := 1; i < n+1; i++ {
		// res *= calcNthTerm_E(i)
		res *= calcKthTerm_E(i)
	}
	return res
}

// original function to estimate. N is for each root
func calcNthTerm_E(n int) float64 {
	res := 1.0
	for i := 0; i < intPow(2, n-1); i++ {
		d := intPow(2, n+1) - i*2
		res *= math.Pow(float64(d)/float64(d-1), 1/float64(intPow(2, n)))
	}
	return res
}

// calcNthTerm split into terms.
func calcKthTerm_E(k int) float64 {
	ceil2n, n := CeilToPowerOf2(k)
	max := ceil2n * 2
	maxAmount := intPow(2, n-1)
	d := max - 2*(maxAmount) + 2*(ceil2n-k)

	return math.Pow(float64(d)/float64(d-1), 1/float64(ceil2n))
}

// Ceil to the nearest power of 2 that is bigger than num.
func CeilToPowerOf2(num int) (int, int) {
	n := 0
	for i := 1; i <= num; i = i << 1 {
		n += 1
	}
	return intPow(2, n), n
}

// Returns x to the power of n.
func intPow(x, n int) int {
	if n == 1 {
		return x
	}
	if n == 0 {
		return 1
	}
	return x * intPow(x, n-1)
}

// Estimate e with *threads* goroutines and n iterations of kindaE
func parallelE(n, threads int) float64 {
	var wg sync.WaitGroup
	wg.Add(threads)

	piRes := &ResWithLock{res: 2.0}

	for thread := 1; thread < threads+1; thread++ {
		go func(tNum int) {
			defer wg.Done()

			threadRes := 1.0
			for ns := tNum; ns < n; ns += threads {
				// threadRes *= calcNthTerm_E(ns)
				threadRes *= calcKthTerm_E(ns)
			}
			piRes.mu.Lock()
			defer piRes.mu.Unlock()
			piRes.res *= threadRes
		}(thread)
	}
	wg.Wait()
	return piRes.res
}
