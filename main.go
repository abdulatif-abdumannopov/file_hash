package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"desktopApp/style"
	"encoding/hex"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/crypto/sha3"
	"hash"
	"image/color"
	"io"
	"log"
)

func main() {
	a := app.New()

	iconUri, err := storage.ParseURI("file://Icon.png")
	if err != nil {
		log.Fatal(err)
	}
	iconRes, err := storage.LoadResourceFromURI(iconUri)
	if err != nil {
		log.Fatal(err)
	}
	a.SetIcon(iconRes)

	w := a.NewWindow("Hash checker")
	w.Resize(fyne.NewSize(400, 600))

	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("File hash: ")
	output.Wrapping = fyne.TextWrapBreak
	output.Disable()

	hashSelect := widget.NewSelect([]string{"SHA256", "MD5", "SHA1", "SHA512", "SHA3-256", "SHA3-512"}, nil)
	hashSelect.SetSelected("SHA256")

	wrapped := container.NewMax(output)

	button := widget.NewButton("Выбрать файл", func() {
		dialog.ShowFileOpen(func(file fyne.URIReadCloser, err error) {
			if err != nil {
				log.Println("Ошибка:", err)
				return
			}
			if file == nil {
				return
			}
			defer file.Close()

			hasher := getMethod(hashSelect.Selected)
			if _, err := io.Copy(hasher, file); err != nil {
				output.SetText("Ошибка чтения файла")
				return
			}
			hash := hex.EncodeToString(hasher.Sum(nil))
			output.SetText(hash)
		}, w)
	})

	pad := canvas.NewRectangle(color.Transparent)
	pad.SetMinSize(fyne.NewSize(10, 20))

	w.SetContent(container.NewVBox(
		style.MarginWrap(widget.NewLabel("Hash method"), 0, 5, 0, 5),
		style.MarginWrap(hashSelect, 0, 5, 0, 5),
		style.MarginWrap(button, 0, 5, 0, 5),
		pad,
		container.NewVBox(
			layout.NewSpacer(),
			wrapped,
		),
	))

	w.ShowAndRun()
}

func getMethod(algorithm string) hash.Hash {
	switch algorithm {
	case "MD5":
		return md5.New()
	case "SHA1":
		return sha1.New()
	case "SHA512":
		return sha512.New()
	case "SHA3-256":
		return sha3.New256()
	case "SHA3-512":
		return sha3.New512()
	default:
		return sha256.New()
	}
}
