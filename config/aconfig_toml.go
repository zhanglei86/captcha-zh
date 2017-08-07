package config

import (
	"github.com/BurntSushi/toml"
	"time"
	"fmt"
)

/**
全局配置，指针变量
 */
var TConfig *tomlConfig

type tomlConfig struct {
	Title   string
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
	Paths	paths	//路径
	CaptchaSys	captchaSys //图形验证码系统参数
}

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

/**
路径的
 */
type paths struct {
	Path	string
}

/**
图形验证码系统配置
 */
type captchaSys struct {
	Initial_count	int
	Check_interval	int
	Threshold		int
	Update_count	int
}

func init() {
	var config tomlConfig
	if _, err := toml.DecodeFile(config.Paths.Path + PATH_CONFIG_TOML, &config); err != nil {
		fmt.Println(err)
		return
	}

	// 赋值
	TConfig = &config
}