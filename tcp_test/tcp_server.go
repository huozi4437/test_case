package main

import (
	"log"
	"net"
)

func main()  {
	Listen(":30002")
}

func Listen(addr string) {
	Listener, err := net.Listen("tcp", addr)
	defer Listener.Close()
	if err != nil {
		log.Panicf("Listen %s panic:%v", addr, err)
	}
	log.Println(Listener.Addr().Network(), Listener.Addr().String())

	for {
		conn, err := Listener.Accept()
		defer conn.Close()
		if err != nil {
			log.Panicf("Accept %s panic:%s", addr, err)
		}
		Handle(conn)
	}
}

func Handle(conn net.Conn) {
	log.Println("Addr:", conn.LocalAddr().String(), "RemoteAddr:", conn.RemoteAddr().String())

	for {
		data := make([]byte, 0)
		n, err := conn.Read(data)
		if err != nil {
			log.Panic("Read err:", err, "n:", n)
		}
		log.Println("read n:", n)

	}
}