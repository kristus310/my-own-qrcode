package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
)

const FixedGridSize = 256

type Salt int

const (
	Null Salt = iota
	WithSalt
	WithoutSalt
)

func hash(url string, salt Salt) ([]int, string) {
	var binary []int
	var hasher = sha256.New()

	if salt == WithSalt {
		randomSalt := strconv.Itoa(rand.Intn(len(url)))
		url += randomSalt
	}

	url = formattingURL(url)
	hasher.Write([]byte(url))
	hashed := hasher.Sum(nil)
	for i := 0; i < len(hashed); i++ {
		converted := intToBinary(hashed[i])
		for j := 0; j < len(converted); j++ {
			binary = append(binary, int(converted[j]))
		}
	}
	return binary, hex.EncodeToString(hashed)
}

func encode(url string) []int {
	encoded := url
	fmt.Println(encoded)

	firstPadding := true
	for len(encoded) < FixedGridSize/8 {
		if firstPadding {
			encoded += "!"
			firstPadding = false
		} else {
			encoded += "A"
		}
	}
	fmt.Println(encoded)
	fmt.Println(stringToBinary(encoded))
	return stringToBinary(encoded)
}
