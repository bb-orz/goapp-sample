package core

import (
	"fmt"
	"goapp-sample/dtos"
	"goapp-sample/services"
	"sync"
)

/*
实现服务逻辑
*/


var _ services.IServiceA = new(ServiceAV1)

func init() {
	var once sync.Once
	once.Do(func() {
		services.SetServiceA(new(ServiceAV1))
	})
}

type ServiceAV1 struct{}


func (s *ServiceAV1) Foo(i dtos.InDTO) error {
	var err error
	// TODO 实现模块业务逻辑domain
	fmt.Println("实现模块业务逻辑domain —— Foo")

	return err
}


func (s *ServiceAV1) Bar(i dtos.InDTO) error {
	var err error
	// TODO 实现模块业务逻辑domain
	fmt.Println("实现模块业务逻辑domain —— Bar")
	return err
}
