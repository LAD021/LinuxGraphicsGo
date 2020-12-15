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
	time.Sleep(1)
	graphicgo.SetBgColor(graphicgo.PaleTurquoise1)
	graphicgo.GraphBye()
}
