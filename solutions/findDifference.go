package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	res := make([][]int, 2)
	n1 := len(nums1)
	n2 := len(nums2)

	num1 := make(map[int]bool, n1)
	num2 := make(map[int]bool, n2)

	for _, i := range nums1 {
		num1[i] = true
	}
	for _, i := range nums2 {
		num2[i] = true
	}

	// temp := []int{}
	// for _, i := range nums2 {
	// 	//fmt.Printf("Map i:%v, v: %d\n", i, v)
	// 	if !num1[nums2[i]] {
	// 		temp = append(temp, nums2[i])
	// 	}
	// }
	res[0] = check(num1, num2)
	res[1] = check(num2, num1)

	return res
}

func check(a, b map[int]bool) []int {
	ans := make([]int, 0)
	for val, _ := range a {
		if _, ok := b[val]; !ok {
			ans = append(ans, val)
		}
	}
	return ans
}
