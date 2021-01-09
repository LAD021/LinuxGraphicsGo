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
	graphicgo.SetBgColor(drawork.BkColor)
	var wg sync.WaitGroup
	job := graphicgo.GetRefreshJob()
	job.SetFPS(drawork.FPS)

	job.SetInit(func() {
		time.Sleep(drawork.Stardelay)
		drawork.InitCube()
		drawork.DrawCube()
	})

	job.SetWork(func() {
		//graphicgo.DrawLine(300, 300, 400, 500, graphicgo.RED, graphicgo.Bold)
		//graphicgo.DrawCircle(200, 200, 150, graphicgo.RED, graphicgo.Bold, false)
		drawork.Rotation()
		drawork.Changing()
		drawork.DrawCube()
		drawork.IfEnd()
		if drawork.End == 0 {
			time.Sleep(drawork.EndDelay)
			job.Stop()
		}
	})
	job.Start()
	defer job.Stop()
	wg.Add(1)
	fmt.Println("Prepared")
	wg.Wait()
}
