package main

import (
	"os"
)
import "LinuxGraphicsGo/graphicgo"

func main() {
	if graphicgo.GraphInit() != nil {
		os.Exit(-1)
	}

	graphicgo.GraphBye()
}
