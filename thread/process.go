package thread

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"os"
	"strings"
)

// 判断某个线程是否已经存在
func CheckProcessByName(serName string) {
	//if pid, _ := tools.GetPid(serName); pid != "" {
	//	panic(fmt.Sprintf("%v is already exist:pid === %v", serName, pid))
	//}

	curPid := os.Getpid()
	processes, err := process.Processes()
	if err != nil {
		fmt.Printf("find process err : %v", err)
	}
	for _, v := range processes {
		name, _ := v.Name()
		if int32(curPid) != v.Pid && strings.Contains(name, serName) {
			panic(fmt.Sprintf("%v is already exist:pid === %v ;  process Name : %v", serName, v.Pid, name))
		}
	}
	fmt.Printf("%v is not exist,continue to start", serName)
}
