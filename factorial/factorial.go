package factorial

func Factorial(i int) (result int) {
	if i <= 1 {
		return 1
	}
	return i * Factorial(i-1)
}
