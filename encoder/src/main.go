package main

import (
	"fmt"
	"math"

	"fyne.io/fyne/v2/container"
)

func main() {
	var myApp Application
	myApp.CreateApp()
	myApp.CreateWindow("Application", NULL)

	const URL string = "https://youtube.com"
	encodedURL := encode(URL)
	fmt.Println(encodedURL)
	/*
		hashedURL, hashedToHex := hash(URL, WithSalt)
		fmt.Println(hashedURL)
		fmt.Println(hashedToHex)

		var db Database
		db.Initialize()
		db.StoreHash(hashedToHex, URL)
		defer db.Database.Close()
	*/

	var code Code
	gridSize := int(math.Sqrt(float64(FixedGridSize)))
	code = Code{
		gridSize,
		20,
		container.NewGridWithColumns(gridSize),
	}

	myApp.Window.DrawCode(code, encodedURL)
	myApp.Window.SetContent(code.Grid)
	myApp.Window.Open()
	myApp.Window.SavePNG("output.png")
	myApp.Run()
}
