package binarysearch

import "sort"

func BinarySearch(input []int, search int) bool {
	sort.Ints(input)
	low, high := 0, len(input)-1
	for low <= high {
		mid := low + (high-low)/2
		if input[mid] == search {
			return true
		}
		if input[mid] < search {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
