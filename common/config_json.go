package common

import(
	"captcha-zh/config"

	"io/ioutil"
	"log"
	"os"
	"encoding/json"
	"sync"
)

type ConfigJson struct {
	InitialCount  int `json:"initial_count"`
	CheckInterval int `json:"check_interval"`
	Threshold     int `json:"threshold"`
	UpdateCount   int `json:"update_count"`
}

var (
	configJson *ConfigJson
	configLock = new(sync.RWMutex)
)

func LoadConfig(flag bool) *ConfigJson {
	file, err := ioutil.ReadFile(config.PATH_CONFIG)
	if err != nil {
		log.Println("open config: ", err)
		if flag {
			os.Exit(1)
		}
	}

	// 临时方法
	temp := new(ConfigJson)

	err2 := json.Unmarshal(file, temp)
	if err2 != nil {
		log.Println("parse config: ", err2)
		if flag {
			os.Exit(1)
		}
	}

	// FIXME 为什么要加锁？！
	configLock.Lock()
	configJson = temp
	configLock.Unlock()

	return configJson
}

func GetConfig() *ConfigJson {
	configLock.RLock()
	defer configLock.RUnlock()
	return configJson
}

func init() {
	LoadConfig(true)
}