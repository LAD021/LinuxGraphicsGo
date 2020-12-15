package graphicgo

import (
	"fmt"
	"sync"
	"time"
)

const (
	StopCmd = iota
)

type refreshJob struct {
	chCmd chan int
	FPS   float64
}

var (
	refresherOnce sync.Once
	instance      *refreshJob = nil
)

func NewRefreshJob() *refreshJob {
	refresherOnce.Do(func() {
		instance = &refreshJob{
			chCmd: make(chan int),
			FPS:   24,
		}
	})
	return instance
}

func (job *refreshJob) SetFPS(fps float64) {
	job.FPS = fps
}

func (job *refreshJob) Start() {
	if job.FPS == 0 {
		job.FPS = 24
	}
	ticker := time.NewTicker(time.Second / time.Duration(1/job.FPS))
	go func() {
		fmt.Println("Refresher Start!")
		for {
			select {
			case <-ticker.C:
				resetScreen()
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
