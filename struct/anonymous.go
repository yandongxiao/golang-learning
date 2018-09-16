// Grouped globals
// NOTICE: 如果你需要使用map[T1]T2来聚合数据，考虑使用匿名类来代替
package main

import (
	"flag"
	"fmt"
)

type Person struct {
	Name string
}

// 在最外层定义了匿名类. 外部访问方式：package.Ps.Man
// 与数组方式相比，它不需要下标，语义更清晰
// 与MAP方式相比，它不需要额外定义key的取值范围，更安全
// 与在外部声明单独声明Man, Woman相比，Ps将同一类型的数据聚合在了一起
var Ps struct {
	Man, Woman Person
}

func init() {
	Ps.Man.Name = "jack"
	Ps.Woman.Name = "lili"
}

func main() {
	// 假设这是在另一个package中使用上面的全局变量
	flag.Parse()
	fmt.Println(Ps.Man)
	fmt.Println(Ps.Woman)
}
