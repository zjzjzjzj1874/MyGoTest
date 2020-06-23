package main

import (
	"net/http"
)

func httpListener() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/active", handleActiveAnother)
	http.HandleFunc("/kill", handleKill)
	http.ListenAndServe("127.0.0.1:8008", nil)

	//in := bytes.NewBuffer(nil)
	//cmd := exec.Command("sh")
	//cmd.Stdin = in
	//go func() {
	//	in.WriteString("echo hello world > test.txt\n")
	//	in.WriteString("exit\n")
	//}()
	//if err := cmd.Run(); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//cmd1 := exec.Command("ps -ef | grep redis")
	//args := []string{"/Users/jiahua-zhoujian/Tools/golang/src/MyGoTest/mytest/mytest", "&"}
	//cmd := exec.Command("/bin/bash", "-c", "nohup /Users/jiahua-zhoujian/Tools/golang/src/MyGoTest/mytest/mytest&") //不加第一个第二个参数会报错

	//后台进程启动
	//fmt.Println("hello world")
	//for i := 0; i < 5; i++ {
	//	fmt.Println(fmt.Sprintf("No.%d,hello %d", i, i))
	//	if i == 2 {
	//		fmt.Println("我提前结束了...")
	//		return
	//	}
	//	time.Sleep(time.Second * 10)
	//}
	//fmt.Println("finally,i am done.")
}
