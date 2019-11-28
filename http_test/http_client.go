package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

type Reader struct {
	file *os.File
}

func (r *Reader) Open(path string)  {
	var err error
	r.file, err = os.Open(path)
	if err != nil {
		log.Println("ReadFile err:", err)
	}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	//log.Print("Read p len:", len(p), "cap:", cap(p))
	//p = make([]byte, 1024*1024)
	n, err = r.file.Read(p)
	//log.Println("Read file n:", n)
	//log.Print("Read p len:", len(p), "cap:", cap(p))
	//os.Exit(0)
	return
}

func Post(wait *sync.WaitGroup) {
	defer wait.Done()
	client := http.Client{}
	defer client.CloseIdleConnections()

	log.Println("http post start...")
	//bytes.NewBufferString("aa")
	//Read := &Reader{}
	//Read.Open("./voice.wav")
	file, err := os.Open("./voice.wav")
	if err != nil {
		fmt.Println("Open file err:", err)
		return
	}
	//Read.Open("./http_server.go")
	postReq, err := http.NewRequest("POST", "http://localhost:30001/upload", file)
	if err != nil {
		fmt.Println("【ERROR】Request err:", err)
		return
	}
	//postReq.Header.Set("h", "Hello")
	//postReq.Header.Set("w", "World")
	rsp, err := client.Do(postReq)
	if err != nil {
		log.Println("【ERROR】http post err:", err)
		return
	}
	defer rsp.Body.Close()

	//log.Println("post rsp status:", rsp.Status, "statusCode:", rsp.StatusCode)
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Println("【ERROR】post err:", err)
		return
	}
	log.Println("post body:", string(body))
}

func Get()  {
	client := http.Client{}
	defer client.CloseIdleConnections()

	req, err := http.NewRequest("GET", "http://localhost:30001/upload", bytes.NewBufferString("君不见黄河之水天上来！"))
	if err != nil {
		log.Panic("NewRequest panic:", err)
	}
	//req.Header.Set("唐", "李白")
	rsp, err := client.Do(req)
	defer rsp.Body.Close()
	if err != nil {
		log.Panic("Do req panic:", err)
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Panic("read body panic:", err)
	}
	log.Println("body:", string(body))
}

func main()  {
	wait := &sync.WaitGroup{}
	for i:=0; i < 1; i++ {
		wait.Add(1)
		go Post(wait)
	}
	wait.Wait()
}