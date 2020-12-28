package services

import "goapp-sample/dtos"

/*
定义服务及数据传输对象
*/
var serviceA IServiceA

func SetServiceA(sv IServiceA) {
	serviceA = sv
}

func GetServiceA() IServiceA {
	return serviceA
}

type IServiceA interface {
	Foo(i dtos.InDTO) error
	Bar(i dtos.InDTO) error
}

