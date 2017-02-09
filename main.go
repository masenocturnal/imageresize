package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/rainycape/magick"
)

func main() {
	var dir = "/home/am/projects/sdk/images/gallery"

	resizeFilesInDir(dir)
}

func resizeFilesInDir(dir string) {
	filesInDir, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Print(err)
	}
	if len(filesInDir) < 1 {
		fmt.Print("Directory contains no files")
		os.Exit(0)
	}

	for _, fileName := range filesInDir {
		file := path.Join(dir, fileName.Name())
		resizeFile(file)
	}

	fmt.Print("All Done")
	os.Exit(0)
}

func resizeFile(fileName string) {

	var myImage *magick.Image
	var err error

	myImage, err = magick.DecodeFile(fileName)

	if err != nil {
		fmt.Println("fileName: {fileName}", fileName)
		fmt.Println(err)
		os.Exit(1)
	}

	var isPortrait = IsPortrait(myImage)

	var newWidth int
	var newHeight int
	var layout string

	if isPortrait {
		newWidth = 200
		newHeight = -1
		layout = "portrait"
	} else {
		newWidth = -1
		newHeight = 200
		layout = "landscape"
	}

	newImage, err := myImage.Resize(newWidth, newHeight, magick.FCubic)

	if err != nil {
		fmt.Printf("Could not resize %s ", fileName)
		fmt.Println(err)
	}

	outputFilename := path.Base(fileName)
	newFilename := path.Join("/tmp/images", layout, outputFilename)

	writeImage(newFilename, newImage)

	fmt.Printf("File: %s has been resized to %d x %d \n", newFilename, newImage.Width(), newImage.Height())
}

func writeImage(newFilename string, newImage *magick.Image) {
	fo, err := os.Create(newFilename)
	if err != nil {
		log.Fatal(err)
	}

	defer fo.Close()
	newImage.Encode(fo, magick.NewInfo())
}

// ShowFormats displays all the formats supported
func ShowFormats() {
	var formats []string
	var err error
	formats, err = magick.SupportedFormats()

	if err != nil {
		fmt.Println(err)
	}

	for _, format := range formats {
		fmt.Println(format)
	}
}

// IsPortrait prints the current info for a given fileName
func IsPortrait(myImage *magick.Image) bool {

	var height = myImage.Height()
	var width = myImage.Width()

	fmt.Printf("Current Height is: %d \n", height)
	fmt.Printf("Current Width is: %d \n", width)

	return (height > width)
}
