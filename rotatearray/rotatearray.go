package rotatearray

func RotateArray(data []int, value int) (result []int) {
	size := len(data)
	if size < value {
		return data
	}
	r := size - value%size
	data = append(data[r:], data[:r]...)
	return data
}
