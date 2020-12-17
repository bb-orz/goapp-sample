package main

import (
	"github.com/bb-orz/goinfras"
	"github.com/bb-orz/goinfras/XLogger"
	"github.com/bb-orz/goinfras/XOAuth"
	"github.com/bb-orz/goinfras/XStore/XGorm"
	"github.com/bb-orz/goinfras/XStore/XRedis"
	"github.com/bb-orz/goinfras/XValidate"
	"github.com/gin-gonic/gin"

	_ "github.com/bb-orz/goinfras-sample/simple/restful" // 自动载入Restful API模块
	"github.com/bb-orz/goinfras/XGin"
	"io"
	"os"
)

// 注册应用组件启动器，把基础设施各资源组件化
func RegisterStarter() {
	// 注册日志记录启动器，并添加一个异步日志输出到文件
	file, err := os.OpenFile("./info.log", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	writers := []io.Writer{file}
	goinfras.RegisterStarter(XLogger.NewStarter(writers...))

	// 注册Cron定时任务
	// 可以自定义一些定时任务给starter启动
	// goinfras.RegisterStarter(XCron.NewStarter())

	// 注册ETCD
	// goinfras.RegisterStarter(XEtcd.NewStarter())

	// 注册mongodb启动器
	// goinfras.RegisterStarter(XMongo.NewStarter())

	// 注册mysql启动器
	goinfras.RegisterStarter(XGorm.NewStarter())
	// 注册Redis连接池
	goinfras.RegisterStarter(XRedis.NewStarter())
	// 注册Oss
	// goinfras.RegisterStarter(XAliyunOss.NewStarter())
	// goinfras.RegisterStarter(XQiniuOss.NewStarter())
	// 注册Mq
	// goinfras.RegisterStarter(XNats.NewStarter())
	// goinfras.RegisterStarter(XRedisPubSub.NewStarter())
	// 注册Oauth Manager
	goinfras.RegisterStarter(XOAuth.NewStarter())


	// 注册gin web 服务启动器
	// add your gin middlewares
	middlewares := make([]gin.HandlerFunc, 0)
	goinfras.RegisterStarter(XGin.NewStarter(middlewares...))

	// 注册验证器
	goinfras.RegisterStarter(XValidate.NewStarter())

	// 对资源组件启动器进行排序
	goinfras.SortStarters()

}
