docker load -i centos-7.9.2009.tgz
执行build.sh脚本构建镜像
```
REPOSITORY             TAG               
app-haproxy            2.2.13 
haproxybase            2.2.13
```
以host模式运行镜像opensigma-haproxy
docker run --name={name} --net=host  -d {imageid}
执行脚本 sh start.sh

haproxy目录下，haproxy文件为2.2.13版本
dataplaneapi目录下，dataplaneapi文件为2.1.0版本
修改版本需下载源码编译出二进制文件进行替换
https://github.com/haproxy/haproxy
https://github.com/haproxytech/dataplaneapi

配置日志：
$ModLoad imudp
$UDPServerRun 514
local0.*                                            /var/log/haproxy.log

