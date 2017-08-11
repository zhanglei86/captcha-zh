package demo

/**
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。
不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。
 */

import (
	"testing"
	"fmt"
)

func TestMap1(t *testing.T)  {

	/* 声明变量，默认 map 是 nil */
	var map1 map[string]string

	// 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
	fmt.Println(map1)
	if(map1 == nil) {
		fmt.Println("map1 is null!")
	}

	map1["k1"] = "v1"
}

func TestMap2(t *testing.T)  {
	/* 声明并初始化 */
	countryCapitalMap := make(map[string]string)

	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* 使用 key 输出 map 值 */
	fmt.Println("原始 map")
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}

	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if(ok){
		fmt.Println("Capital of United States is", captial)
	}else {
		fmt.Println("Capital of United States is not present")
	}

	/* 删除元素 */
	delete(countryCapitalMap,"France");
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}

}