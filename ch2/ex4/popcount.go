/*
Rewrite PopCount to use a loop instead of a single expression. Compare the
performance of the two versions.
*/

package popcount

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountShift(x uint64) int {
	count := 0
	for x != 0 {
		if x&1 == 1 {
			count++
		}
		x = x >> 1
	}
	return count
}
