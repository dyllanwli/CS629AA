package main

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 && i+1 < len(arr) {
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			// use ... to append another slice
			i++
		}
	}
}

// func main() {
// 	arr := []int{1, 2, 3, 0, 4, 5, 0}
// 	duplicateZeros(arr)
// 	fmt.Println(arr)
// }
