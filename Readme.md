# Image Resize

This is very basic application that uses image magick to iterate through 
a directory of images and resize them proportionally to a given width. 
Files will be placed in directories according to thier orientation



```
$ go build imageresizer.go 
$ ./imageresizer -src="/home/am/projects/sdk/images/gallery" -dst="/tmp/" -width=200 



$ tree /tmp/images/
├── landscape
│   ├── 005.JPG
│   ├── 006.JPG
│   ├── 016.JPG
│   ├── 022.JPG
│   ├── 032.JPG
│   ├── DSC_0468.JPG
│   ├── DSC_0526.JPG
│   ├── DSC_0951 (Small).JPG
│   ├── DSC_0960 (Small).JPG
│   ├── DSC_0987.JPG
│   ├── DSCN1970.JPG
│   ├── IMG_0757.JPG
│   ├── IMG_4121.JPG
│   ├── IMG_4146.JPG
│   └── IMG_4354.JPG
└── portrait
    ├── 001.JPG
    ├── 004.JPG
    ├── 015.JPG
    ├── 017.JPG
    ├── 018.JPG
    ├── 021.JPG
    ├── 025.JPG
    ├── 026.JPG
    ├── 037 (2).JPG
    ├── 038.JPG
    ├── 043.JPG
    ├── 046.JPG
    ├── 048.JPG
    ├── 062.JPG
    ├── 069.JPG
    ├── DSC_0501.JPG
    ├── DSC_0580.JPG
    ├── DSC_0931.JPG
    ├── DSCN0768.JPG
    ├── IMG_0167.JPG
    ├── IMG_0182.JPG
    ├── IMG_2669.JPG
    └── IMG_3280.JPG


```
