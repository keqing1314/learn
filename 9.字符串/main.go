// go 字符串记录相关
package main

import (
	"fmt"
	"log"
	"strings"
	"unicode/utf8"
)

// go中字符串本质上就是字节序列
// go官方专门在社区中发布博客来说明这个观点
// 字面意义上，字符串应该是字符序列
// 但是字符是字节组成的，而且英文字符就是一个字节
// 这些字节占用一片连续的内存空间
// 不赘述过多，直接开始示例，感兴趣的话直接去官方网站看博客就行

func main() {
	// 字面量
	// 字符串在go分为原生字符串和普通字符串
	str1 := "这就是一个普通字符串\nqwerty"
	fmt.Println(str1)
	// 普通字符串没办法直接换行，支持通过转义实现换行
	// 上述示例:
	// 这就是一个普通字符串
	// qwerty
	// 原生字符串
	str2 := `我是刻晴
	剑光如我...
	`
	fmt.Println(str2)
	// 我是刻晴
	// 剑光如我...
	// 原生就支持换行，不支持转义符

	// 访问
	// 字符串式字节序列，和字节切片一样，可以使用下标访问
	str3 := "I am Keqing,one of the seven start"
	fmt.Println(string(str3[0])) // I
	// 注意我这里使用的式英文，如果中文字符串，打印就是乱码，因为中文字符多个字节表示
	// 单字节就是乱码
	// 切割字符串
	str3 = str3[0:4] // I am 空格也是字符
	// 但是不能修改
	// 如果修改的话就会编译报错
	// 错误: cannot assign to str3[0] (neither addressable nor a map index expression)
	//str3[0] = 'i'
	// 没办法直接修改某个字节，就可直接覆盖
	str3 = "my name is Keqing"

	// 转换
	// 字符串可以直接转换成字节切片
	// 字节切片可以转换成字符串
	b := []byte(str3)
	fmt.Println(len(b), string(b))
	//【注意】字符串的内容只是可读的，无法修改，但是转换成字节切片就可以修改了

	// 长度
	// len()可以返回字符串的长度
	fmt.Println(len(str3))

	// 拷贝
	// 类似于拷贝数组也是使用copy内置函数
	var dest, src string
	src = "this is a string"
	// copy(目标字符传切片，源头字符串)
	destBytes := make([]byte, len(src))
	copy(destBytes, src)
	dest = string(destBytes)
	fmt.Println(dest)
	// 也可以使用strings包中的Clone函数
	dest = strings.Clone(src)
	fmt.Println(dest)

	// 拼接
	// 字符串拼接最简单就是 +
	hello := "hello "
	world := "world!"
	s5 := hello + world
	fmt.Println(s5)
	// 也可以转换成字节切片
	s5 = string(append([]byte(hello), []byte(world)...))
	// 不过上述方式性能存在问题
	// 使用strings的builder实现拼接
	sb := strings.Builder{}
	sb.WriteString(hello)
	sb.WriteString(world)
	fmt.Println(sb.String())

	// 遍历
	// 前面提到了字符串本质是字节序列
	// 使用for循环遍历
	// 遍历的是字节，打印的也是字节
	for i := 0; i < len(s5); i++ {
		log.Println(string(s5[i]))
	}
	// 使用for range迭代
	// 这个做了处理，遍历的就是rune 也就是字符
	// rune 前面的数据类型中说到了本质就是int32的别名
	// go原生兼容utf8编码，一个utf8合法的字符最大也就4字节
	for _, r := range "我是刻晴" {
		log.Println(string(r))
	}
	// 还可以使用utf8包下的函数
	str := "hello 世界!"
	for i, w := 0, 0; i < len(str); i += w {
		r, width := utf8.DecodeRuneInString(str[i:])
		fmt.Println(string(r))
		w = width
	}
}
