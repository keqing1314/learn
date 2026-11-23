// go 文件相关示例
package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
	"syscall"
)

// Go 语言中提供文件处理的标准库大致有以下几个:
// os库 负责os文件系统交互的具体实现
// io库 读写io的抽象层
// fs库 文件系统的抽象层

// 打开
// 常见的两种打开文件的凡是使用os包提供的两个函数，
// Open函数
func OpenDemo() {
	// 直接输入文件名就行
	// 你看源码的话，发现里面就是OpenFile的调用
	file, err := os.Open("森林书.txt")
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
		return
	} else if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	fmt.Printf("文件打开成功，文件名%s", file.Name())
}

// OpenFile函数
func OpenFileDemo() {
	// 这个可以设置文件描述和文件权限
	// 权限os包中定义这些，我这里复制过来
	const (
		// 只读，只写，读写 三种必须指定一个
		O_RDONLY int = syscall.O_RDONLY // 以只读的模式打开文件
		O_WRONLY int = syscall.O_WRONLY // 以只写的模式打开文件
		O_RDWR   int = syscall.O_RDWR   // 以读写的模式打开文件
		// 剩余的值用于控制行为
		O_APPEND int = syscall.O_APPEND // 当写入文件时，将数据添加到文件末尾
		O_CREATE int = syscall.O_CREAT  // 如果文件不存在则创建文件
		O_EXCL   int = syscall.O_EXCL   // 与O_CREATE一起使用, 文件必须不存在
		O_SYNC   int = syscall.O_SYNC   // 以同步IO的方式打开文件
		O_TRUNC  int = syscall.O_TRUNC  // 当打开的时候截断可写的文件
	)
	// 文件权限的
	const (
		ModeDir        = fs.ModeDir        // d: 目录
		ModeAppend     = fs.ModeAppend     // a: 只能添加
		ModeExclusive  = fs.ModeExclusive  // l: 专用
		ModeTemporary  = fs.ModeTemporary  // T: 临时文件
		ModeSymlink    = fs.ModeSymlink    // L: 符号链接
		ModeDevice     = fs.ModeDevice     // D: 设备文件
		ModeNamedPipe  = fs.ModeNamedPipe  // p: 具名管道 (FIFO)
		ModeSocket     = fs.ModeSocket     // S: Unix 域套接字
		ModeSetuid     = fs.ModeSetuid     // u: setuid
		ModeSetgid     = fs.ModeSetgid     // g: setgid
		ModeCharDevice = fs.ModeCharDevice // c: Unix 字符设备, 前提是设置了 ModeDevice
		ModeSticky     = fs.ModeSticky     // t: 黏滞位
		ModeIrregular  = fs.ModeIrregular  // ?: 非常规文件

		// 类型位的掩码. 对于常规文件而言，什么都不会设置.
		ModeType = fs.ModeType

		ModePerm = fs.ModePerm // Unix 权限位, 0o777
	)
	// 读写，没有就创建
	file, err := os.OpenFile("森林书.txt", os.O_RDWR|os.O_CREATE, 0444)
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
		return
	} else if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	file.Close()
	fmt.Printf("文件打开成功，文件名%s", file.Name())
}

// 读取
// 当成功的打开文件后，便可以进行读取操作了
// 关于读取文件的操作，*os.File类型提供了以下几个公开的方法
// 将文件读进传入的字节切片
// func (f *File) Read(b []byte) (n int, err error)
// 相较于第一种可以从指定偏移量读取
// func (f *File) ReadAt(b []byte, off int64) (n int, err error)

// 一般都是直接使用os包中封装好的
// os.ReadFile()函数
func ReadFileDemo() {
	bytes, err := os.ReadFile("森林书.txt")
	if err != nil {
		fmt.Printf("文件读取失败:%v", err)
		return
	}
	fmt.Println(string(bytes)) // 刻晴的
}

// io.ReadAll() 函数
func ReadAllDemo() {
	file, err := os.OpenFile("森林书.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("打开文件失败:%v", err)
		return
	}
	defer file.Close()

	fmt.Printf("\n文件打开成功:%v", file.Name())
	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("文件读取失败:%v", err)
	}
	fmt.Println(string(bytes))
}

// 写入
// os.File结构体提供了以下几种方法以供写入数据
// 写入字节切片
// func (f *File) Write(b []byte) (n int, err error)

// 写入字符串
// func (f *File) WriteString(s string) (n int, err error)

// 从指定位置开始写，当以os.O_APPEND模式打开时，会返回错误
// func (f *File) WriteAt(b []byte, off int64) (n int, err error)

// 对一个文件写入数据，那么访问修饰符应当是os.O_WRONLY或者os.O_RDWR
// 否则就无法成写入数据
// 这是一个简单的写入文件内容示例
// 打开要写入的文件，没有文件就创建,有就追加，有读写权限
// 向文件里面写入5句话
func WriteDmeo() {
	file, err := os.OpenFile("系统指南.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("文件打开失败:%v", err)
		return
	}
	defer file.Close()

	for range 5 {
		offest, err := file.WriteString("刻晴天下第一!\n")
		if err != nil {
			fmt.Printf("错误:%v", err)
			break
		}
		fmt.Println(offest)
	}
}

// os.WriteFile() 函数
// func os.WriteFile(name string, data []byte, perm os.FileMode) error
func WriteFileDemo() {
	bytes := []byte("hello world!\n")
	err := os.WriteFile("README.txt", bytes, 0666)
	if err != nil {
		fmt.Println(err)
	}
}

