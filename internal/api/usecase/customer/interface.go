package customer

import "github.com/julioshinoda/sample-project/util"

//UseCase interface
type UseCase interface {
	GetByEmail(email string) ([]*util.CustomerEntity, error)
}
