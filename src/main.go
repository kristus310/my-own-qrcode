package main

import (
	"math"

	"fyne.io/fyne/v2/container"
)

func main() {
	var myApp Application
	myApp.CreateApp()
	myApp.CreateWindow("Application", NULL)

	const URL string = "youtube.com"
	hashedURL := hash(URL, WithSalt)

	var code Code
	gridSize := int(math.Ceil(math.Sqrt(float64(len(hashedURL)))))
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
