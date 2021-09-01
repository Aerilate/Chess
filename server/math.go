package main

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	return a + b - min(a, b)
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func mag(n int) int {
	if n == 0 {
		return 0
	}
	return n / abs(n)
}
