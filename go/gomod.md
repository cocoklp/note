包管理

https://www.cnblogs.com/wongbingming/p/12941021.html



https://blog.csdn.net/fly910905/article/details/104297446

```
# 开启
export GO111MODULE=on
# 1.13 之后才支持多个地址，之前版本只支持一个
export GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy,direct
# 1.13 开始支持，配置私有 module，不去校验 checksum
export GOPRIVATE=*.corp.example.com,rsc.io/private
```

```
执行命令go mod init在当前目录下生成一个go.mod文件，执行这条命令时，当前目录不能存在go.mod文件。如果之前生成过，要先删除；
如果你工程中存在一些不能确定版本的包，那么生成的go.mod文件可能就不完整，因此继续执行下面的命令；
执行go mod tidy命令，它会添加缺失的模块以及移除不需要的模块。执行后会生成go.sum文件(模块下载条目)。添加参数-v，例如go mod tidy -v可以将执行的信息，即删除和添加的包打印到命令行；
执行命令go mod verify来检查当前模块的依赖是否全部下载下来，是否下载下来被修改过。如果所有的模块都没有被修改过，那么执行这条命令之后，会打印all modules verified。
执行命令go mod vendor生成vendor文件夹，该文件夹下将会放置你go.mod文件描述的依赖包，文件夹下同时还有一个文件modules.txt，它是你整个工程的所有模块。在执行这条命令之前，如果你工程之前有vendor目录，应该先进行删除。同理go mod vendor -v会将添加到vendor中的模块打印出来；
```

