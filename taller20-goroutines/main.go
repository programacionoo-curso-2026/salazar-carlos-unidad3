package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {

		go ShowGoroutines(1)

	}

	time.Sleep(time.Minute)

}

func ShowGoroutines(id int) {
	delay := rand.Intn(500)
	fmt.Printf("Goroutine #%d\n with %dms\n", id, delay)
	time.Sleep(time.Millisecond * time.Duration(delay))
}
