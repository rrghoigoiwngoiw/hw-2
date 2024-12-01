package binarysearch

func BinarySearch(input []int, search int) int {
	low, high := 0, len(input)-1
	for low <= high {
		mid := low + (high-low)/2
		if input[mid] == search {
			return mid
		}
		if input[mid] < search {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
