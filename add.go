package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("addition of two slices")
	num1 := []int{6, 7, 8}
	num2 := []int{6, 7, 8}
	Add(num1, num2)
}
func Add(num1 []int, num2 []int) {
	c := make([]int, len(num1))
	sort.Sort(sort.Reverse(sort.IntSlice(num2)))
	for i := 0; i < len(num1); i++ {
		c[i] = num1[i] + num2[i]
	}

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Printf("num1 + num2 : %v + %v = %v \n", num1, num2, c)
}
