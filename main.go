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
	graphicgo.SetBgColor(graphicgo.GREEN)
	graphicgo.ResetScreen()
	for i := int64(0); i < int64(100); i++ {
		graphicgo.DrawDot(i, i*2)
	}
}
