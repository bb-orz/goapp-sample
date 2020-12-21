package restful

import (
	"fmt"
	"github.com/bb-orz/goinfras-sample"
	"github.com/bb-orz/goinfras/XGin"
	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化时自动注册该API到Gin Engine
	XGin.RegisterApi(new(SimpleApi))
}

type SimpleApi struct {
	service1 goinfras_sample.IService1
}

// SetRouter由Gin Engine 启动时调用
func (s *SimpleApi) SetRoutes() {
	s.service1 = goinfras_sample.GetService1()

	engine := XGin.XEngine()

	engine.GET("simple/foo", s.Foo)
	engine.GET("simple/bar", s.Bar)
}

func (s *SimpleApi) Foo(ctx *gin.Context) {
	email := ctx.Param("email")
	// 调用服务
	err := s.service1.Foo(goinfras_sample.InDTO{Email: email})

	// 处理错误
	fmt.Println(err)
}

func (s *SimpleApi) Bar(ctx *gin.Context) {
	email := ctx.Param("email")
	// 调用服务
	err := s.service1.Bar(goinfras_sample.InDTO{Email: email})

	// 处理错误
	fmt.Println(err)
}
