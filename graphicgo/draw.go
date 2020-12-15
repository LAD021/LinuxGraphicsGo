package graphicgo

import "fmt"

func getPoint(x int64, y int64) (newPos int64, err error) {
	return dev.Seek((x*screenWidth+y)*pixWidth, 0)
}

func DrawDot(x int64, y int64) {
	_, err := getPoint(x, y)
	if err != nil {
		fmt.Println("error:", err)
	}
	dev.Write(graphBuff[:])
}
