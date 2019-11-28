package main

import (
	"context"
	"fmt"
	"git.speakin.mobi/gongan/identify_api_server.git/grpc_dep"
	"github.com/coreos/etcd/client"
	"google.golang.org/grpc"
	"grpc_test/grpc_depn"
	pb "grpc_test/proto"
	"log"
	"sync"
	"time"
)

func main() {

	// etcd建立连接
	etcdClient, err := initEtcdClient([]string{"http://192.168.0.75:2379"})
	if nil != err {
		log.Fatal("etcd init failed ", err)
	}
	grpc_dep.Init(etcdClient, "prod")

	fmt.Println(etcdClient.Endpoints())

	// grpc缓存池
	grpcPool := grpc_depn.NewGrpcClientPool(etcdClient, "/ALGO_EXTRA", 2, time.Second*5)
	conn, err := grpcPool.PickClientWithBalance(grpc.WithInsecure())

	//log.Println("conn: ",conn)
	if err != nil {
		log.Fatalf("grpcPool.PickClientWithBalance err: %v", err)
	}

	client := pb.NewGenderServiceClient(conn)

	resp, err := client.SearchGender(context.Background(), &pb.GenderRequest{
		VoiceId:   "file20190909121150_4a60299a12b543869612db0c4d1c5abb",
		StartTime: 1063,
		EndTime:   10642,
	})

	fmt.Println(resp.GetGender())
	fmt.Println(resp.GetStatus().GetMsg())

	sy := sync.WaitGroup{}
	sy.Add(1)
	sy.Wait()
}

func initEtcdClient(etcdAddrList []string) (client.Client, error) {

	etcdClient, err := client.New(client.Config{
		Endpoints: etcdAddrList,
	})

	if nil != err {
		return nil, err
	}
	go func() {
		for {
			err = etcdClient.AutoSync(context.Background(), 10*time.Second)
			if err == context.DeadlineExceeded || err == context.Canceled {
				break
			}
		}
	}()
	return etcdClient, nil
}
