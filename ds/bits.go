package ds

func KBitSet(n int, k int) bool {

	var temp int = (1 << (k - 1))

	return n&temp != 0
}

func CountSetBits(n int) int {

	res := 0
	for n > 0 {
		n = n & (n - 1)
		res++
	}

	return res
}
