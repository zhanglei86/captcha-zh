package demo

/**
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
 */

import (
	"testing"
	"fmt"
)

type NokiaPhone struct {}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type HuaweiPhone struct{}

func (hwPhone HuaweiPhone) call() {
	fmt.Println("I am huaweiPhone, I can call you!")
}

func TestPhone(t *testing.T)  {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(HuaweiPhone)
	phone.call()
}
