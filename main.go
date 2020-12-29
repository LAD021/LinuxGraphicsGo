package main

import (
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
	graphicgo.SetBgColor(graphicgo.GREEN)
	job := graphicgo.GetRefreshJob()
	job.SetFPS(100)
	job.Start()
	defer job.Stop()
	go func() {
		for {

			select {
			case cmd := <-job.RefreshSig:
				fmt.Println("we are doing")
				if cmd == graphicgo.StartCmd {
					for i := int64(0); i < int64(100); i++ {
						go graphicgo.DrawDot(i, i*2, graphicgo.RED, graphicgo.Middle)
						fmt.Println(i)
					}
				}
			}
		}
	}()
	fmt.Println("Prepared")
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
