[Docker & docker-compose](#Docker & docker-compose)

[修改镜像源](#修改镜像源)

[构建镜像](#构建镜像)

[运行镜像](#运行镜像)



# Docker & docker-compose

## Docker

安装依赖后再安装，rpm -ivh 命令

 

https://blog.csdn.net/sessionsong/article/details/102628738

 

 sudo rpm -ivh docker-ce-18.09.9-3.el7.x86_64.rpm 

 

containerd.io - daemon to interface with the OS API (in this case, LXC - Linux Containers), essentially decouples Docker from the OS, also provides container services for non-Docker container managers

docker-ce - Docker daemon, this is the part that does all the management work, requires the other two on Linux

docker-ce-cli - CLI tools to control the daemon, you can install them on their own if you want to control a remote Docker daemon



## docker-compose

 

https://docs.docker.com/compose/install/#install-compose

 

sudo curl -L "https://github.com/docker/compose/releases/download/1.27.4/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

 

sudo chmod +x /usr/local/bin/docker-compose

 

sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose

# 修改镜像源

 /etc/docker/daemon.json





# docker 命令

## docker ps

https://blog.csdn.net/woshimeihuo/article/details/90209779

```
CONTAINER ID（container id ） ：顾名思义 ,容器ID的意思，可以通过这id找到唯一的对应容器
IMAGE （image）：该容器所使用的镜像
COMMAND （command）：启动容器时运行的命令
CREATED （created）：容器的创建时间，显示格式为”**时间之前创建“
STATUS （status）：容器现在的状态，状态有7种：created（已创建）|restarting（重启中）|running（运行中）|removing（迁移中）|paused（暂停）|exited（停止）|dead
PORTS （ports）:容器的端口信息和使用的连接类型（tcp\udp）
NAMES （names）:镜像自动为容器创建的名字，也唯一代表一个容器
```

![img](https://images2015.cnblogs.com/blog/697113/201609/697113-20160916150735445-515141805.jpg)



# 构建镜像 

## 使用Dockerfile

```
FROM centos:7
WORKDIR /app
USER root
COPY . .

CMD ["./main"]
```

FROM: 基础镜像。

```
package main

import (
	"io"
	"log"
	"net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
```

go build main.go 得到main二进制文件后，docker build -t hellodocker:v1 . 得到镜像

docker images可以查看镜像信息

docker run 创建一个新的容器并运行

docker run命令

```
docker run [OPTIONS] IMAGE [COMMAND] [ARG...]
OPTIONS:
  -d 后台运行容器
  -i 以交互模式运行
  -t 为容器分配一个伪输入终端，it通常同时使用。
  -P 随即映射端口
  -p 指定端口映射 -p 主机port:容器port
  不指定端口的话，容器外不能访问
  -name="" 指定容器名称
  -v 绑定一个卷
```





# 运行镜像

docker run 运行镜像，会根据镜像创建一个容器。实例化。

docker-compose up 可以运行一系列镜像容器。

## docker run

docker run [OPTIONS] IMAGE [COMMAND] [ARG ...]

```
docker run -d --name=test_nginx -p 80:80 -v e:/docker_files:/var/www/html nginx
-d 后台启动nginx
-name 指定别名
-p 宿主机的80端口映射到容器内的80端口
-v 将宿主机的 e:/docker_files目录挂载到容器内的 /var/www/html目录。一般将宿主机的代码目录和配置文件挂载到容器内部，对于数据库而言，数据目录也要挂载，否则容器重启，数据会丢失
挂载相当于映射 -v src:dst  改src，dst会变，改dst src也会变
```



## docker-compose

通过yml文件进行容器的参数设置，管理。

```
version "3.7"
services:
    nginx:#需要启动的容器服务
      image: nginx
      restart: always
      ports:
        - "8080:80"
      container_name: test-nginx
      depends_on:
        - db
    db:#需要启动的容器服务
      image: mysql
      restart: always
      ports:
        - "3307:3306"
      networks:
        - test-net
      environment:
        - MYSQL_ROOT_PASSWORD=123456
      container_name: test-db
      volumes:
        - e:/docker_files/test_mysql:/var/lib/mysql
networks:
    test-net:
        driver: bridge

```

docker-compose up -d



# 日志

docker logs打印出的是docker容器在控制台上输出的信息。





# 报错

##  docker run 报 structure needs cleaning

需清理docker system的文件系统

du -hs /var/lib/docker/ 查看磁盘使用情况

docker system df 查看docker的磁盘使用情况

```
[root@k8s1master ~]# docker system df
TYPE                TOTAL               ACTIVE              SIZE                RECLAIMABLE
Images              10                  10                  1.213GB             47.93MB (3%)
Containers          19                  19                  0B                  0B
Local Volumes       9                   0                   48B                 48B (100%)
Build Cache         0                   0                   0B                  0B
[root@k8s1master ~]# 
```

docker system prune 可以删除关闭的容器无用的数据、卷、网络以及没有tag的镜像

docker system prune -a 会把没有使用的镜像都删掉 **慎用**

迁移docker目录：

https://blog.csdn.net/weixin_32820767/article/details/81196250



## 使用dockerd排查错误

https://www.imooc.com/article/details/id/70557

https://www.jianshu.com/p/93518610eea1

https://zhuanlan.zhihu.com/p/119401647





