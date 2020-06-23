package main

import "sync"

var mySyncMap sync.Map

func testMySyncMap() {
	mySyncMap.Store("key1", "hello world")
}
