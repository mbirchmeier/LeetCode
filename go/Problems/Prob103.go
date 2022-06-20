package problems

import "fmt"

func isPowerOfThree(n int) bool {
	var test int = 1
	for test < n {
		test *= 3
	}
	return test == n
}

func countPrimes(n int) int {
	// var isPrime = make([]bool, n)

	// for index, element := range isPrime {
	fmt.Println(n)
	// }

	return 0

}
