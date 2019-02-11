package main

import (
	"fmt"
)

func sum(nums ...int) {
	//fmt.Printf("%T\n", nums)
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	//slice
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	//array
	num := [4]int{1, 2, 3, 4}
	sum(num[:]...)
	//fmt.Printf("%T\n", num)
}
