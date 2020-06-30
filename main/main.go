package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	//"sync/atomic"

	"github.com/speps/go-hashids"
	"math/big"
	"math/rand"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"sync"
	"testing"
	"time"
)

type MyTask struct {
	task chan int64
}

var (
	mytask = &MyTask{}
)

func init() {
	mytask.task = make(chan int64, 1000)
}

func main() {
	//http.HandleFunc("/", handle)
	//http.HandleFunc("/hello", handleHello)
	//http.HandleFunc("/active", handleActiveAnother)
	//http.HandleFunc("/kill", handleKill)
	//go consumer()
	//fmt.Println("start http serve")
	//http.ListenAndServe("127.0.0.1:8008", nil)

	//str := "hello world!"
	//fmt.Println(strings.Contains(str, ""))
	//fmt.Println(strings.Contains(str, " "))
	//fmt.Println(strings.IndexAny(str, " "))
	//
	//strings.Replace(str, " ", "", 5)
	//fmt.Println(str)

	// atomicTest() // atomic包一些方法使用
}

func atomicTest() {
	//atu := NewAtomicUser()
	//fmt.Println("age === ", atu.age.Load(), "; type :", reflect.TypeOf(atu.age.Load()))
	//fmt.Println("name === ", atu.name.Load(), "; type :", reflect.TypeOf(atu.name.Load()))
	//
	//val := atu.age.Load().(uint64)
	//atomic.AddUint64(&val, 56)
	//fmt.Println("age === ", atu.age.Load(), "; type :", reflect.TypeOf(atu.age.Load()))
	//
	//atomic.CompareAndSwapUint64(&val, atu.age.Load().(uint64), val)
	//fmt.Println("age === ", atu.age.Load(), "; type :", reflect.TypeOf(atu.age.Load()))
}

// 下面两个函数是有区别的
func errTest(arr [3]int) error {
	return nil
}
func errTest1(arr []int) error {
	return nil
}
func add(a []int, b int, c int) (d []int) {
	arr1 := new([]int)

	*arr1 = append(*arr1, a[:b]...)

	*arr1 = append(*arr1, c)

	*arr1 = append(*arr1, a[b:]...)

	fmt.Println(a)
	fmt.Println(*arr1)
	return a
}

func add1(a []int, b int, c int) (d []int) {
	var abcd []int

	abcd = append(abcd, a[:b]...)

	abcd = append(abcd, c)

	abcd = append(abcd, a[b:]...)

	fmt.Println(abcd)
	fmt.Println(a)
	return a
}

func add2(a []int, b int, c int) (d []int) {

	fmt.Println("before append:a == >", a)

	abcd := append(a[:b], c)

	fmt.Println("after append:a == >", a)

	fmt.Println("after append:abcd == >", abcd)
	return a
}

// 指针类型
func pti() {
	var j int = 5
	fmt.Println(reflect.TypeOf(j))
	fmt.Println(reflect.TypeOf(&j))
	fmt.Println(j)
	fmt.Println(&j)
	i := new(int)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(*i))
	fmt.Println(reflect.TypeOf(&i))
	fmt.Println(i)
	fmt.Println(*i)
	fmt.Println(&i)
}

//var ARGS string
//
//func shellTest() {
//
//	var uptime *bool = new(bool)
//	flag.BoolVar(uptime,"u", false, "print system uptime")
//	flag.Parse()
//
//	ARGS = strings.Join(flag.Args(), " ")
//	if len(os.Args) < 2 {
//		flag.Usage()
//		os.Exit(1)
//	}
//
//	if *uptime {
//		fmt.Println("12 days")
//	}
//}

func hello(num ...int) {
	num[0] = 18
}

func Test13(t *testing.T) {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
}

// 1.不同类型的切片不可复制
// 2.一个切片的原始值不可修改;但是遍历的时候又需要做一些特殊的操作,如赋值等;这时可以使用copy方法;
func testCopy() {
	arr1 := make([]int, 3)
	arr2 := []int{1, 2, 3, 4, 5}
	copy(arr1, arr2)
	arr1[0] = 100
	fmt.Println("copy后arr1 == ", arr1, "arr2 == ", arr2)

	arr3 := arr2
	arr3[0] = 101
	fmt.Println("arr3 == ", arr3, "arr2 == ", arr2)
}

