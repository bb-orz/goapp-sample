package main

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"time"
)

const (
	ConfigFilePathFlag = "config_file"

	RemoteProviderFlag      = "remote_provider"
	RemoteEndpointFlag      = "remote_endpoint"
	RemoteKVPathFlag        = "remote_kv_path"
	RemoteTypeFlag          = "remote_type"
	RemoteWatchDurationFlag = "remote_watch_duration"

	EnvPrefixFlag     = "env_prefix"
	EnvKeysFlag       = "env_keys"
	EnvAllowEmptyFlag = "env_allow_empty"
	EnvAutomaticFlag  = "env_automatic"
)

var (
	err      error        // 错误实例
	viperCfg *viper.Viper // viper 配置实例

	configFilePath string // 接收flag command line argument：配置文件完整路径

	remoteProvider      string        // 接收flag command line argument：远程配置提供者名称接收变量
	remoteEndpoint      string        // 接收flag command line argument：远程配置提供者主机端点
	remoteKVPath        string        // 接收flag command line argument： 	远程配置键值路径
	remoteType          string        // 接收flag command line argument： 	远程配置加载到本地的文件类型
	remoteWatchDuration time.Duration // 接收flag command line argument： 	远程配置加载到本地的文件类型

	envAutomatic  bool     //	是否载入所有环境变量，如设置Prefix，则只筛选有前缀的载入
	envAllowEmpty bool     //	是否允许读取空值的环境变量
	envPrefix     string   //	指定需读取的环境变量前缀
	envKeys       []string //  绑定特定环境变量
)

// 运行初始化时解析命令行参数到viper实例
func init() {
	pflag.StringP(ConfigFilePathFlag, "f", "../config/config.yaml", "Config file,like: ../config/config.yaml")

	pflag.StringP(RemoteProviderFlag, "P", "", "Remote K/V config system provider，support etcd/consul")
	pflag.StringP(RemoteEndpointFlag, "E", "", "Remote K/V config system endpoint，etcd requires http://ip:port  consul requires ip:port")
	pflag.StringP(RemoteKVPathFlag, "K", "", "Remote K/V config path，path is the path in the k/v store to retrieve configuration,like: /configs/myapp.json")
	pflag.StringP(RemoteTypeFlag, "T", "", "Support: 'json', 'toml', 'yaml', 'yml', 'properties', 'props', 'prop', 'env', 'dotenv'")
	pflag.DurationP(RemoteWatchDurationFlag, "D", -1, "Currently, only tested with etcd support")

	pflag.BoolP(EnvAutomaticFlag, "a", false, "是否自动读取全部环境变量")
	pflag.BoolP(EnvAllowEmptyFlag, "e", false, "是否允许读取环境变量空值")
	pflag.StringP(EnvPrefixFlag, "p", "", "读取特定前缀的环境变量")
	pflag.StringSliceP(EnvKeysFlag, "k", []string{}, "读取指定键的环境变量")

	pflag.Parse()
	// 实例化viper
	viperCfg = viper.New()
	err := viperCfg.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic("Command Line Flag Binding Error!")
	}
}

func ViperLoader() {
	// 读取viper flag 解析的命令行参数
	configFilePath = viperCfg.GetString(ConfigFilePathFlag)             // 配置文件路径
	remoteProvider = viperCfg.GetString(RemoteProviderFlag)             // 连接远程配置的类型（e.g. etcd/consul）
	remoteEndpoint = viperCfg.GetString(RemoteEndpointFlag)             // 连接远程配置的机器节点（e.g. etcd requires http://ip:port  consul requires ip:port）
	remoteKVPath = viperCfg.GetString(RemoteKVPathFlag)                 // 连接远程配置的配置节点路径 (e.g. path is the path in the k/v store to retrieve configuration)
	remoteType = viperCfg.GetString(RemoteTypeFlag)                     // 远程配置文件类型
	remoteWatchDuration = viperCfg.GetDuration(RemoteWatchDurationFlag) // etcd watch监听时间间隔

	envAutomatic = viperCfg.GetBool(EnvAutomaticFlag) // 是否自动载入所有环境变量（设置Prefix可以过滤特定环境变量）
	envAllowEmpty = viper.GetBool(EnvAllowEmptyFlag)  // 是否允许环境变量空值
	envPrefix = viperCfg.GetString(EnvPrefixFlag)
	envKeys = viperCfg.GetStringSlice(EnvKeysFlag)

	if configFilePath != "" {
		if err = goinfras.LoadViperConfigFromFile(viperCfg, configFilePath); err != nil {
			panic("Viper Loader Config From File Error!")
		}
		fmt.Println("Viper File Config Was Loaded  ......")
	} else if remoteProvider != "" && remoteEndpoint != "" && remoteKVPath != "" {
		if err = goinfras.LoadViperConfigFromRemote(viperCfg, remoteProvider, remoteEndpoint, remoteKVPath, remoteType, remoteWatchDuration); err != nil {
			panic("Viper Loader Config From Remote Error!")
		}
		fmt.Println("Viper Remote Config Was Loaded  ......")
	}

	if len(envKeys) > 0 || envAutomatic {
		if err = goinfras.LoadViperConfigFromEnv(viperCfg, envPrefix, envKeys, envAllowEmpty, envAutomatic, nil); err != nil {
			panic("Viper Loader Config From Env Error!")
		}
		fmt.Println("Viper Env Config Was Loaded  ......")
	}

}
