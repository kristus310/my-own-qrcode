package main

import (
	"fmt"
	"strings"
	"unsafe"
)

type Randomize struct {
	Randomized bool
	Seed       int
}

// Only for mine and testing purposes
func sizeOf[T any](x T) {
	y := unsafe.Sizeof(x)
	fmt.Printf("Bits: %d & bytes: %d", y*8, y)
}

func reverseByte(bytes [8]byte) [8]byte {
	var reversed [8]byte
	for i := 0; i < 8; i++ {
		reversed[i] = bytes[7-i]
	}
	return reversed
}

func arrayToString(array any) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), ""), "[]")
}

func intToBinary(num uint8) [8]byte {
	var bytes [8]byte
	var bit uint8
	for i := 0; i < 8; i++ {
		bit = num % 2
		bytes[i] = bit
		num = num / 2
	}
	return reverseByte(bytes)
}

func binaryToInt(bits [8]byte) uint8 {
	var num uint8 = 0
	for i := 0; i < 8; i++ {
		num = num*2 + bits[i]
	}
	return num
}

func stringToBinary(str string, random Randomize) []int {
	var binary []int
	var converted [8]byte

	fmt.Println(strings.Index(str, "!"))
	for i := 0; i < len(str); i++ {
		if random.Randomized && i > strings.Index(str, "!") {
			converted = intToBinary(str[i] ^ uint8(i*random.Seed))
		} else {
			converted = intToBinary(str[i])
		}
		for j := 0; j < len(converted); j++ {
			binary = append(binary, int(converted[j]))
		}
	}
	return binary
}

func formattingURL(url string) string {
	return arrayToString(stringToBinary(url, Randomize{false, 0}))
}
