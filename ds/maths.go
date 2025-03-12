package ds

func CountDigits(x int) int {
	var res int

	for x > 0 {
		x = x / 10
		res++
	}

	return res

}
func PalindromeNumber(x int) bool {
	var rev int = 0
	var temp int = x

	for temp > 0 {
		digit := temp % 10

		rev = rev*10 + digit
		temp = temp / 10
	}

	return x == rev
}
func IterFact(x int) int {
	res := 1

	for i := 2; i <= 2; i++ {
		res = res * i
	}
	return res
}

func Factorial(x int) int {

	if x == 0 || x == 1 {
		return 1
	}
	return x * Factorial(x-1)
}

func FactorialZeros(x int) int {
	var res int = 0

	for i := 5; i <= x; i *= 5 {
		res = res + x/i
	}
	return res
}
func FactorialZeros2(x int) int {
	var res int = 0

	for x > 0 {
		res += x / 5
		x = x / 5
	}

	return res
}

func GCD(a int, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(a int, b int) int {
	return (a * b) / GCD(a, b)
}

func CheckPrime(x int) bool {

	if x == 1 {
		return false
	}
	if x == 2 || x == 3 {
		return true
	}
	if x%2 == 0 || x%3 == 0 {
		return false
	}

	for i := 5; i*i <= x; i += 6 {
		if x%i == 0 || x%(i+2) == 0 {
			return false
		}
	}
	return true
}
