package http_builder

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	parentCtx = context.Background()
)

type GrpcSrvAddrs struct {
	Address []resolver.Address `json:"address"`
}

func MarshalGrpcSrvAddrs(srvs ...string) []byte{
	addrs := &GrpcSrvAddrs{}
	for _, srv := range srvs {
		addrs.Address = append(addrs.Address, resolver.Address{
			Addr: srv,
		})
	}

	data, _ := json.Marshal(addrs)
	return data
}

func UmarshalGrpcSrvAddrs(data []byte) (*GrpcSrvAddrs, error) {
	addrs := &GrpcSrvAddrs{}
	if err := json.Unmarshal(data, addrs); err != nil{
		return nil, err
	}

	return addrs, nil
}

func SetParentContext(ctx context.Context) {
	parentCtx = ctx
}

// 從遠端服務器拉取grpc服務地址
type httpBuild struct {
}

// 參考代碼： E:\golang\gopath\src\google.golang.org\grpc\resolver\dns\dns_resolver.go
func (hb *httpBuild) ResolveNow(resolver.ResolveNowOption) {
	return
}

func (hb *httpBuild) Close() {
	return
}

// 從遠程獲取服務器列表
func (hb *httpBuild) remoteAddr(remoteTarget string) (*resolver.State, error) {
	c := http.Client{
		Transport: http.DefaultTransport,
	}

	rsp, err := c.Get(remoteTarget)
	if err != nil{
		return nil, err
	} else if rsp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("Get From %s : %v", remoteTarget, rsp)
	}

	defer rsp.Body.Close()
	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil{
		return nil, err
	}

	grpcSrvs, err := UmarshalGrpcSrvAddrs(data)
	if err != nil{
		return nil, err
	}

	return &resolver.State{
		Addresses: grpcSrvs.Address,
	}, nil
}

func (hb *httpBuild) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	var (
		remoteTarget string
		state        *resolver.State
		err          error
	)

	remoteTarget = target.Scheme + `://` + target.Authority + `/` + target.Endpoint
	if state, err = hb.remoteAddr(remoteTarget); err != nil {
		return nil, err
	}

	cc.UpdateState(*state)
	return hb, nil
}

func (*httpBuild) Scheme() string {
	return "http"
}

func DialWithHttpTarget(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	var (
		b resolver.Builder
	)

	b = &httpBuild{}
	resolver.Register(b)

	ctx, cancel := context.WithTimeout(parentCtx, 2*time.Second)
	defer cancel()

	opts = append(opts, grpc.WithBalancerName(roundrobin.Name))
	opts = append(opts, grpc.WithBlock())

	return grpc.DialContext(ctx, target, opts...)
}
