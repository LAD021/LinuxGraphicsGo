package graphicgo

import (
	"fmt"
	"os"
)

const screenSize = screenWidth * screenHeight * pixWidth

var backgroundBuff [screenSize]byte
var drawBuff [screenSize]byte

var dev *os.File
var bgColor = BLACK

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
	refreshBgColor()
	// start building the screen buff
	resetScreen()
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

func SetBgColor(color [4]byte) {
	bgColor = color
}

func refreshBgColor() {
	for i := 0; i < int(screenSize); i += int(pixWidth) {
		backgroundBuff[i] = bgColor[0]
		backgroundBuff[i+1] = bgColor[1]
		backgroundBuff[i+2] = bgColor[2]
		backgroundBuff[i+3] = bgColor[3]
	}
}

func refreshBg() {
	drawBuff = backgroundBuff
}

/**
 * @Description: to fill screen with bgColor
 */
func resetScreen() {
	dev.Seek(0, 0)
	_, err := dev.Write(drawBuff[:])
	if err != nil {
		fmt.Println(err)
		GraphWrong()
	}
}
