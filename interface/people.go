package people

// 理解go语言接口
type People interface {
	Say()
	Eat()
	Sleep()
}

// 这里只是实现了 People 接口方法，并没有将People和Man显式绑定。这是一种非侵入式的设计。
func DoPeople(people People) {
	people.Say()
	people.Eat()
	people.Sleep()
}
