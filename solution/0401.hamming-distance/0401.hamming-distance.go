package _401_hamming_distance

import "math/bits"

func hammingDistance(x int, y int) int {
	xor := x ^ y
	return bits.OnesCount(uint(xor))
}

//func hammingDistance(x int, y int) int {
//	xor := x ^ y
//	dis := 0
//	for xor != 0 {
//		if xor & 1 == 1 {
//			dis += 1
//		}
//		xor = xor >> 1
//	}
//	return dis
//}

//func hammingDistance(x int, y int) int {
//	xor := x ^ y
//	dis := 0
//	for xor != 0 {
//		dis += 1
//		xor = xor & (xor - 1)
//	}
//	return dis
//}
