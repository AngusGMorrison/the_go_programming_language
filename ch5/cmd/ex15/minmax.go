// Write variadic functions max and min, analogous to sum. What should these functions do when
// called wtih no arguments? Write variants that require at least one argument.
package main

import "fmt"

func main() {
	fmt.Println(min(2, 7, 4, 0))
	fmt.Println(max(8, 3, 5, 8))
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("min: no arguments supplied")
	}

	min := vals[0]
	for _, val := range vals[1:] {
		if val < min {
			min = val
		}
	}
	return min, nil
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("max: no arguments supplied")
	}

	max := vals[0]
	for _, val := range vals[1:] {
		if val > max {
			max = val
		}
	}
	return max, nil
}
