package service

import (
	"github.com/dmosyan/Learning-Go/apis/banking/domain"
	"github.com/dmosyan/Learning-Go/apis/banking/dto"
	"github.com/dmosyan/Learning-Go/apis/banking/errs"
)

//go:generate mockgen -destination=../mocks/service/customer_service_mock.go -package=service . CustomerService
type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "activ" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	cs, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	var csDto []dto.CustomerResponse

	for _, c := range cs {
		resp := c.ToDto()
		csDto = append(csDto, resp)
	}

	return csDto, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	resp := c.ToDto()
	return &resp, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
