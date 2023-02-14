package main

import "fmt"

const (
	pi = 3.14
	e  = 2.7
)

func main() {
	//这是单行注释
	/*
		多行注释
	*/

	var name string
	var age int
	var isOk bool
	var (
		a string
		b int
		c bool
		d float32
	)

	m := 100
	var name2, age2 = "123", 33
	fmt.Println("hello world")
	fmt.Println(name, age, isOk)
	fmt.Println(a, b, c, d)
	fmt.Println(name2, age2)
	fmt.Println(m)
	fmt.Println(pi, e)
}
