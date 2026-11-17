// go map相关记录
package main

import (
	"fmt"
	"log"
	"sync"
)

// map映射表在其他编程语言中如java
// map hashmap 之类的 在go中就是map
// go的实现方式是基于哈希，是一种无序结构
// 所以访问顺序就是无序的

func main() {
	// 下面是基本使用方式
	// 初始化
	// 只声明不初始化
	var mp map[int]string
	v, ok := mp[1]
	fmt.Println(v, ok)
	// 这种情况创建的map只能读取，不能赋值，否则编译失败
	// 警告: nil dereference in map update
	// 内函数make
	mp = make(map[int]string)
	mp[1] = "刻晴"
	v, ok = mp[1]
	// 这种创建的是正常的，可以正常使用
	// 字面量
	mp = map[int]string{}
	mp[1] = "刻晴"
	v, ok = mp[1]
	// 这种创建的是正常的可以正常使用
	// 还有种方式使用new函数创建，不过我个人不推荐，map本身就是引用类型
	// 底层就有指针了

	//【注意】map[key]type key必须是可比较类型

	// 访问
	// 前面提到了map使用hash所以是无序的
	// 访问顺序也是无序的
	mp = map[int]string{1: "刻晴", 2: "凝光", 3: "桑多涅"}
	for k, v := range mp {
		log.Println(k, v)
	}
	// 上面遍历map的顺序就不是固定，具有随机性
	// map可以使用len来获取长度，但是没有cap
	fmt.Println(len(mp)) // 3

	// 使用
	// map使用看起来和数组切片一样
	// 使用key来读取数据
	user1 := mp[1]
	fmt.Println(user1)
	// 使用key来修改数据
	mp[1] = "胡桃"

	// 删除
	// map中提供delete函数
	// delete(要删除元素的map,要删除的元素的key)
	delete(mp, 1)

	// 遍历
	// map使用for range 来迭代
	for k, v := range mp {
		log.Println(k, v)
	}

	// 清空
	// 1.21版本之前，没有clear
	// 想要清空，就遍历一个一个delete掉
	for k := range mp {
		delete(mp, k)
	}
	// 1.21之后clear就行
	clear(mp)

	// Set 是一种无序，且值唯一的数据结构
	// 在其他编程语言中，往往需要单独设计封装
	// 在go中你会发现，完美和map契合
	// map的key就是唯一的，且map就是无序的
	// 实现set就直接使用map就行
	set := map[int]struct{}{}
	fmt.Println(set)
	// 值类型使用空结构体，是因为空结构体不占用内存空间

	//【注意】map在go中不是并发安全的，但是go中提供并发安全的map,在sync包下
	safemp := &sync.Map{}
	safemp.Store(1, "刻晴")
	value, ok := safemp.Load(1)
	if ok {
		v, _ = value.(string)
	}
	safemp.Clear()
	// 除了上面写的方法，还提供其他的方法，自行查看
}
