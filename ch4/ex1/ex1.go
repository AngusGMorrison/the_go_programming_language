/*
Write a function that counts the number of bits that are different in two SHA256 hashes. (See
PopCount from Section 2.6.2.)
*/

package main

import (
	"crypto/sha256"
	"fmt"
	"math/bits"
)

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	h1 := sha256.Sum256([]byte("Hello There"))
	h2 := sha256.Sum256([]byte("Hello there"))
	fmt.Println(popCount(h1, h2))
	fmt.Println(popCountWithStdLib(h1, h2))
}

func popCount(h1, h2 [32]byte) int {
	var popCount byte
	for i, b := range h1 {
		diff := b ^ h2[i]
		popCount += pc[diff]
	}
	return int(popCount)
}

func popCountWithStdLib(h1, h2 [32]byte) int {
	var popCount int
	for i, b := range h1 {
		popCount += bits.OnesCount8(b ^ h2[i])
	}
	return popCount
}
