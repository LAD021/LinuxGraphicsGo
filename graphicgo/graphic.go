package graphicgo

import (
	"fmt"
	"os"
)

var dev *os.File
var bgColor = BLACK
var graphColor = WHITE
var fontColor = WHITE

const screenSize = screenHeight * screenWidth * 3

var backgroundBuff [screenSize]uint8

/**
 * @Description: to start the module
 * @return error
 */
func GraphInit() error {
	file, err := os.OpenFile(devPath, os.O_WRONLY, 0664)
	defer file.Close()
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return err
	} else {
		dev = file
		fmt.Println("GraphGo start successfully!")
	}
	return err
}

/**
 * @Description: to end up the module
 */
func GraphBye() {
	dev.Close()
	dev = nil
}

/**
 * @Description: set BgColor and reset the BgColor Buff
 * @param color
 */
func SetBgColor(color GColor) {
	bgColor = color
	for i := 0; i < screenSize; i += 3 {
		backgroundBuff[i] = bgColor.B
		backgroundBuff[i+1] = bgColor.G
		backgroundBuff[i+2] = bgColor.R
	}
}

func SetGraphColor(color GColor) {
	graphColor = color
}

func SetFontColor(color GColor) {
	fontColor = color
}

/**
 * @Description: to fill screen with bgColor
 */
func ResetScreen() {

}
