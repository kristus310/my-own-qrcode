package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decode(encodedBinary []int) {
	var url string
	var bytis [8]byte
	var key int
	length := len(encodedBinary)
	padding := true

	for i := 0; i < length/8; i++ {
		for j := 0; j < 8; j++ {
			bytis[j] = byte(encodedBinary[i*8+j])
		}
		num := binaryToInt(bytis)

		if padding {
			fmt.Println(string(num))
			url += string(num)

			if string(num) == "!" {
				padding = false
				keyString, _ := strings.CutSuffix(url, "!")
				key, _ = strconv.Atoi(keyString)
				fmt.Println("KEYYY!!!")
				fmt.Println(keyString)
				fmt.Println(key)
			}
		} else {
			fmt.Println(string(num / uint8(i+1)))
			url += string(num ^ uint8(i*key))
		}
	}
	fmt.Println(encodedBinary)
	fmt.Println(url)
}
