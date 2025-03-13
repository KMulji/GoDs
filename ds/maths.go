package ds

import "fmt"

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

func PrimeFactors(x int) {
	if x <= 1 {
		return
	}

	for x%2 == 0 {
		fmt.Printf("%d ", 2)
		x = x / 2
	}

	for x%3 == 0 {
		fmt.Printf("%d ", 3)
		x = x / 3
	}
	for i := 5; i <= x*x; i = i + 6 {
		for x%i == 0 {
			fmt.Printf(" %d", i)
			x = x / i
		}

		for x%(i+2) == 0 {
			fmt.Printf(" %d", i+2)
			x = x / (i + 2)
		}

	}
	fmt.Printf("\n")
	if x > 3 {
		fmt.Printf(" %d", x)
	}
}

func AllDivisors(x int) {

	var i int = 1

	for ; i*i <= x; i++ {
		if x%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	for ; i >= 1; i-- {
		if x%i == 0 {
			fmt.Printf("%d ", x/i)
		}
	}
	fmt.Printf("\n")
}

func Siev(x int) {
	isPrime := make([]bool, x+1)

	for i, _ := range isPrime {
		isPrime[i] = true
	}

	for i := 2; i*i <= x; i++ {
		if isPrime[i] {
			for j := i * 2; j <= x; j = j + i {
				isPrime[j] = false
			}
		}
	}
	for i := 2; i <= x; i++ {
		if isPrime[i] {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")

}

func Power(x int, n int) int {
	if n == 0 {
		return 1
	}

	var temp int = Power(x, n/2)
	temp = temp * temp

	if n%2 == 0 {
		return temp
	}
	return temp * x
}

func IterPower(x int, n int) int {
	var res int = 1
	for n > 0 {
		//bit is 1
		if n%2 != 0 {
			res = res * x

		}
		x = x * x
		n = n / 2
	}

	return res
}
