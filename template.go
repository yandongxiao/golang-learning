// Package main implements methods for saving the world.
//
// Experience has show that a small number of procedures
// can prove helpful when attempting to save the world.
//
// NOTE:
// Line comments are the norm
// block comments appear mostly as package comments, but are useful
// within an expression or to disable large swaths of code.
// Comments do not need extra formatting
package main

import (
	"fmt"
)

const c = "C"

var v int = 5

type T struct{}

func init() {
	if v != 5 && c != "C" {
		panic("fail to init")
	}

	// initialization of package
}

func main() {
	var a int
	Func1()
	// ...
	fmt.Println(a)
}

func (t T) Method1() {
	//...
}

func Func1() { // exported function Func1
	//...
}
