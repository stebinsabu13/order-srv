package services

import (
	"context"
	"net/http"

	"github.com/stebin13/order-srv/pkg/client"
	"github.com/stebin13/order-srv/pkg/db"
	"github.com/stebin13/order-srv/pkg/models"
	"github.com/stebin13/order-srv/pkg/pb"
)

type Server struct {
	H       db.Handler
	ProdCli client.ProductServiceClient
	pb.OrderServiceServer
}

func (c *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	product, err := c.ProdCli.FindOne(req.Productid)
	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	} else if product.Status == http.StatusNotFound {
		return &pb.CreateOrderResponse{
			Status: product.Status,
			Error:  product.Error,
		}, nil
	} else if product.Data.Stock < req.Quantity {
		return &pb.CreateOrderResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "insufficient stock",
		}, nil
	}

	order := models.Order{
		Price:     product.Data.Price,
		ProductId: req.Productid,
		UserId:    req.Userid,
	}

	if err := c.H.DB.Create(&order).Error; err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to create order",
		}, nil
	}
	res, err := c.ProdCli.DecreaseStock(req.Productid, order.Id)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	} else if res.Status == http.StatusConflict {
		c.H.DB.Delete(&models.Order{}, order.Id)
		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: res.Error}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}
