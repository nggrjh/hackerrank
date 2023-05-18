package common

func Factorial(n float64) float64 {
	if n < 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Combinations(n, r float64) float64 {
	return Factorial(n) / (Factorial(r) * Factorial(n-r))
}
