package binarysearch

func BinarySearch(data []int, value int) (result bool) {
	var mid int
	size := len(data)
	low := 0
	high := size - 1
	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == value {
			return true
		}
		if data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
