package main

import (
	"captcha-zh/config"
	"captcha-zh/captcha"

	"fmt"
	"time"
)

func timeTest() {

	//获取时间戳
	tm := time.Now().Unix()

	// TODO fine
	fmt.Println("时间戳：", tm)
	fmt.Println("时间戳：", time.Unix(tm, 0).Format(config.DATE_FORMAT_DEFAULT))
	fmt.Println("当前时间是：", time.Now().Format(config.DATE_FORMAT_DEFAULT))

	fmt.Println("topic", int64(time.Now().Nanosecond()))
}

func main() {
	fmt.Println("hello world!")

	// fun1
	timeTest()

	// fun2,config
	c := config.GetConfig()
	fmt.Println(*c)

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
}
