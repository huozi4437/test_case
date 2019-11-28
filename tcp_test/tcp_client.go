package main

import (
	"log"
	"net"
)

func main()  {
	Dial(":30002")
}

func Dial(addr string) {
	log.Println("Dial to:", addr)
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		log.Panicf("Dial %s panic:%s", addr, err)
	}
}