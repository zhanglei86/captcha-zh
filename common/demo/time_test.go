package demo

import (
	"captcha-zh/config"

	"testing"
	"time"
	"fmt"
)

func TestTime1(t *testing.T)  {
	//获取时间戳
	tm := time.Now().Unix()

	// TODO fine
	fmt.Println("时间戳：", tm)
	fmt.Println("时间戳：", time.Unix(tm, 0).Format(config.DATE_FORMAT_DEFAULT))
	fmt.Println("当前时间是：", time.Now().Format(config.DATE_FORMAT_DEFAULT))
	fmt.Println("topic", int64(time.Now().Nanosecond()))

	fmt.Println(config.DATE_FORMAT_DEFAULT_LEN)
}