kong

调kong的api  主要涉及consumer、auth key / hmac 、acl 这几个插件



8000：监听HTTP请求并将请求转发到上游服务器。

8443：监听http请求但不转发。

8001：admin api

8444：监控

## service

服务实体，上游服务器的抽象。一个service可以有多个route，匹配到的route就会转发到service里。可以执行一个target(物理服务)或upstream



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



# admin api

## service

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

## router

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



## consumer object

服务的消费者或用户。依赖kong作为主要数据存储，也可以将用户自己管理的列表映射到该数据库consumer表。权限控制依赖这个表。

创建消费者

```
 curl -i -X POST   --url  http://localhost:8001/consumers/   --data 'username=zhou'   --data 'custom_id=000150' 
```



## plugin object

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











