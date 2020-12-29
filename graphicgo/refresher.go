package graphicgo

import (
	"fmt"
	"sync"
	"time"
)

const (
	StopCmd = iota
	StartCmd
)

type refreshJob struct {
	chCmd      chan int
	FPS        int64
	RefreshSig chan int
}

var (
	refresherOnce sync.Once
	instance      *refreshJob = nil
)

func GetRefreshJob() *refreshJob {
	refresherOnce.Do(func() {
		instance = &refreshJob{
			chCmd: make(chan int, 100),
			FPS:   24,
		}
	})
	return instance
}

func (job *refreshJob) SetFPS(fps int64) {
	job.FPS = fps
}

func (job *refreshJob) Start() {
	if job.FPS == 0 {
		job.FPS = 24
	}
	timeSpace := time.Duration(int64(time.Second) / job.FPS)
	ticker := time.NewTicker(timeSpace)
	go func() {
		fmt.Println("Refresher Start!")
		for {
			select {
			case <-ticker.C:
				//fmt.Println("Printing")
				go func() {
					a := <-job.RefreshSig
					fmt.Println(a)
				}()
				job.RefreshSig <- StartCmd
				fmt.Println("Done")
				resetScreen()
				refreshBg() // test

			case cmd := <-job.chCmd:
				if cmd == StopCmd {
					ticker.Stop()
					fmt.Println("Refresher Stopped!")
					break
				}
			}
		}
	}()
}

func (job *refreshJob) Stop() {
	fmt.Println("Refresher Stopped")
	job.chCmd <- StopCmd
}
