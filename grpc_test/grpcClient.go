package main

import (
	"./testProto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("grpc client start...")
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := testProto.NewRouteGuideClient(conn)
	//ListFeatures(client)
	RouteChat(client)
}

//服务端流式rpc
func ListFeatures(client testProto.RouteGuideClient) {
	resp, err := client.ListFeatures(context.Background(), &testProto.Rectangle{Hi:&testProto.Point{Message:[]byte("服务端流式rpc"),
		Longitude:1, Latitude:5}})
	if err != nil {
		log.Fatalf("服务端流式rpc client.ListFeatures err: %v", err)
	}

	file, err := os.Create("./testData/recv_16k.wav")
	if err != nil {
		log.Println("Create err:", err)
		return
	}
	for {
		rsp, err := resp.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("Recv end")
				break
			} else {
				log.Println("Recv err:", err)
				return
			}
		}
		n, err := file.Write(rsp.Message)
		if err != nil {
			log.Println("Write err:", err)
			return
		}
		log.Println("Recv n:", n)
	}
}

//双向流式rpc
func RouteChat(client testProto.RouteGuideClient)  {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := client.RouteChat(ctx)
	if err != nil {
		log.Fatalf("双向流式rpc client.ListFeatures err: %v", err)
	}

	file, err := os.Create("./testData/recv_16k.wav")
	if err != nil {
		log.Println("Create err:", err)
		return
	}
	
	log.Println("Recv...")
	for {
		rsp, err := resp.Recv()

		if err == io.EOF {
			log.Println("Recv end")
			break
		}
		if err != nil {
			log.Println("Recv err:", err)
			return
		}
		if string(rsp.Message) == "EOF" {
			log.Println("Recv end..")
			break
		}

		n, err := file.Write(rsp.Message)
		if err != nil {
			log.Println("Write err:", err)
			return
		}
		log.Println("Recv n:", n)
	}

	log.Println("Send...")
	err = resp.Send(&testProto.RouteNote{Message:[]byte("文件传输成功！")})
	if err != nil {
		log.Println("Send err:", err)
		return
	}
	
	//resp.CloseSend()
	//cancel()
	//<-resp.Context().Done()
}