package main

import (
	"fmt"
	"sort"
	// "strings"
)

// var (
// 	a int    = 20
// 	b int    = 30
// 	c string = "liudan"
// 	d string = "xueanguo"
// )

// const (
// 	q, w = iota + 1, iota + 2
// 	e, r
// 	t, y
// )

func main() {
	// fmt.Println("Hello World")
	// fmt.Println(a, b)

	// fmt.Println(a + b)
	// fmt.Println(a - b)

	// fmt.Printf(c + d + "\n")

	// var x = [...]bool{false, true, false}
	// for _, v := range x {
	// 	fmt.Println(v)
	// }

	// var str = []string{"geeks", "for", "geeks"}

	// var scoreMap int = 20

	// fmt.Println(len(c))
	// fmt.Println(strings.Split(c, ""))
	// fmt.Println(strings.Contains(c, "l"))
	// fmt.Println(strings.HasPrefix(c, "l"))
	// fmt.Println(strings.LastIndex(c, "l"))
	// fmt.Println(strings.Join(str, "-"))

	// if scoreMap > 90 {
	// 	fmt.Println("A")
	// } else if scoreMap > 75 && scoreMap < 90 {
	// 	fmt.Println("B")
	// } else {
	// 	fmt.Println("C")
	// }

	// fmt.Println(q, w, e, r, t, y)

	// switch scoreMap := 20; scoreMap {
	// case 20:
	// 	fmt.Println("1")
	// case 30:
	// 	fmt.Println("2")
	// default:
	// 	fmt.Println("xxx")
	// }

	// o := [5]int{1, 2, 3, 4, 5}
	// s := o[1:3]
	// fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	// a := make([]int, 2, 10)
	// fmt.Println(a)
	// fmt.Println(len(a))
	// fmt.Println((cap(a)))

	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"])
	fmt.Printf("type of a:%T\n", scoreMap)

	v, ok := scoreMap["小王子"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	userInfo := map[string]string{
		"username": "沙河小王子",
		"password": "123456",
	}

	for k, v := range userInfo {
		fmt.Println(k, v)
	}

	var x = [...]int{3, 7, 8, 9, 1}
	sort.Ints(x[:])
	fmt.Println(x)

	var y = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		y = append(y, fmt.Sprintf("%v", i))
	}
	fmt.Println(y)

}
