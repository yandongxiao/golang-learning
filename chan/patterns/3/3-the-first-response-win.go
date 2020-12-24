package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 输入参数是channel
func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3)+1
	// Simulate a workload, and sleep 1s/2s/3s.
	time.Sleep(time.Duration(rb) * time.Second)
	c <- ra
}

func main() {
	rand.Seed(time.Now().UnixNano())
	startTime := time.Now()
	// NOTE: c must be a buffered channel.
	c := make(chan int32, 5)
	for i := 0; i < cap(c); i++ {
		go source(c)
	}

	// Only the first response will be used.
	rnd := <-c
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)
}