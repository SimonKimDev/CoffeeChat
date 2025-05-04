package application

import "github.com/SimonKimDev/CoffeeChat/internal/domain"

type Greeter interface {
	Greet() domain.Greeting
}

type greeterSvc struct{}

func NewGreeterService() Greeter {
	return &greeterSvc{}
}

func (*greeterSvc) Greet() domain.Greeting {
	return domain.Greeting{Message: "Hello World"}
}
