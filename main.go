package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"unsafe"
)

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

func hashSalt(url string) string {
	hasher := sha256.New()
	salt := rand.Intn(len(url) - (2 * len(url)) + (2 * len(url)))
	hasher.Write([]byte(url + strconv.Itoa(salt)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	var binary []int

	word := "aaa"
	for i := 0; i < len(word); i++ {
		converted := intToBinary(word[i])
		for j := 0; j < len(converted); j++ {
			binary = append(binary, int(converted[j]))
		}
	}

	final := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(binary)), ""), "[]")

	hashed := hashSalt(final)
	fmt.Println(hashed)

	/*

	 */
	/*
		a := app.New()
		w := a.NewWindow("Hello World")

		gridSize := int(math.Ceil(math.Sqrt(float64(len(binary)))))
		grid := container.NewGridWithColumns(gridSize)

		for i := 0; i < gridSize*gridSize; i++ {
			rect := canvas.NewRectangle(color.Black)

			if i < len(binary) && binary[i] == 1 {
				rect.FillColor = color.White
			}

			rect.SetMinSize(fyne.NewSize(20, 20))
			grid.Add(rect)
		}

		w.SetContent(grid)
		w.ShowAndRun()
	*/
}
