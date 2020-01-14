package main

import "fmt"

// We can do better by exploiting the fact that arguments to deferred
// functions are evaluated when the deferexecutes. The tracing routine
// can set up the argument to the untracing routine. This example:
func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	// attention:the function of trace() will execute first;
	// then the defer will push un() to the function stack.
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	b()
}
