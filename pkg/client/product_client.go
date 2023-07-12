package client

import (
	"context"
	"log"

	"github.com/stebin13/order-srv/pkg/config"
	"github.com/stebin13/order-srv/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProdCli(c *config.Config) ProductServiceClient {
	cc, err := grpc.Dial(c.Product_Svc_Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("could not connect", err)
	}

	return ProductServiceClient{
		Client: pb.NewProductServiceClient(cc),
	}
}

func (c *ProductServiceClient) FindOne(productid int64) (*pb.FindOneResponse, error) {
	req := pb.FindOneRequest{
		Id: productid,
	}
	return c.Client.FindOne(context.Background(), &req)
}

func (c *ProductServiceClient) DecreaseStock(productid, orderid int64) (*pb.DecreaseStockResponse, error) {
	req := pb.DecreaseStockRequest{
		Id:      productid,
		OrderId: orderid,
	}
	return c.Client.DecreaseStock(context.Background(), &req)
}
