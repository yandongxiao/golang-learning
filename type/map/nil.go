// 缺点：如果key的value在map中确实存在，但是值为null. 如何区分？
// Java: m.get(key)). 如果key存在，则返回对应的value；否则返回null.
// python: m.get(key) 与java保持一致；访问形式如果为m[key], 则抛出异常
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	// NOTICE: 类似, golang当中，get永远不会抛出异常，如果该key不存在，返回value的一个zero-value。
	// Person 是一个值类型，它的zero-value是{"", 0}
	persons := make(map[string]Person)
	if persons["jack"].name == "" {
		println("因为struct是值类型，persons返回了一个zero-value")
	}

	// golang中区分zero-value和key的值不存在的方法
	if _, ok := persons["jack"]; !ok {
		println("map中不存在该条记录")
	}

	// range 接收nil值
	var kvs map[string]Person
	for k, v := range kvs {
		fmt.Println(k, v)
	}
}
