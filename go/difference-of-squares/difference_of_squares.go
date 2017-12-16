package diffsquares

func SquareOfSums(n int) int {
	sumN := n * (n + 1) / 2
	return sumN * sumN
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
