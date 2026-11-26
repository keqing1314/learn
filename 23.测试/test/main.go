// go 测试相关记录
package main

// 对于开发者而言，良好的测试可以提前发现程序中的错误
// 避免后续因维护不及bug不能及时处理的情况出现
// go test是go中提供的测试工具

// 编写规范
// 测试包：最好有一个测试包专门用来测试
// 测试文件：测试文件通常以_test.go结尾
// 测试函数：示例测试ExampleXXX，单元测试TestXXX，基准测试BenchmarkXXX，模糊测试FuzzXXX
// 如果包名是testdata，该包通常是为了存储用于测试的辅助数据
// 执行测试的时候go会忽略testdata的包

// 常用参数
// -o file 指定编译后的二进制文件名称
// -c 只编译测试文件，但是不运行
// -json 以json格式输出测试日志
// -exec xprog 使用xprog运行测试，等价于go run
// -bench regexp 选中regexp匹配的基准测试
// -fuzz regexp 选中regexp匹配的模糊测试
// -fuzztime t 模糊测试自动结束的实践，t为实践间隔，当单位为x时，表示次数，例如220x
// -fuzzminimizetime t	模式测试运行的最小时间，规则同上
// -count n	运行测试 n 次，默认 1 次
// -cover	开启测试覆盖率分析
// -covermode set,count,atomic	设置覆盖率分析的模式
// -cpu	为测试执行GOMAXPROCS
// -failfast	第一次测试失败后，不会开始新的测试
// -list regexp	列出regexp匹配的测试用例
// -parallel n	允许调用了t.Parallel的测试用例并行运行，n值为并行的最大数量
// -run regexp	只运行regexp匹配的测试用例
// -skip regexp	跳过regexp匹配的测试用例
// -timeout d	如果单次测试执行时间超过了时间间隔d，就会panic。d为时间间隔，例 1s,1ms,1ns 等
// -shuffle off,on,N	打乱测试的执行顺序，N为随机种子，默认种子为系统时间
// -v	输出更详细的测试日志
// -benchmem	统计基准测试的内存分配
// -blockprofile block.out	统计测试中协程阻塞情况并写入文件
// -blockprofilerate n	控制协程阻塞统计频率，通过命令go doc runtime.SetBlockProfileRate查看更多细节
// -coverprofile cover.out	统计覆盖率测试的情况并写入文件
// -cpuprofile cpu.out	统计 cpu 情况并写入文件
// -memprofile mem.out	统计内存分配情况并写入文件
// -memprofilerate n	控制内存分配统计的频率，通过命令go doc runtime.MemProfileRate查看更多细节
// -mutexprofile mutex.out	统计锁竞争情况并写入文件
// -mutexprofilefraction n	设置统计n个协程竞争一个互斥锁的情况
// -trace trace.out	将执行追踪情况写入文件
// -outputdir directory	指定上述的统计文件的输出目录，默认为go test的运行目录

func main() {

}
