package graphicgo

import "fmt"

//func getPoint(x int64, y int64) (newPos int64, err error) {
//	return dev.Seek((x+screenHeight*y)*pixWidth, 0)
//}
//
//func dot(x int64, y int64, color [4]byte) {
//	_, err := getPoint(x, y)
//	if err != nil {
//		fmt.Println("error:", err)
//	}
//	dev.Write(color[:])
//}

func getPoint(x int64, y int64) (newPos int64, err error) {
	return (x + screenHeight*y) * pixWidth, nil
}

func dot(x int64, y int64, color [4]byte) {
	if x >= 0 && y >= 0 {
		pos, err := getPoint(x, y)
		if err != nil {
			fmt.Println("error:", err)
		}
		drawBuff[pos] = color[0]
		drawBuff[pos+1] = color[1]
		drawBuff[pos+2] = color[2]
		drawBuff[pos+3] = color[3]
	}
}
