package http_builder

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"testing"
	"time"
)

var (
	httpListenSrv = "localhost:20010"
	target        = "http://localhost:20010/srvs"
	srvs          = []string{":31000", ":31001", ":31002"}
)

func testInit() {
	http.HandleFunc("/srvs", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write(MarshalGrpcSrvAddrs(srvs...))
	})
	http.ListenAndServe(httpListenSrv, nil)
}

type greeter struct {
	srv string
}

func (g *greeter) SayHello(ctx context.Context,req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: req.Name + " " + g.srv,
	}, nil
}

func greeterSrv(srvs...string)map[string]*grpc.Server  {
	res := map[string]*grpc.Server{}
	for _, srv := range srvs {
		gSrv := grpc.NewServer()
		helloworld.RegisterGreeterServer(gSrv, &greeter{srv: srv})
		l, _ := net.Listen("tcp", srv)
		go gSrv.Serve(l)
		res[srv] = gSrv
	}

	return res
}

func TestDialWithHttpAddr(t *testing.T) {
	go testInit()

	res := greeterSrv(srvs...)
	cc, err := DialWithHttpTarget("http://192.168.1.108:20004", grpc.WithInsecure())
	if ! assert.NoError(t, err){
		t.SkipNow()
	}

	gc := helloworld.NewGreeterClient(cc)
	for i := 0;i < 1000;i++{
		rsp, _ := gc.SayHello(parentCtx, &helloworld.HelloRequest{
			Name:"Jim Green",
		})

		if rsp != nil{
			log.Println(rsp)
		}
	}
	// 模擬服務崩潰
	go res[srvs[0]].Stop()
	log.Printf("============")
	for i := 0;i < 1000;i++{
		rsp, err := gc.SayHello(parentCtx, &helloworld.HelloRequest{
			Name:"Jim Green",
		})

		if err != nil{
			t.Error(err)
		}

		if rsp != nil{
			log.Println(rsp)
		}
	}

	greeterSrv(srvs[0])
	log.Printf("============")
	time.Sleep(3*time.Second)
	for i := 0;i < 1000;i++{
		rsp, _ := gc.SayHello(parentCtx, &helloworld.HelloRequest{
			Name:"Jim Green",
		})

		if rsp != nil{
			log.Println(rsp)
		}
	}
}
