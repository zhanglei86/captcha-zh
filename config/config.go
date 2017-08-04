package config

import(
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
	"sync"
)

type Config struct {
	InitialCount  int `json:"initial_count"`
	CheckInterval int `json:"check_interval"`
	Threshold     int `json:"threshold"`
	UpdateCount   int `json:"update_count"`
}

var (
	config *Config
	configLock = new(sync.RWMutex)
)

func LoadConfig(flag bool) *Config {
	file, err := ioutil.ReadFile(PATH_CONFIG)
	if err != nil {
		log.Println("open config: ", err)
		if flag {
			os.Exit(1)
		}
	}

	// 临时方法
	temp := new(Config)

	err2 := json.Unmarshal(file, temp)
	if err2 != nil {
		log.Println("parse config: ", err2)
		if flag {
			os.Exit(1)
		}
	}

	// FIXME 为什么要加锁？！
	configLock.Lock()
	config = temp
	configLock.Unlock()

	return config
}

func GetConfig() *Config {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func init() {
	LoadConfig(true)
}