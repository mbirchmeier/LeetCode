package problems

func isPowerOfThree(n int) bool {
	var test int = 1
	for test < n {
		test *= 3
	}
	return test == n
}
