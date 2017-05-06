package select_examples

import (
	"fmt"
	"math/rand"
	"time"
)

func rollDice() {
	channels := make([]chan bool, 6)

	for i, _ := range channels {
		channels[i] = make(chan bool)
	}

	defer func() {
		for _, c := range channels {
			close(c)
		}
	}()

	go func() {
		for {
			channels[rand.Intn(6)] <- true
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 5; i++ {
		var x int
		select {
		case <-channels[0]:
			x = 1
		case <-channels[1]:
			x = 2
		case <-channels[2]:
			x = 3
		case <-channels[3]:
			x = 4
		case <-channels[4]:
			x = 5
		case <-channels[5]:
			x = 6
		}
		fmt.Printf("%d \n", x)
	}
}
