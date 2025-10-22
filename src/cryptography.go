package main

import (
	"crypto/sha256"
	"math/rand"
	"strconv"
)

type Salt int

const (
	Null Salt = iota
	WithSalt
	WithoutSalt
)

func hash(url string, salt Salt) []int {
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
	return binary
}
