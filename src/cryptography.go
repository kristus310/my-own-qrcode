package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

const FixedGridSize = 576

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
	https := "https://"
	http := "http://"
	if strings.HasPrefix(url, https) {
		url = url[len(https):]
	} else if strings.HasPrefix(url, http) {
		url = url[len(http):]
	}

	firstPadding := true
	for len(url) < FixedGridSize/8 {
		if firstPadding {
			url += "!" + strconv.Itoa(len(url)) + "!"
			firstPadding = false
		} else {
			url += "A"
		}
	}
	binary := stringToBinary(url, false)
	fmt.Println(url)
	fmt.Println(binary)
	return binary
}
