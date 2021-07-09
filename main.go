package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "greeter.h"
import "C"
import (
	"fmt"
	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
	"golang.org/x/image/colornames"
)

func GetFastKeyPoints(img gocv.Mat) []gocv.KeyPoint {
	ffd := gocv.NewFastFeatureDetector()
	defer ffd.Close()
	kp := ffd.Detect(img)
	return kp
}

func GetSIFTKeyPoints(img gocv.Mat) []gocv.KeyPoint {
	si := contrib.NewSIFT()
	defer si.Close()
	kp := si.Detect(img)
	//for i := 0 ; i < len(kp) ; i++ {
	//	fmt.Println(kp[i].Y)
	//}
	return kp
}



func main() {

	myResult := C.factorial(5)
	fmt.Println(myResult)

	//name := C.CString("Gopher")
	//defer C.free(unsafe.Pointer(name))
	//
	//year := C.int(2018)
	//
	//ptr := C.malloc(C.sizeof_char * 1024)
	//defer C.free(unsafe.Pointer(ptr))
	//
	//size := C.greet(name, year, (*C.char)(ptr))
	//
	//b := C.GoBytes(ptr, size)
	//fmt.Println("b basildi"+ string(b))

	webcam, _ := gocv.VideoCaptureDevice(1)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		//imgKeyPoint := gocv.NewMatWithSize(img.Rows(),img.Cols(),gocv.MatTypeCV8U)
		imgKeyPoint := img.Clone()
		myEdges := gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)
		gocv.Canny(img, &myEdges, 0, 100)

		keyPoints := GetSIFTKeyPoints(img)
		gocv.DrawKeyPoints(img, keyPoints, &imgKeyPoint, colornames.Gold, gocv.DrawDefault)
		window.IMShow(myEdges)
		window.WaitKey(1)
	}
}
