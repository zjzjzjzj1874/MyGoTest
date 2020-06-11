package thread

import (
	"fmt"
	"testing"
	"time"
)

// 将对应的path.workingDir和标准输出地址修改一下即可
func TestRunner_DaemonProcessing(t *testing.T) {
	var runner = &Runner{
		Name:       "main",
		Path:       "/Users/jiahua-zhoujian/workspace/golang/src/git.querycap.com/ss/srv-ota/srv-ota",
		Args:       nil,
		WorkingDir: "/Users/jiahua-zhoujian/workspace/golang/src/git.querycap.com/ss/srv-ota",
		Envs: []string{
			"RUNNER_DEMO1=10",
			"RUNNER_DEMO2=20",
		},
		Output: "/Users/jiahua-zhoujian/workspace/golang/src/git.querycap.com/ss/srv-ota/srv-ota/demo_output",
	}

	runner.Start()

	for i := 0; i < 2; i++ {
		if err := runner.Restart(); err != nil {
			fmt.Printf("No. %d , ----------- err : %v", i, err)
			return
		}
	}

	time.Sleep(100 * time.Second)
}
