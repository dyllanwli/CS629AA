package main

func validMountainArray(arr []int) bool {
	N := len(arr)
	i := 0
	// walk up
	for i+1 < N && arr[i] < arr[i+1] {
		i++
	}
	// peak cant be first of last
	if i == 0 || i == N-1 {
		return false
	}

	// walk down
	for i+1 < N && arr[i] > arr[i+1] {
		i++
	}
	return i == N-1
}

// func main() {
// 	arr := []int{3, 2, 1}
// 	r := validMountainArray(arr)
// 	fmt.Println(r)
// }
