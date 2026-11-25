// go 模块记录
package main

// 每一个现代编程语言都会有一个属于自己的依赖管理工具
// 例如java中的gradle，python的pip，nodejs的npm等
// 一个好的依赖管理工具可以帮开发者省去不少实践并且可以提升开发效率。
// 然而go早期没有一个成熟的依赖管理方案，那时候可谓是百家争鸣
// 不过我没见识过就是了，我开始学习go的时候go mod就已经出来了，
// 但是当官方再go1.11退出go mod之后，就逐渐主导了go的依赖管理
// 按理说这个早该出了，反正我现在做的项目或者自己写的demo都是go mod管理依赖的

// 编写模块
// go module本质上就是VCS，依旧是版本控制系统，当你在下载依赖的时候
// 实际上执行的是VCS命令，例如git，所以如果你想要分享你编写的库
// 就需要有个能够访问你库的地址或者仓库，我通常使用github来托管我的代码和实现的库
// 这里就不做赘述，自行了解，github用不来的话也不好说自己是程序员了

// 废话不多说，直接上干货

// 设置代理
// 如果没有代理，你下载的依赖那就很费力了，你会科学上网另说
// go虽然没有中央仓库这种东西，但是有个官方的代理仓库。
// 不过这个的服务器部署在国外，国内的效果就不是很好
// 西方不亮东方亮，国内也有做代理的
// 我常用的就是七牛云的代理
// 设置命令如下：
// go env -w GOPROXY=https://goproxy.cn,direct
// direct必须设置，代理下载失败后绕过代理缓存直接尝试网文源代码仓库
// 设置后，后续的下载依赖很快了

// 下载依赖
// 下载依赖很简单
// go get 依赖地址
// 例如现在gin web框架的依赖
// go get github.com/gin-gonic/gin
// 下载过依赖后项目就会多出go.sum文件同时go.mod会有你的依赖信息

// 下载命令工具
// go install命令会将第三方以来下载到本地并编译成二进制文件
// 再linux或者macos中可以再终端中直接以命令的方式运行，windows没试过，应该也可以的

// 模块管理
// 对于go mod 常用命令
// go mod download 下载当前项目的依赖包
// go mod edit 编辑go.mod文件
// go mod graph 输出模块依赖图
// go mod init 在当前目录初始化go mod
// go mod tidy 清理项目模块
// go mod verify 验证项目的依赖合法性
// go mod why 解释项目哪些地方用到了依赖
// go clean -modcache 用于删除项目模块依赖缓存
// go list -m 列出模块

// 模块存储
// 当使用go mod进行项目管理时
// 模块缓存默认存放在$GOPATH/pkg/mod目录下，也可以修改$GOMODCACHE来指定位置
// go env -w GOMODECACHE="指定的模块缓存位置"
// 同意个机器上的所有go module项目共享改目录下的缓存，缓存没有大小限制，且不会自动删除
// 在缓存中解压的依赖源文件都是只读的，清空缓存的命令：
// go clean -modcache
// 在$GOMODECACHE/cache.download目录下存在着依赖的原始文件，包括哈希文件，原始压缩包等，

func main() {}
