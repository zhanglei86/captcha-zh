package common

import(
	"captcha-zh/config"

	"io/ioutil"
	"log"
	"os"
	"encoding/json"
	"sync"
	"os/signal"
	"syscall"
)

type ConfigJson struct {
	InitialCount  int `json:"initial_count"`
	CheckInterval int `json:"check_interval"`
	Threshold     int `json:"threshold"`
	UpdateCount   int `json:"update_count"`
}

var (
	configJson *ConfigJson	// FIXME 定义错误，定义的要大写，这样初始化一次，不用每次都load()
	configLock = new(sync.RWMutex)
)

func loadConfig(flag bool) *ConfigJson {
	file, err := ioutil.ReadFile(config.TConfig.Paths.Path + config.PATH_CONFIG_JSON)
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
	loadConfig(true)

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR2)
	go func() {
		for {
			<-s
			loadConfig(false)
			log.Print(GetConfig())
			log.Println("Reloaded")
		}
	}()
}