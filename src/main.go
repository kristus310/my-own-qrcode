package main

import (
	"math"

	"fyne.io/fyne/v2/container"
)

func main() {
	var myApp Application
	myApp.CreateApp()
	myApp.CreateWindow("Application", NULL)

	const URL string = "prdelprdelprdelprdelprdelprdela"
	hashedURL := hash(URL, WithoutSalt)

	//encodedURL := encode(URL)
	//fmt.Println(encodedURL)

	var code Code
	gridSize := int(math.Sqrt(float64(FixedGridSize)))
	code = Code{
		gridSize,
		20,
		container.NewGridWithColumns(gridSize),
	}

	myApp.Window.DrawCode(code, hashedURL)
	myApp.Window.SetContent(code.Grid)
	myApp.Window.Open()
	//myApp.Window.SavePNG("build/output.png")
	myApp.Run()
}
