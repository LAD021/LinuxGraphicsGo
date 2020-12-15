package main

import (
	"os"
	"time"
)
import "LinuxGraphicsGo/graphicgo"

func main() {
	if graphicgo.GraphInit() != nil {
		os.Exit(-1)
	}
	defer graphicgo.GraphBye()
	time.Sleep(100)
	graphicgo.SetBgColor(graphicgo.BLUE)
	graphicgo.ResetScreen()
	graphicgo.DrawDot(100, 100)
}
