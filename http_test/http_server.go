package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var succeNum int

func main() {
	http.HandleFunc("/upload", UploadFile)
	//http.Handle("/upload", nil)
	log.Println("http listen 30001...")
	if err := http.ListenAndServe("localhost:30001", nil); err != nil {
		log.Panic("http err:", err)
	}
}

func UploadFile(write http.ResponseWriter, request *http.Request) {
	//log.Println("http request method:", request.Method)
	defer request.Body.Close()

	switch request.Method {
	case "GET":
		{
			data, err := ioutil.ReadAll(request.Body)
			if err != nil {
				log.Panic("read body panic:", err)
			}
			log.Println("Get Head:", request.Header, "body:", string(data))

			picture, err := ioutil.ReadFile("./daylight.jpg")
			if err != nil {
				log.Panic("read file panic:", err)
			}
			_, _ = write.Write(picture)
		}
	case "POST":
		go func() {
			//body, err := ioutil.ReadAll(request.Body)
			//if err != nil {
			//	log.Println("post get body err:", err)
			//	return
			//}
			name := fmt.Sprintf("./voice_%v_%v.wav", time.Now().Format("05"), rand.Intn(100000))
			log.Printf("Post file name:%s num:%v", name, succeNum)
			file, err := os.Create(name)
			if err != nil {
				fmt.Println("【ERROR】 create file err:", err)
			}
			defer file.Close()

			data := make([]byte, 1024*1024) //1M
			var num int
			for {
				n, err := request.Body.Read(data)
				if err == io.EOF {
					//log.Println("read body end n:", n)
					_, err = file.Write(data[:n])
					break
				}
				if n == 0 {
					//log.Print("http err n=0")
					break
				}
				num += n
				//log.Println("post body start write n:", n, "num:", num)
				_, err = file.Write(data[:n])
				if err != nil {
					log.Println("【ERROR】 WriteFile err:", err)
					return
				}
			}
			succeNum++
			log.Println("post body write end num:", succeNum)
			//log.Println("request Form:", request.Form, "\n host:", request.Host, "\n Proto", request.Proto, "\n ContentLength", request.ContentLength,
			//	"\n URL:", request.URL, "\n TLS", request.TLS, "\n RequestURI", request.RequestURI)

			_, _ = write.Write([]byte(fmt.Sprintf("POST: success %v!", succeNum)))
		}()
	case "PUT":
		{
			_, _ = write.Write([]byte("PUT: success!"))
		}
	}
}