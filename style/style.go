package style

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

func HSidePadding(inner fyne.CanvasObject, padSize float32) fyne.CanvasObject {
	leftPad := canvas.NewRectangle(color.Transparent)
	leftPad.SetMinSize(fyne.NewSize(padSize, 1))

	rightPad := canvas.NewRectangle(color.Transparent)
	rightPad.SetMinSize(fyne.NewSize(padSize, 1))

	return container.NewHBox(
		leftPad,
		inner,
		rightPad,
	)
}

func MarginWrap(obj fyne.CanvasObject, top, right, bottom, left float32) fyne.CanvasObject {
	topPad := canvas.NewRectangle(color.Transparent)
	topPad.SetMinSize(fyne.NewSize(1, top))

	bottomPad := canvas.NewRectangle(color.Transparent)
	bottomPad.SetMinSize(fyne.NewSize(1, bottom))

	leftPad := canvas.NewRectangle(color.Transparent)
	leftPad.SetMinSize(fyne.NewSize(left, 1))

	rightPad := canvas.NewRectangle(color.Transparent)
	rightPad.SetMinSize(fyne.NewSize(right, 1))

	return container.NewVBox(
		topPad,
		container.NewHBox(
			leftPad,
			container.NewMax(obj),
			rightPad,
		),
		bottomPad,
	)
}
