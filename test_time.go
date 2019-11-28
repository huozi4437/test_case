package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("M:01 Y:2006 D:02 15 04:05"))
	fmt.Println(time.Now())
}
