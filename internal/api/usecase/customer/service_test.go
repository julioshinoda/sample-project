package customer

import (
	"errors"
	"reflect"
	"testing"

	"github.com/julioshinoda/sample-project/util"
	"github.com/julioshinoda/sample-project/util/mocks"
	"github.com/stretchr/testify/mock"
)

func TestService_GetByEmail(t *testing.T) {
	type fields struct {
		customerClient util.CustomerConnectorClient
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*util.CustomerEntity
		wantErr bool
	}{
		{
			name: "Success to return customer by email",
			fields: fields{
				customerClient: &mocks.CustomerConnectorClient{},
			},
			args: args{
				email: "gehrmann.j@nekom.com",
			},
			want: []*util.CustomerEntity{
				{Customerid: 4711},
			},
		},
		{
			name: "Error to return customer by email",
			fields: fields{
				customerClient: &mocks.CustomerConnectorClient{},
			},
			args: args{
				email: "gehrmann.j@nekom.com",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				customerClient: tt.fields.customerClient,
			}
			switch tt.name {
			case "Success to return customer by email":
				s.customerClient.(*mocks.CustomerConnectorClient).On(
					"GetCustomerByEmail",
					mock.Anything,
					&util.Email{Email: tt.args.email},
				).Return(
					&util.EmailResponse{
						Success: true,
						Customers: []*util.CustomerEntity{
							{Customerid: 4711},
						},
					},
					nil)
			case "Error to return customer by email":
				s.customerClient.(*mocks.CustomerConnectorClient).On(
					"GetCustomerByEmail",
					mock.Anything,
					&util.Email{Email: tt.args.email},
				).Return(
					nil,
					errors.New("some error"),
				)
			}
			got, err := s.GetByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
