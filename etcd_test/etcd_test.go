package main

import (
	"fmt"
	"testing"
)

func TestEtcd(t *testing.T) {
	etcdClient, err := initEtcdClient([]string{"http://192.168.0.75:2379"})
	if err != nil {
		t.Fatal("initEtcdClient err:", err)
	}

	apis := etcdClient.Endpoints()
	fmt.Println("apis:", apis)

	etcdClient
}
