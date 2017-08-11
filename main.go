package main

import (
	"captcha-zh/config"
	"captcha-zh/common"
	"captcha-zh/captcha"

	"fmt"
)

/*
测试时候，改成main()
 */
func amain() {
	fmt.Println("hello world!")

	// fun1

	// fun2,config_json
	c := common.GetConfig()
	fmt.Println(*c)

	tConfig := config.TConfig
	fmt.Println(tConfig.CaptchaSys.Initial_count)

	// fun3
	d := captcha.Random(1,10)
	fmt.Println(d)

	d2 := captcha.RandomGen()
	fmt.Println(<-d2)

	//time.Sleep(time.Hour)

	// fun4
	var tpc captcha.Topic = captcha.RandTopic()
	fmt.Println(tpc.Subject, "=", tpc.Result)

	// fun5
	fmt.Println(tConfig.Paths.Path)

}
