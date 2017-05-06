package main

import "fmt"

func createCounter(start int) chan int {
	next := make(chan int)

	go func(i int) {
		for {
			next <- i
			i++
		}
	}(start)

	return next
}

func main() {
	count := createCounter(0)

	fmt.Printf("Count is %d\n", <-count)
	fmt.Printf("Count is %d\n", <-count)
	fmt.Printf("Count is %d\n", <-count)
	fmt.Printf("Count is %d\n", <-count)
	fmt.Printf("Count is %d\n", <-count)
	fmt.Printf("Count is %d\n", <-count)

	close(count)
}
