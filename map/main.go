package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	// var scoreMap = make(map[string]int, 8)
	// scoreMap["刘丹"] = 100
	// scoreMap["薛安果"] = 130

	// for k, v := range scoreMap {
	// 	fmt.Println(k, v)
	// }

	// delete(scoreMap, "薛安果")
	// fmt.Println(scoreMap)

	var scoreMap = make(map[string]int, 100)
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	keys := make([]string, 0, 100)
	for k := range scoreMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

	var mapSlice = make([]map[string]int, 8, 8)
	mapSlice[0] = make(map[string]int)
	mapSlice[0]["name"] = 100
	fmt.Println(mapSlice)

	var sliceMap = make(map[string][]int, 8)
	sliceMap["中国"] = make([]int, 8)
	sliceMap["中国"][0] = 100
	sliceMap["中国"][1] = 200
	sliceMap["中国"][2] = 300

	sliceMap["日本"] = make([]int, 10)
	sliceMap["日本"][0] = 0
	sliceMap["日本"][1] = 1

	for k, v := range sliceMap {
		fmt.Println(k, v)
	}

	var s = "how do you do"
	var wordCount = make(map[string]int, 10)
	var words = strings.Split(s, "")
	for _, word := range words {
		k, ok := wordCount[word]
		if ok {
			wordCount[word] = k + 1
		} else {
			wordCount[word] = 1
		}
	}

	for k, v := range wordCount {
		if k == " " {
			continue
		}
		fmt.Println(k, v)
	}
}
