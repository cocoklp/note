[构建镜像](#构建镜像)

[运行镜像](#运行镜像)

[docker-compose](#docker-compose)





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

