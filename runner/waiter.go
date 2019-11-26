package runner

import (
	"time"

	"github.com/criyle/go-judge/pkg/runner"
)

const minCPUPercent = 40 // 40%
const checkIntervalMS = 50

type waiter struct {
	timeLimit time.Duration
}

func (w *waiter) Wait(done chan struct{}, cg runner.CPUAcctor) bool {
	var lastCPUUsage uint64
	var totalTime time.Duration
	lastCheckTime := time.Now()
	// wait task done (check each interval)
	ticker := time.NewTicker(checkIntervalMS * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case now := <-ticker.C: // interval
			cpuUsage, err := cg.CpuacctUsage()
			if err != nil {
				return true
			}

			cpuUsageDelta := time.Duration(cpuUsage - lastCPUUsage)
			timeDelta := now.Sub(lastCheckTime)

			totalTime += durationMax(cpuUsageDelta, timeDelta*minCPUPercent/100)
			if totalTime > w.timeLimit {
				return true
			}

			lastCheckTime = now
			lastCPUUsage = cpuUsage

		case <-done: // returned
			return false
		}
	}
}

func durationMax(a, b time.Duration) time.Duration {
	if a > b {
		return a
	}
	return b
}