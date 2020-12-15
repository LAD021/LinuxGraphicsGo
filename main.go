package main

import (
	"os"
	"sync"
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
	job := graphicgo.NewRefreshJob()
	job.SetFPS(100)
	job.Start()
	//for i := int64(0); i < int64(100); i++ {
	//	graphicgo.DrawDot(i, i*2, graphicgo.RED)
	//}
	var wg sync.WaitGroup
	wg.Add(2)
	wg.Wait()
}