// 比较两个结构体
func compareStruct() {
	type (
		T1 struct {
			Name string
			Age  int
		}
		T2 struct {
			Name   string
			Age    int
			Gender int
		}
	)

	t1 := T1{"张三", 18}
	t2 := T1{"张三", 18}
	//t2 := T2{"张三", 18, 1}
	equal := reflect.DeepEqual(t1, t2)
	fmt.Println("比较内容 ==> T1 == T2 :", equal)

	t3 := &T1{"张三", 18} // 地址:0xc0000584a0
	t4 := &T1{"张三", 18} // 地址:0xc0000584c0
	//t4 := t3 // 这俩相等,
	t4.Name = "王二"
	fmt.Println("t3 地址:", t3, "  t4 地址:", t4, "    t3 === t4:", t3 == t4)
}

func defaultRoute(w http.ResponseWriter, response *http.Request) {
	fmt.Println("hello world!")
}

type A struct {
	a1 int
	a2 []int
	a3 int
}

type duoliaoResult struct {
	ResTodo string `json:"resTodo"`
	ResMsg  string `json:"resMsg"`
	ResData string `json:"resData"`
}

// 将int类型的数字加密一下
// 文档传送门 ==> https://godoc.org/github.com/speps/go-hashids
// GitHub传送门 ==> https://github.com/speps/go-hashids
func hashIds() {
	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	//e, _ := h.Encode([]int{45, 434, 1313, 99})
	e, _ := h.Encode([]int{45, 434, 1313, 99})
	fmt.Println(e)
	d, _ := h.DecodeWithError(e)
	fmt.Println(d)
}

const (
	BWillChange    = iota // 0换玩家位置 ==> 换01
	ChangeOneTwo          // 换12
	ChangeTwoThree        // 换23
)

func drawByRank2(before [4]int) [4]int {
	if rand.Intn(2) != BWillChange {
		return before
	}

	res := [4]int{}
	copy(res[:], before[:])

	switch rand.Intn(3) {
	case ChangeOneTwo:
		res[1], res[2] = res[2], res[1]
	case ChangeTwoThree:
		res[2], res[3] = res[3], res[2]
	default:
		res[0], res[1] = res[1], res[0]
	}
	return res
}

// 随机抽签 ==> 4个玩家是否做队友
func drawByRank1(before [4]int) (res [4]int) {
	teamA := [2]int{before[0], before[2]}
	teamB := [2]int{before[1], before[3]}

	tmp := []int{0, 1, 2, 3}
	iCardPosition, iCardIndex := 0, 0
	rNode := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 4; i++ {
		iCardPosition = rNode.Intn(len(tmp))
		res[iCardIndex] = tmp[iCardPosition]
		iCardIndex++
		tmp = append(tmp[:iCardPosition], tmp[iCardPosition+1:]...)
	}

	newA := [2]int{res[0], res[2]}
	//newB := [2]int{res[1], res[3]}

	fmt.Printf("teamA == %v;newA == %v \n", teamA, newA)
	fmt.Printf("before == %v;res == %v \n", before, res)
	sort.Ints(newA[:])
	sort.Ints(teamB[:])
	sort.Ints(teamA[:])
	if reflect.DeepEqual(newA, teamA) || reflect.DeepEqual(newA, teamB) {
		fmt.Println("相等")
		return before
	} else {
		fmt.Println("不相等")
		return res
	}
}

//var ALL_CARDS = []byte{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
//var ALL_CARDS = []byte{0,1,2,3,4,5,6}
var ALL_CARDS = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, //红桃23456789 10 JQKA 共13张牌
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, //黑桃
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, //梅花
	0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, //方块
	0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4A, 0x4B, 0x4C, 0x4D, //方块
	0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5A, 0x5B, 0x5C, 0x5D, //方块
	0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A, 0x6B, 0x6C, 0x6D, //方块
	0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7A, 0x7B, 0x7C, 0x7D, //方块
	0x8E, 0x8F, //小王
	0x9E, 0x9F}

