package main

import (
	"fmt"
	// "math"
)

const (
// pi = 3.14
// e  = 2.7
)

func main() {
	//这是单行注释
	/*
		多行注释
	*/

	// var name string
	// var age int
	// var isOk bool
	// var (
	// 	a string
	// 	b int
	// 	c bool
	// 	d float32
	// )

	// m := 100
	// var name2, age2 = "123", 33
	// fmt.Println("hello world")
	// fmt.Println(name, age, isOk)
	// fmt.Println(a, b, c, d)
	// fmt.Println(name2, age2)
	// fmt.Println(m)
	// fmt.Println(pi, e)
	// var a int = 10
	// fmt.Printf("%d \n", a) // 10
	// fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制
	// fmt.Println(math.MaxFloat32)

	// s1 := "您好"
	// s2 := "hello"
	// fmt.Println(s1, s2)
	// fmt.Println(s1 + s2)
	// s3 := fmt.Sprintf("%s%s", s1, s2)
	// fmt.Println(s3)
	s := "hello沙河"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}

	for _, r := range s {
		fmt.Printf("%c\n", r)
	}

	s1 := "big"
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

}
