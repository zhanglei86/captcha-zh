package demo

import (
	"captcha-zh/config"

	"testing"
	"fmt"
	"strconv"
)

const (
	Unknown = iota	// iota 初始化后会自动递增
	Female
	Male
)

/* 声明全局变量 */
var g string


func TestBase(t *testing.T)  {
	fmt.Println("hello, world!")
}

//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
func TestRangeUnicode(t *testing.T)  {
	g = "go"
	for i, c := range g {
		fmt.Println(i, c)
	}
}

// 多返回值函数
func TestMoreRet(t *testing.T)  {
	var flag int
	var sex, callbackOutput string
	sex, callbackOutput, flag = getSex(Female, "数据将原路返回")
	fmt.Printf("性别：%s, 回调：%s, 返回值：%d\n", sex, callbackOutput, flag)
}

// 类型转换
func TestExpression(t *testing.T)  {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum)/float32(count)
	fmt.Printf("mean 的值为: %f\n",mean)
}

// 月份
func TestMonth(t *testing.T)  {

	// test
	fmt.Println(config.Weeks[0])
	fmt.Println(config.Months[1])

	// 循环取
	var i int = 5
	mySlice := make([]int, i)

	// TODO 数组越界时候，如何捕获到异常？？
	for k, v := range mySlice {
		fmt.Printf("k=%d, v=%d\n", k, v)
		fmt.Println(config.Months[k])
	}
}

// 选择
func TestSwitch(t *testing.T)  {

	ret := getSwitchGrade(90)
	fmt.Println(ret)

}

// 获取性别
func getSex(sex int, callBack string) (string, string, int) {
	var retFlag int = 100

	retSex := config.SexConstant[strconv.Itoa(sex)]

	return retSex, callBack, retFlag
}

func getSwitchGrade(store int) string {

	// 定义
	var ret string
	var grade string

	// 算评级
	switch store {
	case 90: grade = "A"
	case 80: grade = "B"
	case 50,60,70 : grade = "C"
	default: grade = "D"
	}

	// 描述
	switch {
	case grade == "A" :
		ret = "优秀!"
	case grade == "B", grade == "C" :
		ret = "良好"
	case grade == "D" :
		ret = "及格"
	case grade == "F":
		ret = "不及格"
	default:
		ret = "差"
	}
	return ret
}
