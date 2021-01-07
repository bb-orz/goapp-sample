package services

import "goapp/dtos"

/*
定义服务及数据传输对象
*/
var serviceExample IServiceExample

func SetServiceExample(sv IServiceExample) {
	serviceExample = sv
}

func GetServiceExample() IServiceExample {
	return serviceExample
}

type IServiceExample interface {
	Foo(i dtos.InDTO) error
	Bar(i dtos.InDTO) error
}

