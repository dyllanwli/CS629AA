package main

import "fmt"

func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for _, m := range arr[i+1:] {
			if (arr[i] == m*2) || (arr[i]*2 == m) {
				return true
			}
		}
	}
	return false
}

func checkIfExist2(arr []int) bool {
	// hash table solution
	dummy := make(map[int]bool, len(arr))
	for _, n := range arr {
		fmt.Println(n)
		// check previous stored element is can be assignment or not
		if ok := dummy[2*n]; ok {
			return true
		}

		// if even, we can be doubled value ourself
		if n%2 == 0 {
			if ok := dummy[n/2]; ok {
				return true
			}
		}
		dummy[n] = true
	}
	return false
}

// func main() {
// 	arr := []int{-20, 8, -6, -14, 0, -19, 14, 4}
// 	fmt.Println(checkIfExist2(arr))
// }
