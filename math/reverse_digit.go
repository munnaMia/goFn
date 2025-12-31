package math

func ReverseDigit(n int) int {
	result := 0

	for n > 0 {
		result = (10 * result) + (n % 10)
		n /= 10
	}
return  result
}
