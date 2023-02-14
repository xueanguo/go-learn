package main

import (
	"fmt"
	// "sort"
)

// func main() {
// 	a := 2
// 	b := 5

// 	fmt.Println(a + b)
// 	fmt.Println(a - b)
// 	fmt.Println(a * b)
// 	fmt.Println(a / b)
// 	fmt.Println(a % b)
// }

// func main() {
// 	var score = 65
// 	if score >= 90 {
// 		fmt.Println("A")
// 	} else if score > 75 {
// 		fmt.Println("B")
// 	} else {
// 		fmt.Println("C")
// 	}
// }

// func main() {
// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println(i)
// 	// }
// 	for i := 0; i < 5; i++ {

// 		if i == 3 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	finger := 3
// 	switch finger {
// 	case 1:
// 		fmt.Println("1")
// 	case 2:
// 		fmt.Println("2")
// 	case 3:
// 		fmt.Println("3")
// 	case 4:
// 		fmt.Println("4")
// 	default:
// 		fmt.Println("无效输入")
// 	}
// }

// func main() {
// 	age := 30
// 	switch {
// 	case age < 25:
// 		fmt.Println("haohaoxuexi")
// 	case age > 25 && age < 35:
// 		fmt.Println("haohaogongzuo")
// 	case age > 60:
// 		fmt.Println("hhh")
// 	}
// }

// func main() {
// 	var a [3]int
// 	var b [4]int

// 	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
// 	var boolArray = [...]bool{true, false, false, true, false}

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(cityArray)
// 	fmt.Println(boolArray)
// }

// func main() {
// 	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
// 	for i := 0; i < len(cityArray); i++ {
// 		fmt.Println(cityArray[i])
// 	}

// 	for _, value := range cityArray {
// 		fmt.Println(value)
// 	}
// }

// func main() {
// 	cityArray := [4][2]string{
// 		{"北京", "西安"},
// 		{"上海", "杭州"},
// 		{"重庆", "成都"},
// 		{"广州", "深圳"},
// 	}
// 	for _, v1 := range cityArray {
// 		for _, v2 := range v1 {
// 			fmt.Println(v2)
// 		}
// 	}
// }

//切片

// func main() {
// 	// a := [5]int{55, 56, 57, 58, 59}
// 	// b := a[1:4]
// 	// fmt.Println(b)
// 	// fmt.Printf("%T\n ", b)

// 	var a []int
// 	a = append(a, 10)
// 	fmt.Println(a, len(a), cap(a))
// 	b := []int{11, 12, 13, 14}
// 	a = append(a, b...)
// 	fmt.Println(a)
// }

// func main() {
// 	var a = make([]string, 5, 10)
// 	for i := 0; i < 10; i++ {
// 		a = append(a, fmt.Sprintf("%v", i))
// 	}
// 	fmt.Println(a)
// 	fmt.Println(len(a))
// }

// func main() {
// 	var a = [...]int{3, 7, 8, 9, 1}
// 	sort.Ints(a[:])
// 	fmt.Println(a)

// 	userInfo := map[string]string{
// 		"username": "沙河小王子",
// 		"pasword":  "123456",
// 	}
// 	fmt.Println(userInfo)
// }

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v, ok)
	} else {
		fmt.Println("查无此人")
	}
}
