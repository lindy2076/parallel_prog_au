package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var threads int = 2 * 12

func main() {
	var m, n int

	fmt.Println("THIS PROGRAM FINDS ALL PRIMES FROM M TO N.\nENTER M N:")
	_, err := fmt.Scanf("%d %d", &m, &n)
	if err != nil {
		panic(err)
	}
	if m > n {
		panic(fmt.Errorf("%d is more than %d", m, n))
	}

	fmt.Printf("INTERVAL IS [%d, %d]\n", m, n)
	fmt.Println("PRIMES ARE:")
	fmt.Println(findAllPrimesFromInterval(m, n, threads))

	for t := 1; t < threads+1; t++ {
		t1 := time.Now().UnixMicro()
		findAllPrimesFromInterval(m, n, t)
		t2 := time.Now().UnixMicro()
		fmt.Printf("%d GOROUTINE(S) FOUND ALL PRIMES IN %d MICROS\n", t, t2-t1)
	}
}

// find all primes from interval [m,n] in *threads* goroutines (list is not sorted)
func findAllPrimesFromInterval(m, n, threads int) []int {
	var primes []int
	var lock sync.Mutex
	var wg sync.WaitGroup

	wg.Add(threads)
	for t := 0; t < threads; t++ {
		go func(tNum int) {
			defer wg.Done()
			var res []int

			for i := m + tNum; i < n+1; i += threads {
				if isPrime(i) {
					res = append(res, i)
				}
			}
			lock.Lock()
			primes = append(primes, res...)
			lock.Unlock()
		}(t)
	}
	wg.Wait()
	return primes
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	numSqrt := int(math.Floor(math.Sqrt(float64(num))))

	for div := 2; div <= numSqrt; div++ {
		if num%div == 0 {
			return false
		}
	}
	return true
}
