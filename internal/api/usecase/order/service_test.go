package order

import (
	"errors"
	"reflect"
	"testing"

	pb "github.com/julioshinoda/sample-project/pb/proto"
	"github.com/julioshinoda/sample-project/pb/proto/mocks"
	"github.com/stretchr/testify/mock"
)

func TestService_GetByCustomerID(t *testing.T) {
	type fields struct {
		orderClient pb.OrdersConnectorClient
	}
	type args struct {
		ID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*pb.OrderResponse
		wantErr bool
	}{
		{
			name: "Success to return orders by customer ID",
			fields: fields{
				orderClient: &mocks.OrdersConnectorClient{},
			},
			args: args{
				ID: int64(4711),
			},
			want: []*pb.OrderResponse{
				{Uuid: "fdfd497828effdfda2b06c8b78efa26c"},
			},
		},
		{
			name: "Error to return orders by customer ID",
			fields: fields{
				orderClient: &mocks.OrdersConnectorClient{},
			},
			args: args{
				ID: int64(4711),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				orderClient: tt.fields.orderClient,
			}
			switch tt.name {
			case "Success to return orders by customer ID":
				s.orderClient.(*mocks.OrdersConnectorClient).On(
					"GetOrdersForCustomer",
					mock.Anything,
					&pb.CustomerId{Customerid: int32(tt.args.ID)},
				).Return(
					&pb.Response{
						Success: true,
						Orders: []*pb.OrderResponse{
							{Uuid: "fdfd497828effdfda2b06c8b78efa26c"},
						},
					},
					nil)
			case "Error to return orders by customer ID":
				s.orderClient.(*mocks.OrdersConnectorClient).On(
					"GetOrdersForCustomer",
					mock.Anything,
					&pb.CustomerId{Customerid: int32(tt.args.ID)},
				).Return(
					nil,
					errors.New("some error on server"))
			}
			got, err := s.GetByCustomerID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetByCustomerID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetByCustomerID() = %v, want %v", got, tt.want)
			}
		})
	}
}
