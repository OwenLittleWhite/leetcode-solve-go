package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	var concat = make([]int, len1+len2)
	i, j, k := 0, 0, 0
	for i < len1 && j < len2 {
		if nums1[i] < nums2[j] {
			concat[k] = nums1[i]
			i++
		} else {
			concat[k] = nums2[j]
			j++
		}
		k++
	}
	for i < len1 {
		concat[k] = nums1[i]
		k++
		i++
	}
	for j < len2 {
		concat[k] = nums2[j]
		k++
		j++
	}
	// 获取下标
	len3 := len(concat)
    fmt.Println(concat)
	if len3%2 > 0 {
		return float64(concat[len3/2])
	} else {
		index1 := len3 / 2
		index2 := index1 - 1
		return (float64(concat[index1]) + float64(concat[index2])) / 2
	}
}

func main() {
	a, b := []int{1, 2}, []int{3, 4}
	fmt.Println(findMedianSortedArrays(a, b))
}
