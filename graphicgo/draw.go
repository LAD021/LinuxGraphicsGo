package graphicgo

import "fmt"

func getPoint(x int64, y int64) (newPos int64, err error) {
	return dev.Seek((x+screenHeight*y)*pixWidth, 0)
}

func DrawDot(x int64, y int64, color [4]byte) {
	_, err := getPoint(x, y)
	if err != nil {
		fmt.Println("error:", err)
	}
	dev.Write(color[:])
}
