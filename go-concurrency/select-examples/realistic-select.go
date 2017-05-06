package select_examples

import (
	"fmt"
	"math/rand"
	"time"
)

func expensiveComputation(data []int, answer chan int, done chan bool) {
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))

	sum := 0
	for _, n := range data {
		sum += n
	}

	answer <- sum
	done <- true
}

func compute() {
	rand.Seed(time.Now().UTC().UnixNano())

	const allDone = 2
	doneCount := 0

	answerA := make(chan int)
	answerB := make(chan int)
	defer func() {
		close(answerA)
		close(answerB)
	}()

	done := make(chan bool)
	defer func() {
		close(done)
	}()

	go expensiveComputation([]int{1, 2, 3}, answerA, done)
	go expensiveComputation([]int{100, 200, 30000}, answerB, done)

	for doneCount != allDone {
		var which, result int
		select {
		case result = <-answerA:
			which = 'a'
		case result = <-answerB:
			which = 'b'
		case <-done:
			doneCount++
		}

		if which != 0 {
			fmt.Printf("%c -> %d\n", which, result)
		}
	}
}
