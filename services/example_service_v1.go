package services

import (
	"fmt"
	"goapp/dtos"
	"sync"
)

/*
实现服务逻辑
*/


var _ IServiceExample = new(ServiceExampleV1)

func init() {
	var once sync.Once
	once.Do(func() {
		SetServiceExample(new(ServiceExampleV1))
	})
}

type ServiceExampleV1 struct{}


func (s *ServiceExampleV1) Foo(i dtos.InDTO) error {
	var err error
	// TODO 实现模块业务逻辑domain
	fmt.Println("实现模块业务逻辑domain —— Foo")

	return err
}


func (s *ServiceExampleV1) Bar(i dtos.InDTO) error {
	var err error
	// TODO 实现模块业务逻辑domain
	fmt.Println("实现模块业务逻辑domain —— Bar")
	return err
}
