package main

import (
	"LinuxGraphicsGo/drawork"
	"fmt"
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
	graphicgo.SetBgColor(drawork.BkColor)
	job := graphicgo.GetRefreshJob()
	job.SetFPS(drawork.FPS)
	job.SetWork(func() {
		graphicgo.DrawLine(300, 300, 400, 500, graphicgo.RED, graphicgo.Bold)
		graphicgo.DrawCircle(200, 200, 150, graphicgo.RED, graphicgo.Bold, false)
	})
	job.Start()
	defer job.Stop()

	fmt.Println("Prepared")
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
