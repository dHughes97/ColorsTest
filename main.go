//**************************************************************
//  P03- "ASCII" Art Program
//
//   Name: Delton Hughes
//
//   Data Structures Date: 10/6/2023
//***************************************************************
//About: This program was to create other modules and import
// said modules in to a ColorsTest module. Each module with
// manipulate an image in a different way from changing the
// text to colored text. Also we had receive an image and download
// it, while also changing the image to gray scale.
//
//************************************************************************

package main

import (
	"github.com/dHughes97/Img_colors"
	"github.com/dHughes97/Img_get"
	"github.com/dHughes97/Img_gray_scale"
	"github.com/dHughes97/Img_text"
)

func main() {
	//Calling function to go get the image
	Img_get.ImageGet()
	//calling to print the colored text to image
	Img_text.Img_text()
	//changing the image to grayscale
	Img_gray_scale.Grayscale()
	//Is processing the pixel in the images
	Img_colors.ImgColors()
}
