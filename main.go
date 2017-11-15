package main

import (
	"flag"
	"fmt"
)

func fib(n int) {
	sequence := []int{0, 1}
	for i := 1; i < n; i++ {
		sequence = append(sequence, sequence[i]+sequence[i-1])
	}
	for i := range sequence {
		fmt.Println(sequence[i])
	}
}

func main() {
	n := flag.Int("n", 5, "How many numbers in the sequence")
	flag.Parse()
	fib(*n)
}
