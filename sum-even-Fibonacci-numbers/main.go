package main

func SumEvenFibonacci(limit int) (result int) {

	return result
}

func findFibonacciN(n int) (fibonacciN int) {

	if (n == 1) || (n == 2) {
		return 1
	} else if n == 0 {
		return 0

	}

	return findFibonacciN(n-1) + findFibonacciN(n-2)
}
