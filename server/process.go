package server

import (
	"captcha-zh/config"
	"captcha-zh/captcha"

	"log"
	"time"
	"strings"
	"os"
)

func captchaGenerate(size int) []string {
	s := make([]string, 0)
	for i := 0; i < size; i++ {
		Topic := captcha.RandTopic()
		fileName := captcha.RandomName() + ".gif"
		captcha.Draw(Topic.Subject, config.PATH_CONFIG_IMAGE_TEMP + fileName)
		s = append(s, fileName + config.SEPARATOR_VERTICAL_LINE + Topic.Result)
	}
	return s
}

func Start() {
	c := config.GetConfig()
	captchas := captchaGenerate(c.InitialCount)
	captcha.CaptchaContainer.Append(captchas...)
	log.Print("init success.")
	ticker := time.NewTicker(time.Second * time.Duration(c.CheckInterval))
	go func() {
		for _ = range ticker.C {
			go workder()
		}
	}()
}

func workder() {
	if !captcha.CaptchaContainer.UpdateNeed() {
		return
	}
	captchas := captchaGenerate(config.GetConfig().UpdateCount)
	captcha.CaptchaContainer.Lock()
	oldCaptchas := captcha.CaptchaContainer.Update(captchas...)
	captcha.CaptchaContainer.Unlock()
	for _, captcha := range oldCaptchas {
		fileName := strings.Split(captcha, "|")[0]
		os.Remove(config.PATH_CONFIG_IMAGE_TEMP + fileName)
	}
	log.Print("update suceess.")
}
