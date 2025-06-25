package main

import (
	"fmt"
	"time"
)

func main() {
	var a = 2
	var arr = [...]int32{1, 2, 3, 4}
	fmt.Println(arr[3], a)
	arr2 := arr[0:2]
	fmt.Println(arr2)
	arr2[0] = 111111111
	fmt.Print(arr)
	var slice []int = []int{4, 5, 6}
	slice = append(slice, 6)
	fmt.Println(slice)
	fmt.Printf("\n the l is %v and capacity is %v\n", len(slice), cap(slice))
	s := make([]int, 4, 6)
	s1 := []int{1, 2, 3}
	fmt.Println(s1[0])
	s1 = []int{4, 5, 6}
	fmt.Println(s1[0])
	s[0] = 1
	fmt.Println(s[0])
	//map {key: value} map[string]int32
	// var m map[string]int32
	scores := map[string]int32{
		"alice": 91,
		"jk":    27,
	}
	fmt.Printf("%#v %% %#v\n", scores["alice"], scores["jk"])
	//now if we have arrays we need loops
	nums := []int{1, 2, 3}
	for i, v := range nums {
		fmt.Println(i, v)
	}
	for name, score := range scores {
		fmt.Printf("Name: %v Score: %v ", name, score)
	}
	//while
	var i int = 0
	for i < 6 {
		fmt.Println(i)
		i = i + 1
	}

}

func TimeLoop(slice []int, n int) time.Duration {
	t0 := time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
