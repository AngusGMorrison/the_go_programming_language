/*
Write a general-purpose unit-conversion program analogous to cf that reads
numbers from its command-line arguments for from the  standard input if there
are no arguments, and converts each number into units like temperature in
Celsius and Fahrenheit, length in feet and metres, weight in pounds and
kilograms, and the like.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"the_go_programming_language/ch2/ex2/convweight"
)

func main() {
	if len(os.Args) < 2 {
		convertFromInput()
	} else {
		convertFromArgs()
	}
}

func convertFromInput() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		f, err := strconv.ParseFloat(input.Text(), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex2: %v\n", err)
		}
		displayConversions(f)
	}
}

func convertFromArgs() {
	for _, arg := range os.Args[1:] {
		f, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex2: %v\n", err)
		}
		displayConversions(f)
	}
}

func displayConversions(f float64) {
	kg := convweight.Kg(f)
	lb := convweight.Lb(f)
	fmt.Printf("%s = %s, %s = %s\n",
		kg, convweight.KgToLb(kg), lb, convweight.LbToKg(lb))
}
