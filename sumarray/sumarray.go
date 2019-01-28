package sumarray

func Sumarray(data []int) (total int) {
	total = 0
	for _, v := range data {
		total = total + v
	}
	return total
}
