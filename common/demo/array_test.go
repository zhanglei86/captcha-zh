package demo

/**
数组，是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。
 */

import (
	"testing"
	"fmt"
)

func TestArr1(t *testing.T)  {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i,j int
	addParam := 10000;

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + addParam /* 设置元素为 i + addParam */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j] )
	}

	var balance1 [10] float32
	var balance2 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Printf("balance1: %f , balance2:%d\n", balance1, balance2[4])

}