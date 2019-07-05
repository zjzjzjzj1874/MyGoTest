package util

import (
	"fmt"
	"time"
)

func Timer(second time.Duration)  {
	timer(second)
}

// 简单的定时器实现
func timer(duration time.Duration)  {
	tick := time.NewTicker(duration * time.Second)
	for {
		select {
		case <-tick.C:
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}
