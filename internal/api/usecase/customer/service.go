package customer

import (
	"context"

	"github.com/julioshinoda/sample-project/pkg/serverconnector"
	"github.com/julioshinoda/sample-project/util"
)

//Service  interface
type Service struct {
	customerClient util.CustomerConnectorClient
}

func NewService(cc util.CustomerConnectorClient) UseCase {
	return &Service{
		customerClient: cc,
	}
}

//GetByEmail Get customer by email
func (s *Service) GetByEmail(email string) ([]*util.CustomerEntity, error) {
	customerResp, err := s.customerClient.GetCustomerByEmail(serverconnector.GetMetadata(context.Background()), &util.Email{Email: email})
	if err != nil {
		return nil, err
	}
	return customerResp.Customers, nil
}
