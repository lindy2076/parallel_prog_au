package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var k int
	fmt.Println("Enter the goroutines amount")
	fmt.Scanf("%d", &k)

	var wg sync.WaitGroup
	wg.Add(k)
	for i := 0; i < k; i++ {
		go func(i int) {
			defer wg.Done()

			fmt.Printf("Hello! I am goroutine %d from %d goroutines\n", i, k)
			time.Sleep(5 * time.Second)
			if i%2 == 0 {
				fmt.Printf("Hello again. I am goroutine %d. I am even\n", i)
			}
		}(i)
	}
	wg.Wait()
}
