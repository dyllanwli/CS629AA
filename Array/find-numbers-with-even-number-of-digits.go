package main

import (
	"fmt"
	"strconv"
)

func findNumbers(nums []int) int {
	res := 0
	for _, v := range nums {
		res += 1 - len(strconv.Itoa(v))%2
	}
	return res
}

func main() {
	fmt.Println(len(strconv.Itoa(12)) % 2)
}
