package main

import (
	"./testProto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
)

type RouteGuide struct{}

//简单的rpc
func (s *RouteGuide) GetFeature(ctx context.Context, req *testProto.Point) (*testProto.Feature, error) {
	log.Printf("简单的rpc GetFeature Latitude:%d Longitude:%d", req.Latitude, req.Longitude)
	rsp := &testProto.Feature{
		Message: []byte("简单的rpc服务"),
		Location: &testProto.Point{Latitude:req.Latitude+1, Longitude:req.Longitude+1},
	}
	return rsp, nil
}

//服务端流式rpc
func (s *RouteGuide) ListFeatures(req *testProto.Rectangle, srv testProto.RouteGuide_ListFeaturesServer) error {
	log.Printf("%s GetFeature Latitude:%d Longitude:%d", string(req.Hi.Message), req.Hi.Latitude, req.Hi.Longitude)
	file, err := os.Open("./testData/16k.wav")
	if err != nil {
		log.Println("ListFeatures Open file err:", err)
		return err
	}

	data := make([]byte, 1024*1024)
	for {
		n, err := file.Read(data)
		if err != nil {
			if err == io.EOF {
				log.Println("Read end n:", n)
				break
			} else {
				log.Println("Read end err:", err)
				return err
			}
		}
		
		_ = srv.Send(&testProto.Feature{Message:data[:n]})
		log.Println("Send data n:", n)
	}

	return nil
}

//客户端流式rpc
func (s *RouteGuide) RecordRoute(srv testProto.RouteGuide_RecordRouteServer) error {
	return nil
}

//双向流式rpc
func (s *RouteGuide) RouteChat(srv testProto.RouteGuide_RouteChatServer) error {
	file, err := os.Open("./testData/16k.wav")
	if err != nil {
		log.Println("ListFeatures Open file err:", err)
		return err
	}

	log.Println("Send...")
	data := make([]byte, 1024*1024)
	for {
		n, err := file.Read(data)
		if err != nil {
			if err == io.EOF {
				log.Println("Read end n:", n)
				break
			} else {
				log.Println("Read err:", err)
				return err
			}
		}

		_ = srv.Send(&testProto.RouteNote{Message:data[:n]})
		log.Println("Send data n:", n)
	}
	_ = srv.Send(&testProto.RouteNote{Message:[]byte("EOF")})

	log.Println("Recv...")
	for {
		data, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("Recv end")
				break
			} else {
				log.Println("Recv err:", err)
				return err
			}
		}
		log.Println("Recv data:", data)
		log.Println("Recv data:", data.Message)
		log.Println("Recv data:", string(data.Message))
	}
	
	return nil
}

func main() {
	server := grpc.NewServer()
	//g_client.RegisterVprServiceServer(server, &SearchService{})
	testProto.RegisterRouteGuideServer(server, &RouteGuide{})
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	fmt.Println("start grpc listen...")
	_ = server.Serve(lis)
}
