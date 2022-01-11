package order

import (
	pb "github.com/julioshinoda/sample-project/pb/proto"
)

//UseCase interface
type UseCase interface {
	GetByCustomerID(ID int64) ([]*pb.OrderResponse, error)
}