func shuffle1() []byte {
	//结果数据
	var result_cards []byte
	result_cards = make([]byte, len(ALL_CARDS))
	//存储临时牌的变量
	tmp_cards := make([]byte, len(ALL_CARDS))
	copy(tmp_cards, ALL_CARDS)
	//记录牌的序号
	iCardIndex := 0
	//记录牌的位置
	iCardPosition := 0
	//生成随机数种子
	r_node := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < len(ALL_CARDS); i++ {
		iCardPosition = r_node.Intn(len(tmp_cards))
		result_cards[iCardIndex] = tmp_cards[iCardPosition]
		iCardIndex++
		tmp_cards = append(tmp_cards[:iCardPosition], tmp_cards[iCardPosition+1:]...)
	}
	return result_cards
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func shuffle2() []byte {
	begin := time.Now()
	tmp_cards := make([]byte, len(ALL_CARDS))
	copy(tmp_cards, ALL_CARDS)
	for i := len(ALL_CARDS); i > 0; i-- {
		lastIdx := i - 1
		idx := rand.Intn(i)
		tmp_cards[lastIdx], tmp_cards[idx] = tmp_cards[idx], tmp_cards[lastIdx]
	}
	fmt.Println(time.Now().Sub(begin))
	return tmp_cards
}
func setData(x int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("恢复异常", err)
		}
	}()

	arr := make([]int, 10)
	arr[x] = 88
}

