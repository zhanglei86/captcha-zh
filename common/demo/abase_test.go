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

// 类型转换
func TestExpression(t *testing.T)  {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum)/float32(count)
	fmt.Printf("mean 的值为: %f\n",mean)
}