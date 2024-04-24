package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

const FILENAME string = "mf"

var wg sync.WaitGroup
var c = make(chan string)

func write(c chan<- string) {
	defer close(c)
	defer wg.Done()

	f, err := os.Open(FILENAME)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		symb := scanner.Text()
		time.Sleep(time.Duration(rand.Int()%100) * time.Millisecond)
		c <- symb
	}
	c <- "\n"

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func read(c <-chan string) {
	defer wg.Done()
	for symbol := range c {
		fmt.Print(symbol)
		time.Sleep(time.Duration(rand.Int()%50) * time.Millisecond)
	}
}

func main() {
	wg.Add(2)

	go write(c)
	go read(c)

	wg.Wait()
}
