# kong

## 官网

https://docs.konghq.com/

https://docs.konghq.com/enterprise/2.3.x/admin-api/#service-object



## 安装

### 源码安装(失败)

安装依赖包

https://docs.konghq.com/install/source/

准备：

 yum -y install patch

yum install -y zlib-devel

yum install -y gcc-c++

yum install libyaml-devel libyaml

git clone git@github.com:Kong/kong-build-tools.git

./kong-build-tools/openresty-build-tools/kong-ngx-build     --prefix deps     --work work     --openresty 1.17.8.2     --openssl 1.1.1i     --kong-nginx-module master     --luarocks 3.4.0     --pcre 8.44     --jobs 6     --force

export OPENSSL_DIR**=**$(pwd)/deps/openssl 
export PATH**=**$(pwd)/deps/openresty/bin:$PATH 
export PATH**=**$(pwd)/deps/openresty/nginx/sbin:$PATH 
export PATH**=**$(pwd)/deps/openssl/bin:$PATH
export PATH**=**$(pwd)/deps/luarocks/bin:$PATH

检查依赖

nginx -V 
resty -v 
openresty -V
openssl version -a 
luarocks --version

安装kong

git clone git@github.com:Kong/kong.git

cd kong

git checkout 2.3.2

make install

 export PATH**=**$(pwd)/kong/bin:$PATH

kong version --vv 检查版本号

kong命令报错

 module '****' not found

安装libyaml：yum install -y libyaml

https://www.cnblogs.com/zhoujie/p/kong2.html
https://www.cnblogs.com/chenjinxi/p/8724564.html



### docker 安装(成功)

https://registry.hub.docker.com/_/kong

docker run -d --name kong-database                 -p 5432:5432                 -e "POSTGRES_USER=kong"   -e POSTGRES_HOST_AUTH_METHOD=trust"              -e "POSTGRES_DB=kong"                 postgres:9.6

docker run --rm     --link kong-database:kong-database     -e "KONG_DATABASE=postgres"     -e "KONG_PG_HOST=kong-database"     -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database"     kong kong migrations bootstrap

docker run -d --name kong1     --link kong-database:kong-database     -e "KONG_DATABASE=postgres"     -e "KONG_PG_HOST=kong-database"     -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database"     -e "KONG_PROXY_ACCESS_LOG=/dev/stdout"     -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout"     -e "KONG_PROXY_ERROR_LOG=/dev/stderr"     -e "KONG_ADMIN_ERROR_LOG=/dev/stderr"     -e "KONG_ADMIN_LISTEN=0.0.0.0:8001, 0.0.0.0:8444 ssl"     -p 8000:8000     -p 8443:8443     -p 8001:8001     -p 8444:8444     kong



调kong的api  主要涉及consumer、auth key / hmac 、acl 这几个插件

## 原理

proxy_listen：

​	默认8000，接收公网流量并代理到upstream

admin_listen：

​	默认8001，配置

### 术语

client：客户端，访问kong proxy 端口的下游服务器

upstream service：上游服务器，客户端请求转发到的地方。

service：服务实体。上有服务器抽象。

route：kong的路由实体。kong的入口。根据请求匹配已定义好的规则，路由给关联的service，每个route都关联一个service

plugin：插件。通过admin api配置 全局、routes、services维度的插件。

### 流程

有请求到达后，与配置的route比较看是否能匹配，每个route都会关联service，kong会应用配置到route和service的插件，然后把流量代理到upstream。不匹配则返回404

hosts  paths  methods 请求同时满足三个元素则匹配成功。host支持泛域名。

8000：监听HTTP请求并将请求转发到上游服务器。

8443：监听http请求但不转发。

8001：admin api

8444：监控

## service

服务实体，上游服务器的抽象。一个service可以有多个route，匹配到的route就会转发到service里。可以执行一个target(物理服务)或upstream。

```
{
    "id": "9748f662-7711-4a90-8186-dc02f10eb0f5",
    "created_at": 1422386534,
    "updated_at": 1422386534,
    "name": "my-service",
    "retries": 5,
    "protocol": "http",
    "host": "example.com",
    "port": 80,
    "path": "/some_api",
    "connect_timeout": 60000,
    "write_timeout": 60000,
    "read_timeout": 60000,
    "tags": ["user-level", "low-priority"],
    "client_certificate": {"id":"4e3ad2e4-0bc4-4638-8e34-c84a417ba39b"}
}
```



## routes

路由的抽象，客户端的入口。将实际的request映射到service。每个route都会关联一个service，一个service可以关联多个router。



## consumers

消费者，使用service的用户。可以为其添加plugin插件，自定义他的请求行为。可以创建多个app

## upstream

虚拟服务，多个targets集合，可负载均衡。

service的host可以指定为upstream对象，service的host就是upstream的name。

```
--添加upstream
curl -i -X POST --url http://localhost:8001/upstreams/ --data 'name=nhs.wilmar.service'
```



## target

每个终端节点作为一个target。



## admin api

### service

新增service

