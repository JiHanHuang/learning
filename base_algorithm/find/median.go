package main

//https://leetcode.cn/problems/median-of-two-sorted-arrays/
//求第K个最小数，可以使用二分求解

import (
	"sort"
)

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	n1 := len(nums1)
// 	n2 := len(nums2)
// 	if n1+n2 == 0 {
// 		return 0
// 	}
// 	if n1 <= 1 || n2 <= 1 {
// 		merge := mergeArray(nums1, nums2)
// 		return middleN(merge)
// 	}
// 	//fmt.Println(">>>>>>:", n1, n2)
// 	mergeArray(nums1, nums2)
// 	start1 := 0
// 	end1 := n1 - 1
// 	start2 := 0
// 	end2 := n2 - 1
// 	for {
// 		if end1-start1 <= 1 || end2-start2 <= 1 {
// 			break
// 		}
// 		if middleN(nums1[start1:end1+1]) > middleN(nums2[start2:end2+1]) {
// 			e1 := (start1 + end1 + 1) / 2
// 			s2 := (start2 + end2) / 2
// 			min := end1 - e1
// 			if s2-start2 < min {
// 				min = s2 - start2
// 			}
// 			end1 = end1 - min
// 			start2 = start2 + min
// 		} else {
// 			e2 := (start2 + end2 + 1) / 2
// 			s1 := (start1 + end1) / 2
// 			min := end2 - e2
// 			if s1-start1 < min {
// 				min = s1 - start1
// 			}
// 			end2 = end2 - min
// 			start1 = start1 + min
// 		}
// 		//fmt.Println(start1, end1, start2, end2)
// 		mergeArray(nums1[start1:end1+1], nums2[start2:end2+1])
// 	}
// 	merge := mergeArray(nums1[start1:end1+1], nums2[start2:end2+1])
// 	return middleN(merge)
// }

// func middleN(a []int) float64 {
// 	L := len(a)
// 	if L == 0 {
// 		panic("nil array")
// 	}
// 	if L%2 == 0 {
// 		return float64(a[L/2]+a[(L-1)/2]) / 2
// 	}
// 	return float64(a[L/2])
// }

// func mergeArray(a, b []int) []int {
// 	//fmt.Println(a, b)
// 	totalLen := len(a) + len(b)
// 	merge := make([]int, totalLen)
// 	i := 0
// 	j := 0
// 	for k := 0; k < totalLen; k++ {
// 		if i >= len(a) {
// 			for ; j < len(b); j++ {
// 				merge[k] = b[j]
// 				k++
// 			}
// 			break
// 		}
// 		if j >= len(b) {
// 			for ; i < len(a); i++ {
// 				merge[k] = a[i]
// 				k++
// 			}
// 			break
// 		}
// 		if a[i] < b[j] {
// 			merge[k] = a[i]
// 			i++
// 		} else {
// 			merge[k] = b[j]
// 			j++
// 		}

// 	}
// 	//fmt.Println(merge)
// 	return merge
// }

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)
	if n1+n2 == 0 {
		return 0
	}
	if n1 <= 1 || n2 <= 1 {
		return getArrayMidle(nums1, nums2)
	}
	//fmt.Println(">>>>>>:", n1, n2)
	//mergeArray(nums1, nums2)
	start1 := 0
	end1 := n1 - 1
	start2 := 0
	end2 := n2 - 1
	for {
		if end1-start1 <= 1 || end2-start2 <= 1 {
			break
		}
		//m1 := (start1 + end1) / 2
		m1 := start1 + (end1-start1)>>1
		//m2 := (start2 + end2) / 2
		m2 := start2 + (end2-start2)>>1
		//if middleN(nums1[start1:end1+1]) > middleN(nums2[start2:end2+1]) {
		if nums1[m1] > nums2[m2] {
			//e1 := (start1 + end1 + 1) / 2
			e1 := start1 + (end1-start1+1)>>1
			min := end1 - e1
			if m2-start2 < min {
				min = m2 - start2
			}
			end1 = end1 - min
			start2 = start2 + min
		} else {
			//e2 := (start2 + end2 + 1) / 2
			e2 := start2 + (end2-start2+1)>>1
			min := end2 - e2
			if m1-start1 < min {
				min = m1 - start1
			}
			end2 = end2 - min
			start1 = start1 + min
		}
		//fmt.Println(start1, end1, start2, end2)
		//mergeArray(nums1[start1:end1+1], nums2[start2:end2+1])
	}
	//merge := mergeArray(nums1[start1:end1+1], nums2[start2:end2+1])
	//return middleN(merge)
	return getArrayMidle(nums1[start1:end1+1], nums2[start2:end2+1])
}

func middleN(a []int) float64 {
	L := len(a)
	if L == 0 {
		return 0
	}
	if L&1 == 0 {
		return float64(a[L>>1]+a[(L-1)>>1]) / 2
	}
	return float64(a[L>>1])
}

func mergeArray(a, b []int) []int {
	//fmt.Println(a, b)
	totalLen := len(a) + len(b)
	merge := make([]int, totalLen)
	i := 0
	j := 0
	for k := 0; k < totalLen; k++ {
		if i >= len(a) {
			for ; j < len(b); j++ {
				merge[k] = b[j]
				k++
			}
			break
		}
		if j >= len(b) {
			for ; i < len(a); i++ {
				merge[k] = a[i]
				k++
			}
			break
		}
		if a[i] < b[j] {
			merge[k] = a[i]
			i++
		} else {
			merge[k] = b[j]
			j++
		}

	}
	//fmt.Println(merge)
	return merge
}

func getArrayMidle(a, b []int) float64 {
	//fmt.Println(a, b)
	totalLen := len(a) + len(b)
	if totalLen == 0 {
		return 0
	}
	var large, small []int
	if len(a) > len(b) {
		large = a
		small = b
	} else {
		large = b
		small = a
	}
	if len(small) > 2 {
		panic("invlid input array")
	}
	L := len(large)
	if L < 2 {
		if len(small) < 1 {
			return float64(large[0])
		}
		return float64(large[0]+small[0]) / 2
	}

	var largeS int
	var largeE int
	if len(large) > 4 {
		if len(large)&1 == 0 {
			largeS = (L - 3) >> 1
			largeE = (L + 2) >> 1
		} else {
			largeS = L>>1 - 1
			largeE = (L + 1) >> 1
		}
	} else {
		largeS = 0
		largeE = len(large) - 1
	}
	//fmt.Println("midle:", large[largeS:largeE+1])
	result := mergeArray(large[largeS:largeE+1], small)
	//fmt.Println("sort:", result)
	return middleN(result)
}

//======================================================================
func findMedianSortedArraysOK(nums1 []int, nums2 []int) float64 {
	var merge []int
	merge = append(merge, nums1...)
	merge = append(merge, nums2...)
	sort.Ints(merge)
	//fmt.Println("Merge:", merge)
	L := len(merge)
	if L == 0 {
		return 0
	}
	if L%2 == 0 {
		return float64(merge[L/2]+merge[(L-1)/2]) / 2
	}
	return float64(merge[L/2])
}
