package demo

/**
递归，就是在运行的过程中调用自己。
Go 语言支持递归。但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。
递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列等。
 */

import (
	"testing"
	"fmt"
)

// 测试
func TestFactorial(t *testing.T)  {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(i))
}

// 测试
func TestFibonacci(t *testing.T)  {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}

// 阶乘
func Factorial(x int) (result int) {
	if x == 0 {
		result = 1;
	} else {
		result = x * Factorial(x - 1);
	}
	return;
}

// 斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
