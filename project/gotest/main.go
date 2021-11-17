package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main() {
	// fmt.Println("Hello World")

	c := CustomerRepositoryMock{}
	c.On("GetCustomer", 1).Return("Bond", 18, nil)
	c.On("GetCustomer", 2).Return("", 0, errors.New("not found"))

	//=====
	name, age, err := c.GetCustomer(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name, age)
}

type CustomerReposity interface {
	GetCustomer(id int) (name string, age int, err error)
	Hello()
}

type CustomerRepositoryMock struct {
	mock.Mock
}

func (m *CustomerRepositoryMock) GetCustomer(id int) (name string, age int, err error) {
	args := m.Called(id)
	return args.String(0), args.Int(1), args.Error(2)
}

func (m *CustomerRepositoryMock) Hello() {
	println(m)
}
