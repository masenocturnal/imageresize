package main

import (
	//"bytes"
	//"image"
	//"image/gif"
	"fmt"
	"log"
	"os"

	"github.com/rainycape/magick"
)

func main() {

	var fileName = "/home/am/projects/sdk/images/gallery/DSC_0501.JPG"

	var myImage *magick.Image
	var err error

	myImage, err = magick.DecodeFile(fileName)

	if err != nil {
		fmt.Println("fileName: {0}", fileName)
		fmt.Println(err)
		os.Exit(1)
	}

	var isPortrait = IsPortrait(myImage)

	var newWidth = 200
	var newHeight = -1

	if isPortrait {

		// resize to width
		// var newImage *magick.Image

		newImage, err := myImage.Resize(newWidth, newHeight, magick.FCubic)

		if err != nil {
			fmt.Printf("Could not resize %s ", fileName)
			fmt.Println(err)
		}

		fo, err := os.Create("/tmp/output.jpg")
		if err != nil {
			log.Fatal(err)
			panic(err)

		}

		defer fo.Close()
		newImage.Encode(fo, magick.NewInfo())
	}
}

// ShowFormats displays all the formats supported
func ShowFormats() {

	/*
		var formats []string
		var err error
		formats, err = magick.SupportedFormats()

		if err != nil {
			fmt.Println(err)
		}
	*/
	/*
		for _, format := range formats {
			fmt.Println(format)
		}
	*/
}

// IsPortrait prints the current info for a given fileName
func IsPortrait(myImage *magick.Image) bool {

	var height = myImage.Height()
	var width = myImage.Width()

	fmt.Printf("height %d \n", height)
	fmt.Printf("width %d \n", width)

	return (height > width)
}
