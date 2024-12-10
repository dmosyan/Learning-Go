package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "John", City: "Anytown", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "1002", Name: "Emily", City: "New York", Zipcode: "110012", DateofBirth: "2003-04-04", Status: "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}
