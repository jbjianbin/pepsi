package main

import (
	"context"
	"flag"
	"github.com/saileifeng/pepsi/example/shopping/name"
	"github.com/saileifeng/pepsi/example/shopping/pb"
	"github.com/saileifeng/pepsi/registry/consul"
	"log"
	"time"
)

var consulAddr = "127.0.0.1:8500"
var port = 0
var serviceName = name.SrvOrder

//OrderService ...
type OrderService struct {
}

//CreateOrder 创建购物订单
func (os *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderInfoRequest) (*pb.CreateOrderInfoResponse, error) {
	log.Println("CreateOrder :", req)
	return &pb.CreateOrderInfoResponse{OrderID: time.Now().Unix()}, nil
}

func main() {
	flag.StringVar(&consulAddr, "registry_address", "127.0.0.1:8500", "registry address")
	flag.Parse()
	r := consul.NewRegister(consulAddr, serviceName, port)
	pb.RegisterOrderServiceServer(r.Server, &OrderService{})
	r.Run()
}
