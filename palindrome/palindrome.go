package palindrome

import "algorithms/stringreversal"

func Palindrome(s string) (result bool) {
	return s == stringreversal.Reverse(s)
}