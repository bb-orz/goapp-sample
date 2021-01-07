package main

import (
	"github.com/bb-orz/goinfras"
	"github.com/bb-orz/goinfras/XGin"
	"github.com/bb-orz/goinfras/XGlobal"
	"github.com/bb-orz/goinfras/XLogger"
	"github.com/bb-orz/goinfras/XStore/XGorm"
	"github.com/bb-orz/goinfras/XValidate"
	"github.com/spf13/viper"
	_ "goapp/services"    // 自动载入业务核心，注册service实例
	_ "goapp/restful" // 自动载入Restful API模块
	"goapp/restful/middleware"
)


// 注册应用组件启动器，把基础设施各资源组件化
func RegisterStarter(viperConfig *viper.Viper) {
	goinfras.RegisterStarter(XGlobal.NewStarter())

	goinfras.RegisterStarter(XLogger.NewStarter())

	// 注册mysql启动器
	goinfras.RegisterStarter(XGorm.NewStarter())

	// 注册验证器
	goinfras.RegisterStarter(XValidate.NewStarter())

	// 注册gin web 服务启动器
	// TODO add your gin middlewares
	// 尾部中间件设置为统一错误处理和统一http响应
	goinfras.RegisterStarter(XGin.NewStarter(
		middleware.CorsMiddleware(viperConfig),
		middleware.ResponseMiddleware(),
		middleware.ErrorMiddleware(),
	))

	// 对资源组件启动器进行排序
	goinfras.SortStarters()

}

