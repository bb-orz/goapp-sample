package services

import (
	"fmt"
	"github.com/bb-orz/goinfras-sample"
	"sync"
)

/*
实现服务逻辑
*/

var _ goinfras_sample.IService1 = new(Service1)

func init() {
	var once sync.Once
	once.Do(func() {
		goinfras_sample.SetService1(new(Service1))
	})
}

type Service1 struct{}

func (s *Service1) Foo(i goinfras_sample.InDTO) error {
	var err error
	// TODO 实现模块业务逻辑domain
	fmt.Println("实现模块业务逻辑domain —— Foo")

	return err
}

func (s *Service1) Bar(i goinfras_sample.InDTO) error {
	var err error
	// TODO 实现模块业务逻辑domain
	fmt.Println("实现模块业务逻辑domain —— Bar")
	return err
}
