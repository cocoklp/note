# golang unrecognized import path 完美解决方案.

https://blog.csdn.net/yuhao818/article/details/100557931

## go get层面增加代理

go 1.11版本新增了 GOPROXY 环境变量，go get会根据这个环境变量来决定去哪里取引入库的代码

```
$ export GOPROXY=https://goproxy.io
```