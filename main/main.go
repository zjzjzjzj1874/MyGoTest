package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()
	timeString := n.AddDate(0,0,-1).Format("2006-01-02")
	yesterday := timeString + " 22:00:00"
	fmt.Println("yesterday ==== ",yesterday)
}





