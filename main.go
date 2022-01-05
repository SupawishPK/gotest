package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main() {
	c := CustomerRepositoryMock{}
	c.On("GetCustomer", 1).Return("Beer", 25, nil)
	c.On("GetCustomer", 2).Return("", 0, errors.New("not found!"))

	name, age, err := c.GetCustomer(2)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(name, age)
}

type CustomerRepository interface {
	GetCustomer(id int) (name string, age int, err error)
}

type CustomerRepositoryMock struct {
	mock.Mock
}

func (m *CustomerRepositoryMock) GetCustomer(id int) (name string, age int, err error) {
	arge := m.Called(id)
	return arge.String(0), arge.Int(1), arge.Error(2)
}
