package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		// spacial case
		copy(nums1, nums2)
	}

	i, j, k := m-1, n-1, m+n-1
	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			// start from the end
			// so we can avoid append slice operation
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
	for ; j >= 0; k-- {
		nums1[k] = nums2[j]
		j--
	}
	// fmt.Println(nums1)

}

// func main() {
// 	nums, m, nums2, n := []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3
// 	merge(nums, m, nums2, n)
// }
