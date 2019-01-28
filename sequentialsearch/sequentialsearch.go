package sequentialsearch

func SequentialSearch(data []int, value int) (result bool) {
	for _, v := range data {
		if v == value {
			return true
		}
	}
	return false
}
