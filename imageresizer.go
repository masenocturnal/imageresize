package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

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
	flag.Parse()

	if len(source) < 1 {
		fmt.Println("Source directory must be provided")
		os.Exit(1)
	}

	if len(dest) < 1 {
		fmt.Println("Destination directory must be provided")
		os.Exit(1)
	}

	resizeFilesInDir(source, dest, width)
}

//
func resizeFilesInDir(dir string, dest string, size int) {
	filesInDir, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	if len(filesInDir) < 1 {
		fmt.Println("Directory contains no files")
		os.Exit(2)
	}

	if size < 1 {
		fmt.Printf("Size : %v must be > 0 ", size)
		os.Exit(2)
	}

	// ensure the destination dir exists
	dirExist, err := exists(path.Base(dest))
	if !dirExist {
		// try and create it
		err := os.MkdirAll(dest, 0500)
		if err != nil {
			fmt.Printf("Could not create dir %s", dest)
			os.Exit(2)
		}

	}

	for _, file := range filesInDir {
		fileName := file.Name()

		// @todo import mime/magic and compare the headers
		// with what image magic supports
		if strings.HasSuffix(strings.ToLower(fileName), ".jpg") {
			filePath := path.Join(dir, fileName)
			resizeFile(filePath, dest)
		} else {
			fmt.Printf("Skipping %s \n", fileName)
		}
	}

	fmt.Print("All Done")
	os.Exit(0)
}

func resizeFile(srcFile string, destDir string) {

	var myImage *magick.Image
	var err error

	myImage, err = magick.DecodeFile(srcFile)

	if err != nil {
		fmt.Println("fileName: {fileName}", srcFile)
		fmt.Println(err)
		os.Exit(1)
	}

	var isPortrait = IsPortrait(myImage)

	var newWidth int
	var newHeight int
	var layout string

	// @todo we can later on give
	// other resizing options
	if isPortrait {
		newWidth = 200
		newHeight = -1
		layout = "portrait"
	} else {
		newWidth = 200
		newHeight = -1
		layout = "landscape"
	}

	newImage, err := myImage.Resize(newWidth, newHeight, magick.FCubic)

	if err != nil {
		fmt.Printf("Could not resize %s ", srcFile)
		fmt.Println(err)
	}

	outputFilename := path.Base(srcFile)

	newFilename := path.Join(destDir, layout, outputFilename)

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

	fo, err := os.Create(newFilename)

	if err != nil {
		return false, errors.New("Path does not exist")
	}

	defer fo.Close()
	newImage.Encode(fo, magick.NewInfo())
	fmt.Println("File Written to : " + newFilename)
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
