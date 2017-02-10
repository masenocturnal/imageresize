package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"flag"

	"github.com/rainycape/magick"
)

var source string
var dest string
var width int

func init() {

	flag.StringVar(&source, "src", "", "Directory containing files that we want to convert")
	flag.StringVar(&dest, "dst", "", "Directory we want to output to")
	flag.IntVar(&width, "width", 200, "Width to resize to")
}

func main() {
	//var dir = "/home/am/projects/sdk/images/gallery"
	fmt.Println("rar:", flag.Args())
	flag.Parse()

	fmt.Println("rar: ", source)
	fmt.Println("foo: ", dest)
	fmt.Println("tail:", flag.Args())

	// args := flag.Args()
	// fmt.Printf("rar %d", len(args))

	/*
		if len(args) < 1 {

			flag.Usage()
			os.Exit(2)
		}
	*/
	resizeFilesInDir(source, dest, width)
}

func resizeFilesInDir(dir string, dest string, size int) {
	filesInDir, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Print(err)
	}
	if len(filesInDir) < 1 {
		fmt.Print("Directory contains no files")
		os.Exit(2)
	}

	if size < 1 {
		fmt.Printf("Size : {size} must be > 0")
		os.Exit(2)
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
		newWidth = 200
		newHeight = 1
		layout = "landscape"
	}

	newImage, err := myImage.Resize(newWidth, newHeight, magick.FCubic)

	if err != nil {
		fmt.Printf("Could not resize %s ", fileName)
		fmt.Println(err)
	}

	outputFilename := path.Base(fileName)
	newFilename := path.Join("/tmp/images", layout, outputFilename)

	var imgWritten bool

	imgWritten, err = writeImage(newFilename, newImage)

	if !imgWritten {
		if err != nil {
			fmt.Print(err)
		}
	}

	fmt.Printf("File: %s has been resized to %d x %d \n", newFilename, newImage.Width(), newImage.Height())
}

func writeImage(newFilename string, newImage *magick.Image) (bool, error) {

	doesExist, err := exists(path.Base(newFilename))

	// @todo use case
	if err != nil {
		return false, err
	}

	if !doesExist {
		return false, nil
	}

	fo, err := os.Create(newFilename)

	if err != nil {
		return false, err
	}

	defer fo.Close()
	newImage.Encode(fo, magick.NewInfo())
	return true, nil
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
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