```
curl -i -X POST \
  --url  http://localhost:8001/services/ \
  --data 'name=service-stock' \
  --data 'url=http://hq.sinajs.cn'
或
  curl -i -X POST --url  http://localhost:8001/services/ --data 'name=service-stock' --data 'protocol=http' --data 'host=hq.sinajs.cn'   --data 'port=80'
```

获取service

根据service-name或者route-id

```
curl -X GET http://127.0.0.1:8001/services/773e3a07-fb7b-4459-93ed-b9fda63cd037
```

根据routerid获取service

```
GET /routes/{route id}/service
```

### router

定义匹配客户端请求的规则，匹配给定路由的每个请求将被代理到与之关联的服务。

```
curl -i -X POST --url  http://localhost:8001/routes/ --data 'protocols[]=http&protocols[]=https' --data 'hosts[]=hq.sinajs.cn'   --data 'service.id=9ec3c166-f29a-4b04-a33e-c17ac42a3429'
```

添加了路由，并且关联了一个service。根据host匹配。

```
#原接口
curl http://hq.sinajs.cn/list=sh601006

#通过kong代理访问
curl 127.0.0.1:8000/list=sh601006 --header 'host:hq.sinajs.cn'
```

服务关联的路由

```
GET /services/{service name or id}/routes
```



### consumer object

服务的消费者或用户。依赖kong作为主要数据存储，也可以将用户自己管理的列表映射到该数据库consumer表。权限控制依赖这个表。

创建消费者

```
 curl -i -X POST   --url  http://localhost:8001/consumers/   --data 'username=zhou'   --data 'custom_id=000150' 
```



### plugin object

在http请求/响应生命周期期间执行的插件配置。身份验证、速率限制。

客户端对服务每个请求都运行添加的插件，如果需要针对消费者，可以指定consumer_id实现。

一个插件对于一个请求只运行一次。可以为各种实体 实体组合 全局配置插件。

```
当多次配置插件时，优先级的完整顺序是：
1、在以下组合上配置插件：a Route, a Service, and a Consumer。 （消费者意味着请求必须被认证）。
2、在Route和Consumer组合上配置的插件。 （消费者意味着请求必须被认证）。
3、插件配置在Service和Consumer的组合上。 （消费者意味着请求必须被认证）。
4、插件配置在Route和Service的组合上。
5、在Consumer上配置的插件。 （消费者意味着请求必须被认证）。
6、在Route上配置的插件。
7、在Service上配置的插件。
8、插件配置为全局运行。
```



service:

```
 curl -X GET http://127.0.0.1:8001/services/httpbin
{"host":"10.12.4.71","created_at":1612783966,"connect_timeout":60000,"id":"d4724441-447e-4bb7-a373-7af2c4bd32be","protocol":"http","name":"httpbin","read_timeout":60000,"port":80,"path":"\/anything\/req","updated_at":1612812973,"retries":5,"write_timeout":60000,"tags":null,"client_certificate":null}
```



route for service:

## 插件

### ip-restriction

可以在service、router、consumer、全局四个维度添加黑白名单。ipv4、ipv6均支持。

### service：

添加白名单

curl -X POST http://127.0.0.1:8001/services/service/plugins     --data "name=ip-restriction"      --data "config.whitelist=54.13.21.1"     --data "consumer.id=45c08aeb-c08c-409c-a573-38bffdb9f204"

更新白名单

覆盖更新

 curl -X PUT http://127.0.0.1:8001/services/service/plugins/e0a66af3-ef02-4afc-918a-fb348eca442e     --data "name=ip-restriction"      --data "config.whitelist=54.13.21.1"  --data "config.whitelist=54.13.21.2"   --data "consumer.id=45c08aeb-c08c-409c-a573-38bffdb9f204"

获取白名单

curl -X GET http://127.0.0.1:8001/services/service/plugins/e0a66af3-ef02-4afc-918a-fb348eca442e     --data "name=ip-restriction"      --data "config.whitelist=54.13.21.1"  --data "config.whitelist=54.13.21.2"   --data "consumer.id=45c08aeb-c08c-409c-a573-38bffdb9f204"

删除白名单

curl -X DELETE http://127.0.0.1:8001/services/service/plugins/e0a66af3-ef02-4afc-918a-fb348eca442e     --data "name=ip-restriction"      --data "config.whitelist=54.13.21.1"  --data "config.whitelist=54.13.21.2"   --data "consumer.id=45c08aeb-c08c-409c-a573-38bffdb9f204"

禁用

curl -X PUT http://127.0.0.1:8001/services/service/plugins/5d4d52df-c623-4fef-ab35-308e9b16aba3     --data "name=ip-restriction"      --data "enabled=false"

### 获取真实ip

https://www.imooc.com/article/288036

-e “KONG_TRUSTED_IPS=0.0.0.0/0,::/0”
-e “KONG_REAL_IP_HEADER=X-Forwarded-For”   // 当前镜像是x-real-ip





## 负载均衡

创建upstream

upstream中添加target

创建service，host指向upstream name

创建与service关联的route

如何绑定VIP？？？ 监听其他端口？？？




