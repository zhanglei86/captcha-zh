package demo

import (
	"testing"
	"fmt"
)

func TestBase(t *testing.T)  {
	fmt.Println("hello world!")

	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}