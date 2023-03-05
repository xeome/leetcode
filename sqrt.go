package main

// Sqrt returns an approximation to the square root of x. Using Newton's method.
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}
