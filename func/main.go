package main

import "fmt"

func sayHello(name string) {
	fmt.Println("hello", name)
}

func intSum(a int, b int) int {
	ret := a + b
	return ret
}

func intsum2(a ...int) int {
	ret := 0
	for _, v := range a {
		ret = ret + v
	}
	return ret
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc(x int, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	sayHello("xueanguo")

	r := intSum(10, 20)
	fmt.Println(r)
	x := intsum2(20, 20)
	fmt.Println(x)

	r1 := calc(1, 2, add)
	r2 := calc(1, 2, sub)

	fmt.Println(r1, r2)
}
