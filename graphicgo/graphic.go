package graphicgo

import (
	"fmt"
	"os"
)

const screenSize = screenWidth * screenHeight * pixWidth

var backgroundBuff [screenSize]byte
var graphBuff [4]byte

var dev *os.File
var bgColor = BLACK
var graphColor = WHITE
var fontColor = WHITE

/**
 * @Description: to start the module
 * @return error
 */
func GraphInit() error {
	file, err := os.OpenFile(devPath, os.O_RDWR, 0664)
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return err
	} else {
		dev = file
		fmt.Println("GraphGo start successfully!")
	}
	SetBgColor(RED)
	SetGraphColor(GREEN)
	// start building the screen buff
	ResetScreen()
	return err
}

/**
 * @Description: to end up the module
 */
func GraphBye() {
	fmt.Println("GraphGo goodbye!")
	dev.Close()
	dev = nil
}

func GraphWrong() {
	fmt.Println("Wrong")
	dev.Close()
	dev = nil
}

/**
 * @Description: set BgColor and reset the BgColor Buff
 * @param color
 */
func SetBgColor(color [4]byte) {
	for i := 0; i < int(screenSize); i += int(pixWidth) {
		backgroundBuff[i] = color[0]
		backgroundBuff[i+1] = color[1]
		backgroundBuff[i+2] = color[2]
		backgroundBuff[i+3] = color[3]
	}
}

func SetGraphColor(color [4]byte) {
	graphColor = color
}

func SetFontColor(color [4]byte) {
	fontColor = color
}

/**
 * @Description: to fill screen with bgColor
 */
func ResetScreen() {
	dev.Seek(0, 0)
	_, err := dev.Write(backgroundBuff[:])
	if err != nil {
		fmt.Println(err)
		GraphWrong()
	}
}
