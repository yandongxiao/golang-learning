package main

import "fmt"

var deposits = make(chan int)
var balances = make(chan int)

func CDeposit(money int) { deposits <- money } // 只有一行语句的函数
func CBalance() int      { return <-balances }

func broker() {
	money := 0 // 原本共享的变量，变成协程独有
	for {
		select {
		case inc := <-deposits:
			money += inc
		case balances <- money: // NOTICE: get money的方法. case语句为空
		}
	}
}

func ExampleChan() {
	go broker()
	CDeposit(100)
	CDeposit(100)
	fmt.Println(CBalance())
	//Output:
	//200
}