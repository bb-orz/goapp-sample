package main

import (
	"github.com/bb-orz/goinfras"
	"github.com/bb-orz/goinfras/XCron"
	"github.com/bb-orz/goinfras/XEtcd"
	"github.com/bb-orz/goinfras/XLogger"
	"github.com/bb-orz/goinfras/XMQ/XNats"
	"github.com/bb-orz/goinfras/XMQ/XRedisPubSub"
	"github.com/bb-orz/goinfras/XOAuth"
	"github.com/bb-orz/goinfras/XOss/XAliyunOss"
	"github.com/bb-orz/goinfras/XOss/XQiniuOss"
	"github.com/bb-orz/goinfras/XStore/XGorm"
	"github.com/bb-orz/goinfras/XStore/XMongo"
	"github.com/bb-orz/goinfras/XStore/XRedis"
	"github.com/bb-orz/goinfras/XValidate"
	"github.com/gin-gonic/gin"

	_ "goinfras-sample/restful" // 自动载入Restful API模块
	"github.com/bb-orz/goinfras/XGin"
)

// 注册应用组件启动器，把基础设施各资源组件化
func RegisterStarter() {

	goinfras.RegisterStarter(XLogger.NewStarter())

	// 注册Cron定时任务
	// 可以自定义一些定时任务给starter启动
	goinfras.RegisterStarter(XCron.NewStarter())

	// 注册ETCD
	goinfras.RegisterStarter(XEtcd.NewStarter())

	// 注册mongodb启动器
	goinfras.RegisterStarter(XMongo.NewStarter())

	// 注册mysql启动器
	goinfras.RegisterStarter(XGorm.NewStarter())
	// 注册Redis连接池
	goinfras.RegisterStarter(XRedis.NewStarter())
	// 注册Oss
	goinfras.RegisterStarter(XAliyunOss.NewStarter())
	goinfras.RegisterStarter(XQiniuOss.NewStarter())
	// 注册Mq
	goinfras.RegisterStarter(XNats.NewStarter())
	goinfras.RegisterStarter(XRedisPubSub.NewStarter())
	// 注册Oauth Manager
	goinfras.RegisterStarter(XOAuth.NewStarter())


	// 注册gin web 服务启动器
	// TODO add your gin middlewares
	middlewares := make([]gin.HandlerFunc, 0)
	goinfras.RegisterStarter(XGin.NewStarter(middlewares...))

	// 注册验证器
	goinfras.RegisterStarter(XValidate.NewStarter())

	// 对资源组件启动器进行排序
	goinfras.SortStarters()

}
