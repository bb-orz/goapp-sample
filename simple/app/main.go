package main

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"path"
	"path/filepath"
)

var app *goinfras.Application
var viperCfg *viper.Viper  	// viper 配置实例
var configFilePath string 	// 接收flag command line argument：配置文件完整路径
var remoteProvider string 	// 接收flag command line argument：远程配置提供者名称接收变量
var remoteEndpoint string 	// 接收flag command line argument：远程配置提供者主机端点
var remotePath string 		// 接收flag command line argument： 	远程配置提供者起始键路径

// 应用启动时注册资源组件启动器并按启动优先级进行排序
func init() {
	pflag.String("f", "../config/config.yaml", "Config file,like: ../config/config.yaml")
	pflag.String("T", "", "Remote K/V config system provider，support etcd/consul")
	pflag.String("E", "", "Remote K/V config system endpoint，etcd requires http://ip:port  consul requires ip:port")
	pflag.String("P", "", "Remote K/V config path，path is the path in the k/v store to retrieve configuration,like: /configs/myapp.json")
	pflag.Parse()
	// 实例化viper
	viperCfg = viper.New()
	err := viperCfg.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic("Command Line Flag Binding Error!")
	}
}

// 解析command line flag，生成ViperLoader()导入参数
func ViperCfgArgs(viperCfg *viper.Viper) (envCfgArgs *goinfras.EnvConfigArgs,fileCfgArgs *goinfras.FileConfigArgs,remoteCfgArgs *goinfras.RemoteConfigArgs) {
	// 读取viper flag 解析的命令行参数
	configFilePath = viperCfg.GetString("f") // 配置文件路径
	remoteProvider = viperCfg.GetString("T") // 连接远程配置的类型（e.g. etcd/consul）
	remoteEndpoint = viperCfg.GetString("E") // 连接远程配置的机器节点（e.g. etcd requires http://ip:port  consul requires ip:port）
	remotePath     = viperCfg.GetString("P") // 连接远程配置的配置节点路径 (e.g. path is the path in the k/v store to retrieve configuration)

	if configFilePath != "" { // 若读取配置文件，则分解文件名，返回所需要的文件配置参数
		var (
			cPath string 	  // 配置文件路径部分
			cFileName string  // 配置文件文件全名部分
			cName string	  // 配置文件名称部分（不包含扩展名）
			cExt string		  // 配置文件扩展名部分
		)
		cPath, cFileName = filepath.Split(configFilePath) // 分离路径和文件全名
		cExt = path.Ext(cFileName) // 获取包含点符的扩展名
		cExt = cExt[1:] // 去掉点符
		cName = cFileName[0 : len(cFileName)-len(cFileName)] // 获取分离拓展名的文件名
		return nil, &goinfras.FileConfigArgs{Path:cPath,Name:cName,Type:cExt},nil

	} else if remoteProvider != "" && remoteEndpoint != "" &&  remotePath != "" { // 若读取远程配置，返回远程配置参数
		return nil,nil, &goinfras.RemoteConfigArgs{Provider:remoteProvider,Endpoint:remoteEndpoint,Path:remotePath,Type:"json"}
	}

	return
}

func main() {
	var envCfgArgs *goinfras.EnvConfigArgs
	var fileCfgArgs *goinfras.FileConfigArgs
	var remoteCfgArgs *goinfras.RemoteConfigArgs
	envCfgArgs, fileCfgArgs, remoteCfgArgs = ViperCfgArgs(viperCfg)

	if err := goinfras.ViperLoader(viperCfg,envCfgArgs, fileCfgArgs, remoteCfgArgs);err != nil {
		panic("Viper Loader Config Error!")
	}else {
		fmt.Println("Viper Config Was Loaded  ......")
	}

	// 注册应用组件启动器
	fmt.Println("Register Starters  ......")
	RegisterStarter()

	// 创建应用程序启动管理器
	app = goinfras.NewApplication(viperCfg)

	// 运行应用,启动已注册的资源组件
	fmt.Println("Application Start ......")
	app.Up()
}