// io.WriteString() 函数
func WriteStringDemo() {
	flie, err := os.OpenFile("刻晴圣经.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("打开文件错误%v", err)
		return
	}

	offset, err := io.WriteString(flie, "刻晴晴晴啊啊啊啊啊!")
	if err != nil {
		fmt.Printf("写入文件失败:%v", err)
		return
	}
	fmt.Println(offset)
}

// os.Create() 函数
// 其实就是调用os.OpenFile，简单封装了下
func CreateDemo() {
	//【注意】这里只能创建文件，不能创建目录
	// 如果目录存在可以带目录，没有的话就会panic
	file, err := os.Create("玉衡星")
	if err != nil {
		fmt.Printf("创建文件失败:%v", err)
	}
	fmt.Println("文件创建成功", file.Name())
	_, err = io.WriteString(file, "我是玉衡星刻晴")
	if err != nil {
		fmt.Printf("写入文件失败:%v", err)
	}
}

// 复制
// 对于复制文件，其操作就是同时打开两个文件，一个是源文件一个是要复制的目标文件
// 读取源文件的内容，写入到目标文件中
func CopyDemo() {
	// 从源文件中读取内容
	data, err := os.ReadFile("系统指南.txt")
	if err != nil {
		fmt.Printf("读取文件失败:%v", err)
		return
	}

	// 向目标文件中写入数据
	err = os.WriteFile("系统指南复制.txt", data, 0666)
	if err != nil {
		fmt.Printf("文件写入失败%v", err)
		return
	}
}

// *os.File.ReadFrom
// 使用os.File中的方法ReadFrom，打开文件时，一个只读，一个只写
// 这种复制方式需要将源文件的全部内容读取到内存中却，然后写入目标文件
// 文件较大就不建议使用这种方式
func ReadFromDemo() {
	// 只读的方式打开源文件
	origin, err := os.OpenFile("森林书.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer origin.Close()

	// 只写的方式打开复制文件
	target, err := os.OpenFile("森林书复制.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer target.Close()

	// 复制文件
	_, err = target.ReadFrom(origin)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// io.Copy() 函数
// 这种复制方式是边读边写，先将内容读取到缓冲区去，缓冲区默认是32KB
// func Copy(dst Writer, src Reader) (writtem int64, err error)
// 使用起来和上面的os.File.ReadFile没啥区别
func IOCopyDemo() {
	// 只读，打开文件
	dst, err := os.OpenFile("玉衡星", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("打开源文件失败%v", err)
		return
	}
	defer dst.Close()

	// 只写，打开文件
	src, err := os.OpenFile("玉衡星.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("打开目标文件失败:%v", err)
		return
	}
	defer src.Close()

	// 复制操作
	offset, err := io.Copy(dst, src)
	if err != nil {
		fmt.Printf("复制失败:%v", err)
		return
	}
	fmt.Println("文件复制成功", offset)
}

// 重命名
// os.Rename() 函数
func RenameDemo() {
	err := os.Rename("玉衡星", "玉衡星刻晴")
	if err != nil {
		fmt.Printf("重命名文件失败%v", err)
		return
	}
}

// 删除
// os.Remove和os.RemoveAll
// 前者是删除文件或者空目录，后者是删除指定目录的所有文件和子目录
func RemoveDemo() {
	err := os.Remove("玉衡星")
	if err != nil {
		fmt.Printf("删除失败%v", err)
		return
	}
}

func RemoveAllDemo() {
	// 删除当前目录下所有的文件与子目录
	err := os.RemoveAll(".")
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 刷新
// os.File.Sync这一个函数封装了底层的系统调用Fsync
// 用于将操作系统中缓存的IO写入落实到磁盘上去
func SyncDmeo() {
	file, err := os.Create("测试刷新")
	if err != nil {
		fmt.Printf("文件创建失败%v", err)
		return
	}
	defer file.Close()

	// 写入模拟数据
	_, err = io.WriteString(file, "我是刻晴")
	if err != nil {
		fmt.Printf("文件写入失败%v", err)
		return
	}

	// 同步刷新
	if err := file.Sync(); err != nil {
		fmt.Printf("文件刷新失败%v", err)
		return
	}
}

// 文件夹
// 1、读取文件夹
func ReadDirDemo() {
	d, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("文件读取失败%v", err)
		return
	}

	// 打印目录所有文件文件
	for _, f := range d {
		fmt.Println(f.Name())
	}
}

func ReadDirDmeo2() {

	var readDir func(path string)
	readDir = func(path string) {
		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}
		file, err := os.Open(path)
		if err != nil {
			fmt.Printf("文件打开失败%v\n", err)
			return
		}
		defer file.Close()

		dirs, err := file.ReadDir(-1)
		if err != nil {
			fmt.Printf("读取失败")
		}

		for _, f := range dirs {
			if f.IsDir() {
				fmt.Println("目录", f.Name())
				readDir(path + f.Name())
			} else {
				fmt.Println("文件", f.Name())
			}
		}
	}

	readDir(".")

}

// 创建
// os包下面的Mkdir和MkdirAll
// 使用这两个函数创建文件
func MkdirDemo() {
	err := os.Mkdir("测试目录", 0666)
	if err != nil {
		fmt.Printf("创建目录失败%v", err)
		return
	}
}

func main() {
	//WriteDmeo()
	//CreateDemo()
	ReadDirDmeo2()
	// MkdirDemo()
}
