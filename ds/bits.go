package ds

func KBitSet(n int, k int) bool {

	var temp int = (1 << k)

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

func PowerOfTwo(n int) bool {
	res := 0

	for n > 0 {
		n = n & (n - 1)
		res++
	}

	return res == 1
}
func PowerOfTwo2(n int) bool {
	if n == 0 {
		return false
	}
	return n&(n-1) == 0
}

func FindOdd(a []int, n int) int {
	res := a[0]

	for i := 1; i < n; i++ {
		res = res ^ a[i]
	}

	return res
}
