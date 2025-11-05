package main

import (
	"image/color"
	"image/png"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

type Code struct {
	GridSize int
	TileSize float32
	Grid     *fyne.Container
}

type Window struct {
	Window fyne.Window
}

type Application struct {
	App    fyne.App
	Window Window
}

func (a *Application) CreateApp() {
	a.App = app.New()
}

// CreateWindow - For no/default icon, just put either "" or NULL
func (a *Application) CreateWindow(title string, iconPath string) {
	a.Window.Window = a.App.NewWindow(title)

	if iconPath != "" {
		res, err := fyne.LoadResourceFromPath(iconPath)
		checkError(err, "Loading icon from path", false)
		a.Window.Window.SetIcon(res)
	}
}

func (a *Application) Run() {
	a.App.Run()
}

func (w *Window) DrawCode(code Code, hashed []int) {
	for i := 0; i < code.GridSize*code.GridSize; i++ {
		rect := canvas.NewRectangle(color.Black)

		if i < len(hashed) && hashed[i] == 1 {
			rect.FillColor = color.White
		}

		rect.SetMinSize(fyne.NewSize(code.TileSize, code.TileSize))
		code.Grid.Add(rect)
	}
}

func (w *Window) SetContent(object fyne.CanvasObject) {
	w.Window.SetContent(object)
}

func (w *Window) Open() {
	w.Window.Show()
}

func (w *Window) SavePNG(fileName string) {
	var err error
	var file *os.File

	buildDirectory := ".build/"
	err = createDirectory(buildDirectory)
	checkError(err, "Creating the build directory", true)

	file, err = createFile(buildDirectory + fileName)
	checkError(err, "Creating the image file", false)
	defer closeFile(file)

	image := w.Window.Canvas().Capture()
	err = png.Encode(file, image)
	checkError(err, "Saving the image", false)
}
