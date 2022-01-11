package order

import (
	"context"

	pb "github.com/julioshinoda/sample-project/pb/proto"
	"github.com/julioshinoda/sample-project/pkg/serverconnector"
)

//Service  interface
type Service struct {
	orderClient pb.OrdersConnectorClient
}

func NewService(cc pb.OrdersConnectorClient) UseCase {
	return &Service{
		orderClient: cc,
	}
}

//GetByCustomerID Get orders by customer ID
func (s *Service) GetByCustomerID(ID int64) ([]*pb.OrderResponse, error) {
	customerID := &pb.CustomerId{Customerid: int32(ID)}
	orderResp, err := s.orderClient.GetOrdersForCustomer(serverconnector.GetMetadata(context.Background()), customerID)
	if err != nil {
		return nil, err
	}
	return orderResp.GetOrders(), nil
}