func main111() {

	setData(20)
	fmt.Println("走到这里 ")

	//ch := make(chan int, 10)
	//for i := 0; i < 1; i++ {
	//	select {
	//	case ch <- i: // 数字无意义，只为了通知
	//		fmt.Println("Notify userId:", i, " user:", i, " for task complete")
	//	default:
	//		fmt.Println("User:", i, " user:", i, " task channel has reached the limit")
	//	}
	//}
	//fmt.Println(len(ch))
	//time.Sleep(60 * time.Second)

	//var cnt1 = map[int]int{}
	////for i := 0; i < 10000000; i++ {
	////	c1 := shuffle1()
	////	cnt1[int(c1[0])]++
	////}
	//var cnt2 = map[int]int{}
	//for i := 0; i < 2; i++ {
	////for i := 0; i < 1000000; i++ {
	//	c2 := shuffle2()
	//	cnt2[int(c2[0])]++
	//}
	//fmt.Println(cnt1,"\n",cnt2)

	//for i := 0; i < 100; i++ {
	//	fmt.Println(drawByRank2([4]int{0,1,2,3}))
	//}

	//fmt.Println(drawByRank1([4]int{0, 1, 2, 3}))

	//hashIds()

	//mmp := make(map[uint8]int, 4)
	//mmp[1] = 1
	//mmp[2] = 2
	//
	//if v, ok := mmp[0]; ok {
	//	fmt.Println("存在", v)
	//} else {
	//	fmt.Println("不存在", v)
	//}
	//
	//if mmp[0] == 0 {
	//	fmt.Println("不对")
	//} else {
	//	fmt.Println(" ok")
	//}

	//
	//duoliaoRes := duoliaoResult{}
	//res := `{"resTodo":"01","resMsg":"\u6210\u529f","resData":[]}`
	//err := json.Unmarshal([]byte(res),&duoliaoRes)
	//fmt.Println(err,duoliaoRes)
	//if duoliaoRes.ResTodo == "01"{
	//	fmt.Println("")
	//}

	//a := &A{}
	//a.a2 = make([]int, 5)
	//a.a2[5] = 10
	//
	//fmt.Println(a)

	//a := uint8(255)
	//fmt.Println(a + 1)

	//ch := make(chan int, 10)
	//for i := 0; i <= 11; i++ {
	//	select {
	//	case ch <- i:
	//		fmt.Println("放入:", i)
	//	default:
	//		fmt.Println("放不进去:", i)
	//	}
	//}
	//
	//i := <-ch
	//fmt.Println(i,ch)
	//time.Sleep(time.Second * 60 * 60)

	//http.HandleFunc("/", defaultRoute)
	//http.ListenAndServe("localhost:20202", nil)

	//compareStruct()
	//testCopy()

	// 数组操作
	//a := []int{1,2,3}
	//fmt.Println("a ==== ",a)
	//a = a[3:]
	//fmt.Println("a ==== ",a)

	// 切片排序
	//arr := []int{1,4,2,5,123,45,123}
	//
	//sort.Ints(arr)
	//fmt.Println(arr)

	// map的初始化
	//var m map[uint8]int
	//m := make(map[uint8]int)
	//m[1] = 2
	//fmt.Println("map==>",m)

	/*// 位运算
	var a uint8 = 0x82
	var b uint8 = 0x02
	fmt.Printf("%08b [A]\n", a)
	fmt.Printf("%08b [B]\n", b)
	fmt.Printf("%08b (NOT B)\n", ^b)
	fmt.Printf("%08b ^ %08b = %08b [B XOR 0xff]\n", b, 0xff, b^0xff)
	fmt.Printf("%08b ^ %08b = %08b [A XOR B]\n", a, b, a^b)
	fmt.Printf("%08b & %08b = %08b [A AND B]\n", a, b, a&b)
	fmt.Printf("%08b &^%08b = %08b [A 'AND NOT' B]\n", a, b, a&^b)
	fmt.Printf("%08b&(^%08b)= %08b [A AND (NOT B)]\n", a, b, a&(^b))*/

	// the time unix
	//timeUnix := time.Now().Unix()
	//fmt.Println("时间戳转化为日期格式 ==> ", time.Unix(timeUnix, 0).Format("2006-01-02"))

	// ♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥
	// ♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥ 分 ♥♥ 隔 ♥♥ 符 ♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥
	// ♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥
	// ASCII 字符串长度使用 len() 函数。
	// Unicode 字符串长度使用 utf8.RuneCountInString() 函数。
	//fmt.Println(utf8.RuneCountInString("我是一个石头,abc"))

	//fmt.Println(time.Now().AddDate(0, 0, 1).Unix())
	//l := list.New()
	//l.PushFront(1)
	//l.PushFront(2)
	//l.PushFront(3)
	//l.PushFront(4)
	//l.PushFront(4)
	////l.PushFront(5)
	//
	//fmt.Println(l.Len())
	//for e := l.Front(); e != nil; e = e.Next() {
	//	if i, ok := e.Value.(int); ok {
	//		fmt.Println(e.Value,";i = ",i)
	//	}
	//}
	//fmt.Println()

	//u, err := user.Current()
	//fmt.Println(u,err)
	//if err != nil {
	//	fmt.Printf("Current: %v (got %#v) \n", err, u)
	//}
	//if u.HomeDir == "" {
	//	fmt.Printf("didn't get a HomeDir")
	//}
	//fmt.Println(u.HomeDir)
	//if u.Username == "" {
	//	fmt.Printf("didn't get a username")
	//}
	//fmt.Println(u.Username)
	//pti()

	//abc := []int{1, 2, 3, 4, 5, 6}
	//var x int = 3
	//var i int = 100
	//add2(abc, x, i)

	//s := ""
	//fmt.Println(s)
	//fmt.Println(&s)
	//
	//err := errTest()
	//fmt.Println(err)
	//fmt.Println(&err)

	//for {
	//	fmt.Println("请输入两个参数,用空格号隔开: ")
	//	fmt.Scanln(&firstName, &lastName)
	//	fmt.Printf("input == %v,%v \n", firstName, lastName)
	//}
	//fmt.Sscanf(firstName, lastName,  test1)
	//fmt.Println("From the string we read: ", firstName, lastName, test1)
	//fmt.Println("From the string we read: ", f, i, s)

	//fmt.Println(ToBigFloat("12.df2354541bhj"))
	//fmt.Println(BigFloatToString(big.NewFloat(10.241)))
	//fmt.Println(BigFloatToString(big.NewFloat(10.241254)))
	//fmt.Println(BigFloatToString(big.NewFloat(10.2419254)))
	//fmt.Println(BigFloatToString(big.NewFloat(10.24)))
	//fmt.Printf("a2 = %s\n", a2.String())

	//for i := 0;i < 10000;i ++ {
	//	go appendValueNoMutex(i)	// 不加锁会丢失
	//	//go appendValueWithMutex(i)	// 加锁不会丢失
	//}
	////time.Sleep(time.Second)
	//sort.Ints(s)
	//for i,v := range s{
	//	fmt.Println(i,":",v)
	//}
}
func ToBigFloat(str string) *big.Float {
	f, _, _ := big.ParseFloat(str, 10, 12, big.ToNearestEven)
	return f
}

// big.float转string(保留三位小数)
func BigFloatToString(b *big.Float) string {
	return b.Text('f', 3)
}

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

var s []int
var lock sync.Mutex

// 不加锁
func appendValueNoMutex(i int) {
	s = append(s, i)
}

func appendValueWithMutex(i int) {
	lock.Lock()
	s = append(s, i)
	lock.Unlock()
}

