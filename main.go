package main

import (
	"crypto/sha256"
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"unsafe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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

func stringToBinary(str string) []int {
	var binary []int
	for i := 0; i < len(str); i++ {
		converted := intToBinary(str[i])
		for j := 0; j < len(converted); j++ {
			binary = append(binary, int(converted[j]))
		}
	}
	return binary
}

func hashSalt(url string) []int {
	var binary []int
	hasher := sha256.New()
	salt := rand.Intn(len(url))

	hasher.Write([]byte(url + strconv.Itoa(salt)))
	hashed := hasher.Sum(nil)
	fmt.Println(hashed)
	for i := 0; i < len(hashed); i++ {
		converted := intToBinary(hashed[i])
		for j := 0; j < len(converted); j++ {
			binary = append(binary, int(converted[j]))
		}
	}
	return binary
}

func main() {
	url := "youtube.com"
	binary := stringToBinary(url)
	final := arrayToString(binary)
	hashed := hashSalt(final)

	test := arrayToString(hashed)
	fmt.Println(test)

	a := app.New()
	w := a.NewWindow("Hello World")

	fmt.Println(len(hashed))
	gridSize := int(math.Ceil(math.Sqrt(float64(len(hashed)))))
	fmt.Println(gridSize)
	grid := container.NewGridWithColumns(gridSize)

	for i := 0; i < gridSize*gridSize; i++ {
		rect := canvas.NewRectangle(color.Black)

		if i < len(hashed) && hashed[i] == 1 {
			rect.FillColor = color.White
		}

		rect.SetMinSize(fyne.NewSize(20, 20))
		grid.Add(rect)
	}

	w.SetContent(grid)
	w.ShowAndRun()
}
