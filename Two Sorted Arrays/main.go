package main

import (
	"log"
	"sort"
)

// https://leetcode.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	r := []int{}
	r = append(r, nums1...)
	r = append(r, nums2...)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return getMedian(r)
}

func getMedian(items []int) float64 {
	if len(items) == 0 {
		return 0
	}
	n := len(items)
	if n%2 == 0 {
		return float64(items[n/2]+items[(n/2)-1]) / 2
	}
	return float64(items[n/2])
}

func main() {
	a := []int{}
	b := []int{1}
	log.Println(findMedianSortedArrays(a, b))
}
