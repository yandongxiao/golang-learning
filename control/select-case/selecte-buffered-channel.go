package main

import "fmt"

func ExampleD() {
	ch := make(chan int, 1)
	select {
	case <-ch:
		fmt.Println("receive from chan")
	case ch <- 10:
		fmt.Println("send to chan")
	}
}