// 打招呼
func greet() {
	now := time.Now()
	zeroOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if now.Before(zeroOfToday.Add(15 * time.Hour)) {

	}
}

func handleActiveAnother(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Cmd{
		Path: "/Users/jiahua-zhoujian/Tools/golang/src/MyGoTest/main/httptest",
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Start child process with pid", cmd.Process.Pid)
	// Wait releases any resources associated with the Cmd
	go func() {
		if err := cmd.Wait(); err != nil {
			fmt.Printf("Child process %d exit with err: %v\n", cmd.Process.Pid, err)
		}
	}()

	//args := []string{"a", "b", "c"}
	//cmd := exec.Command("/Users/jiahua-zhoujian/Tools/golang/src/MyGoTest/main/httptest", args...) //不加第一个第二个参数会报错
	////cmd := exec.Command("/bin/bash", "-c", "/Users/jiahua-zhoujian/Tools/golang/src/MyGoTest/main/mytest") //不加第一个第二个参数会报错
	//fmt.Println("fork子进程ing ")
	//
	var HtmlBuffer bytes.Buffer
	//
	//con, err := cmd.Output()
	//if err != nil {
	//	fmt.Println("命令执行错误: ===== ", err)
	//	HtmlBuffer.WriteString("fork子进程失败:" + err.Error())
	//	io.WriteString(w, HtmlBuffer.String())
	//	return
	//} else {
	//	fmt.Println(" 输出内容 ==== ", string(con))
	//}
	//ProId = cmd.ProcessState.Pid()

	fmt.Println(" ExitCode === ", cmd.ProcessState.ExitCode())
	fmt.Println(" Exited === ", cmd.ProcessState.Exited())
	fmt.Println(" ards === ", cmd.Args)

	pro, err := os.FindProcess(cmd.ProcessState.Pid())
	if err != nil {
		fmt.Println("pid == ", cmd.ProcessState.Pid(), " err === ", err)
	}
	ps, err := pro.Wait()
	if err != nil {
		fmt.Println("pid == ", cmd.ProcessState.Pid(), " err === ", err)
	}
	fmt.Println(" ps.ExitCode() ===  ", ps.ExitCode())

	HtmlBuffer.WriteString("hello world")
	HtmlBuffer.WriteString(req.Host)
	io.WriteString(w, HtmlBuffer.String())
}
func handleKill(w http.ResponseWriter, req *http.Request) {
	fmt.Println(" kill .....")
	var HtmlBuffer bytes.Buffer

	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("kill %d", 0))
	//cmd := command.Cmd{}
	//err := cmd.StdRun(fmt.Sprintf("ps aux|grep %s|grep -v \"grep\"|awk '{print $1}' | xargs kill -9", "httptest"))
	err := cmd.Run()
	if err != nil {
		HtmlBuffer.WriteString("err === " + err.Error())
	}

	HtmlBuffer.WriteString("kill")
	HtmlBuffer.WriteString(req.Host)
	io.WriteString(w, HtmlBuffer.String())
}
func handleHello(w http.ResponseWriter, req *http.Request) {
	randseed := int64(rand.Int())
	mytask.task <- randseed
	var HtmlBuffer bytes.Buffer

	HtmlBuffer.WriteString("hello world" + string(randseed))
	HtmlBuffer.WriteString(req.Host)
	io.WriteString(w, HtmlBuffer.String())
}
func handle(w http.ResponseWriter, req *http.Request) {

	var HtmlBuffer bytes.Buffer

	HtmlBuffer.WriteString("HTTP Method : ")
	HtmlBuffer.WriteString(req.Method)

	HtmlBuffer.WriteString("\nHTTP Proto : ")
	HtmlBuffer.WriteString(req.Proto)

	HtmlBuffer.WriteString("\nhttp Host : ")
	HtmlBuffer.WriteString(req.Host)

	HtmlBuffer.WriteString("\nRequest RemoteAddr : ")
	HtmlBuffer.WriteString(req.RemoteAddr)

	HtmlBuffer.WriteString("\nRequestURI : ")
	HtmlBuffer.WriteString(req.RequestURI)

	io.WriteString(w, HtmlBuffer.String())
}

func consumer() {
	for {
		select {
		case v := <-mytask.task:
			fmt.Printf("hello, %v has been dealed \n", v)
		}
	}
}
