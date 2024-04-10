package main

import (
	"fmt"
	"math"
)

// Chosen function: x*sin(x)
func f(x float64) float64 {
	return x * math.Sin(x)
}

func main() {
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

func someParallelIntegration(a, b, err float64, threads int) {

	for k := 1; k < threads+1; k++ {

	}
}
