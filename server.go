package main

import (
	"captcha-zh/config"
	"captcha-zh/captcha"
	"captcha-zh/common/security"
	"captcha-zh/server"

	"net/http"
	"strings"
	"io/ioutil"
)

func CaptchaStream(w http.ResponseWriter, req *http.Request) {
	str, err := captcha.CaptchaContainer.Next()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	strArr := strings.Split(str, "|")
	file, err := ioutil.ReadFile(config.PATH_CONFIG_IMAGE_TEMP + strArr[0])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Write([]byte(security.EncodeFile(file) + "|" + strArr[1]))
}

func main() {
	server.Start()

	http.HandleFunc("/", CaptchaStream)
	http.ListenAndServe(":8002", nil)
}
