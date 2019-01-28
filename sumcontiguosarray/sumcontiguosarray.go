package sumcontiguosarray
// given a list of positive and negative intergers, find a contigous subarray whos sum
// (sum of elements) is maximun

func MaxSubArraySum(data []int) (total int) {
	maxSoFar := 0
	maxEndingHere := 0
	for _, v := range data {
		//fmt.Printf("max_endingbefore::%d,data::%d\n", maxEndingHere, v)
		maxEndingHere = maxEndingHere + v
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
		//fmt.Printf("max_endingafter::%d,index::%d\n", maxEndingHere, k)
		//fmt.Printf("max_sofar::%d\n", maxSoFar)
	}
	return maxSoFar
}
