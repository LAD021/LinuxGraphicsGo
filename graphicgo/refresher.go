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
	work       func()
	init       func()
	wg         *sync.WaitGroup
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
	job.wg.Add(1)
	go func() {
		fmt.Println("Refresher Start!")
		job.init()
		ticker := time.NewTicker(timeSpace)
		for {
			select {
			case <-ticker.C:
				resetScreen()
				refreshBg()
				job.work()

			case cmd := <-job.chCmd:
				if cmd == StopCmd {
					ticker.Stop()
					fmt.Println("Refresher Stopped!")
					job.wg.Done()
					break
				}
			}
		}
	}()
}

func (job *refreshJob) SetWork(f func()) {
	job.work = f
}

func (job *refreshJob) SetInit(f func()) {
	job.init = f
}

func (job *refreshJob) Stop() {
	fmt.Println("Refresher Stopped")
	job.chCmd <- StopCmd
}

func (job *refreshJob) SetWaitGroup(wg *sync.WaitGroup) {
	job.wg = wg
}
