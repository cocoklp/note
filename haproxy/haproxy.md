# 网站

   https://www.haproxy.org/ （官方网站）

   https://www.haproxy.org/download/1.8/src/haproxy-1.8.14.tar.gz （下载地址）

   http://cbonte.github.io/haproxy-dconv/1.8/configuration.html （文档Haproxy 1.8 文档）



# 负载均衡

## 四层负载均衡

根据ip和端口转发流量

![img](https://img-blog.csdn.net/20181025091707483)

## 七层负载均衡

可以根据用户请求的内容将请求转发到不同的后端服务器。允许在一个ip+port下运行多个应用服务器。比如某个uri发到后端1，其他uri发到后端2.

![img](https://img-blog.csdn.net/20181025091950153)



# haproxy部署

haproxy.org 下载安装包 （LTS长期支持稳定）

编译安装

```
yum -y install gcc readline-devel
```

安装lua

```
cd /usr/local/src
 wget http://www.lua.org/ftp/lua-5.3.5.tar.gz
 tar -xf lua-5.3.5.tar.gz 
cd lua-5.3.5
make linux test
 
 
[root@localhost lua-5.4.0]# src/lua -v
Lua 5.3.5  Copyright (C) 1994-2018 Lua.org, PUC-Rio
```



```
yum -y install gcc openssl-devel pcre-devel systemd-devel
```

解压安装

```
tar -xf haproxy-2.0.16.tar.gz -C /usr/local/src/
cd /usr/local/src/haproxy-2.0.16/
 make ARCH=x86_64 TARGET=linux-glibc USE_PCRE=1 USE_OPENSSL=1 USE_ZLIB=1 USE_SYSTEMD=1 USE_LUA=1 LUA_INC=/usr/local/src/lua-5.3.5/src/ LUA_LIB=/usr/local/src/lua-5.3.5/src/
make install PREFIX=/apps/haproxy
ln -s /apps/haproxy/sbin/haproxy /usr/sbin/
haproxy -v
HA-Proxy version 2.0.16 2020/07/17 - https://haproxy.org/
```

启动配置

```
cat /usr/lib/systemd/system/haproxy.service
[Unit]
Description= root
After=network.target
 
[Service]
Type=forking
PIDFile=/var/lib/haproxy/haproxy.pid
ExecStart=/usr/sbin/haproxy -f /etc/haproxy/haproxy.cfg
ExecReload=/bin/kill -s HUP $MAINPID
ExecStop=/bin/kill -s TERM $MAINPID
 
[Install]
WantedBy=multi-user.target


```

systemctl daemon-reload
mkdir /etc/haproxy
mkdir /var/lib/haproxy
useradd -r -s /sbin/nologin -d /var/lib/haproxy/ haproxy

代理配置

```
global
        maxconn 100000
        chroot /apps/haproxy
        stats socket /var/lib/haproxy/haproxy.sock mode 600 level admin
	user haproxy
	group haproxy 
	daemon 
	nbproc 1
        spread-checks 3
	pidfile /var/lib/haproxy/haproxy.pid
        log 127.0.0.1 local2 info
defaults
        option http-keep-alive
        option forwardfor
        maxconn 100000
        mode http
        timeout connect 300000ms
        timeout client 300000ms
        timeout server 300000ms
 
listen proxy_status 
        mode http
        bind 0.0.0.0:9999
        stats enable
        log global
        stats uri     /haproxy-status
        stats auth   haadmin:123456
listen http_80
        log global
	bind 0.0.0.0:9998 
        server web_1 114.67.81.96:9191 check inter 3000 fall 2 rise 5 
        server web_2 114.67.81.96:9192 check inter 3000 fall 2 rise 5
```



重启

```
systemctl daemon-reload
systemctl start haproxy
```



# 配置

- global

  全局参数，通常和操作系统相关.

  ```
  global   # 全局参数的设置
      log 127.0.0.1 local0 info
      # log语法：log <address_1>[max_level_1] # 全局的日志配置，使用log关键字，指定使用127.0.0.1上的syslog服务中的local0日志设备，记录日志等级为info的日志
      user haproxy
      group haproxy
      # 设置运行haproxy的用户和组，也可使用uid，gid关键字替代之
      daemon
      # 以守护进程的方式运行
      nbproc 16
      # 设置haproxy启动时的进程数，根据官方文档的解释，我将其理解为：该值的设置应该和服务器的CPU核心数一致，即常见的2颗8核心CPU的服务器，即共有16核心，则可以将其值设置为：<=16 ，创建多个进程数，可以减少每个进程的任务队列，但是过多的进程数也可能会导致进程的崩溃。这里我设置为16
      maxconn 4096
      # 定义每个haproxy进程的最大连接数 ，由于每个连接包括一个客户端和一个服务器端，所以单个进程的TCP会话最大数目将是该值的两倍。
      #ulimit -n 65536
      # 设置最大打开的文件描述符数，在1.4的官方文档中提示，该值会自动计算，所以不建议进行设置
      pidfile /var/run/haproxy.pid
      # 定义haproxy的pid 
  ```

  // log

  ```
  1)在/etc/rsyslog.conf中增加如下配置后重启rsyslog服务
  local2.info           /var/log/test.log
  执行命令
  [root@localhost ~]# logger -p local2.info "hello world"
  查看 /var/log/test.log
  [root@localhost ~]# tail /var/log/test.log             
  Nov 18 22:36:30 localhost root: hello world
  ```

  

- defaults

  ```
  
  defaults # 默认部分的定义
      mode http
      # mode语法：mode {http|tcp|health} 。http是七层模式，tcp是四层模式，health是健康检测，返回OK
      log 127.0.0.1 local3 err
      # 使用127.0.0.1上的syslog服务的local3设备记录错误信息
      retries 3
      # 定义连接后端服务器的失败重连次数，连接失败次数超过此值后将会将对应后端服务器标记为不可用
      option httplog
      # 启用日志记录HTTP请求，默认haproxy日志记录是不记录HTTP请求的，只记录“时间[Jan 5 13:23:46] 日志服务器[127.0.0.1] 实例名已经pid[haproxy[25218]] 信息[Proxy http_80_in stopped.]”，日志格式很简单。
      option redispatch
      # 当使用了cookie时，haproxy将会将其请求的后端服务器的serverID插入到cookie中，以保证会话的SESSION持久性；而此时，如果后端的服务器宕掉了，但是客户端的cookie是不会刷新的，如果设置此参数，将会将客户的请求强制定向到另外一个后端server上，以保证服务的正常。
      option abortonclose
      # 当服务器负载很高的时候，自动结束掉当前队列处理比较久的链接
      option dontlognull
      # 启用该项，日志中将不会记录空连接。所谓空连接就是在上游的负载均衡器或者监控系统为了探测该服务是否存活可用时，需要定期的连接或者获取某一固定的组件或页面，或者探测扫描端口是否在监听或开放等动作被称为空连接；官方文档中标注，如果该服务上游没有其他的负载均衡器的话，建议不要使用该参数，因为互联网上的恶意扫描或其他动作就不会被记录下来
      option httpclose
      # 这个参数我是这样理解的：使用该参数，每处理完一个request时，haproxy都会去检查http头中的Connection的值，如果该值不是close，haproxy将会将其删除，如果该值为空将会添加为：Connection: close。使每个客户端和服务器端在完成一次传输后都会主动关闭TCP连接。与该参数类似的另外一个参数是“option forceclose”，该参数的作用是强制关闭对外的服务通道，因为有的服务器端收到Connection: close时，也不会自动关闭TCP连接，如果客户端也不关闭，连接就会一直处于打开，直到超时。
      contimeout 5000
      # 设置成功连接到一台服务器的最长等待时间，默认单位是毫秒，新版本的haproxy使用timeout connect替代，该参数向后兼容
      clitimeout 3000
      # 设置连接客户端发送数据时的成功连接最长等待时间，默认单位是毫秒，新版本haproxy使用timeout client替代。该参数向后兼容
      srvtimeout 3000
      # 设置服务器端回应客户度数据发送的最长等待时间，默认单位是毫秒，新版本haproxy使用timeout server替代。该参数向后兼容
  ```

  

- frontend

  ```
  frontend http_80_in # 定义一个名为http_80_in的前端部分
      bind 0.0.0.0:80
      # http_80_in定义前端部分监听的套接字
      mode http
      # 定义为HTTP模式
      log global
      # 继承global中log的定义
      option forwardfor
      # 启用X-Forwarded-For，在requests头部插入客户端IP发送给后端的server，使后端server获取到客户端的真实IP
      acl static_down nbsrv(static_server) lt 1
      # 定义一个名叫static_down的acl，当backend static_sever中存活机器数小于1时会被匹配到
      acl php_web url_reg /*.php$
      #acl php_web path_end .php
      # 定义一个名叫php_web的acl，当请求的url末尾是以.php结尾的，将会被匹配到，上面两种写法任选其一
      acl static_web url_reg /*.(css|jpg|png|jpeg|js|gif)$
      #acl static_web path_end .gif .png .jpg .css .js .jpeg
      # 定义一个名叫static_web的acl，当请求的url末尾是以.css、.jpg、.png、.jpeg、.js、.gif结尾的，将会被匹配到，上面两种写法任选其一
      use_backend php_server if static_down
      # 如果满足策略static_down时，就将请求交予backend php_server
      use_backend php_server if php_web
      # 如果满足策略php_web时，就将请求交予backend php_server
      use_backend static_server if static_web
      # 如果满足策略static_web时，就将请求交予backend static_server
  ```

  

- backend

  ```
  backend php_server #定义一个名为php_server的后端部分
      mode http
      # 设置为http模式
      balance source
      # 设置haproxy的调度算法为源地址hash
      cookie SERVERID
      # 允许向cookie插入SERVERID，每台服务器的SERVERID可在下面使用cookie关键字定义
      option httpchk GET /test/index.php
      # 开启对后端服务器的健康检测，通过GET /test/index.php来判断后端服务器的健康情况
      server php_server_1 10.12.25.68:80 cookie 1 check inter 2000 rise 3 fall 3 weight 2
      server php_server_2 10.12.25.72:80 cookie 2 check inter 2000 rise 3 fall 3 weight 1
      server php_server_bak 10.12.25.79:80 cookie 3 check inter 1500 rise 3 fall 3 backup
      # server语法：server [:port] [param*] # 使用server关键字来设置后端服务器；为后端服务器所设置的内部名称[php_server_1]，该名称将会呈现在日志或警报中、后端服务器的IP地址，支持端口映射[10.12.25.68:80]、指定该服务器的SERVERID为1[cookie 1]、接受健康监测[check]、监测的间隔时长，单位毫秒[inter 2000]、监测正常多少次后被认为后端服务器是可用的[rise 3]、监测失败多少次后被认为后端服务器是不可用的[fall 3]、分发的权重[weight 2]、最后为备份用的后端服务器，当正常的服务器全部都宕机后，才会启用备份服务器[backup]
      backend static_server
      mode http
      option httpchk GET /test/index.html
      server static_server_1 10.12.25.83:80 cookie 3 check inter 2000 rise 3 fall 3
  ```

  

- status

  ```
  listen status # 定义一个名为status的部分
      bind 0.0.0.0:1080
      # 定义监听的套接字
      mode http
      # 定义为HTTP模式
      log global
      # 继承global中log的定义
      stats refresh 30s
      # stats是haproxy的一个统计页面的套接字，该参数设置统计页面的刷新间隔为30s
      stats uri /admin?stats
      # 设置统计页面的uri为/admin?stats
      stats realm Private lands
      # 设置统计页面认证时的提示内容
      stats auth admin:password
      # 设置统计页面认证的用户和密码，如果要设置多个，另起一行写入即可
      stats hide-version
      # 隐藏统计页面上的haproxy版本信息
  ```

  

- listen

# haproxy重新加载配置

## 重启

systemctl restart haproxy

1. 连接多时容易卡死

## hot-reload

haproxy -f /etc/haproxy/haproxy.cfg -st

haproxy -f /etc/haproxy/haproxy.cfg -sf

用-sf执行的时候，发现老进程有长连接挂着而长时间无法结束。
而-st了之后，则立刻关闭了老进程。
-st（terminates）是强行关闭进程，而-sf（finishes）则是优雅结束。



检查配置文件

```
haproxy -f configuration.conf -c
```



# 健康检查

default backend listen 中使用以下配置，frontend中不可用。

```
option httpchk
option httpchk <uri>
option httpchk <method> <uri>
option httpchk <method> <uri> <version>
```

1. global里有external-check，可以使用外部agent进行健康检查。

2. 通过监听端口，只探测端口，并不保证服务可用

   ```
   option httpchk
   sever web1 ip:port check
   ```

3. 通过uri进行检查

   ```
    option httpchk GET /index.html
   ```

4. request获取的头部信息进行匹配进行健康检测

   ```
   option httpchk HEAD /index.jsp HTTP/1.1\r\nHost:\ www.xxx.com
   ```





# dataplane


https://www.e-learn.cn/content/qita/2633565

https://github.com/haproxytech/dataplaneapi

http://114.67.72.40:9999/haproxy-status

https://www.haproxy.com/documentation/hapee/1-9r1/reference/dataplaneapi/

## 安装

https://github.com/haproxytech/dataplaneapi

下载源代码，然后make build，进入build后，

```
./dataplaneapi --port 5555 -b /usr/sbin/haproxy -c /etc/haproxy/haproxy.cfg  -d 5 -r "service haproxy reload" -s "service haproxy restart" -u dataplaneapi -t /tmp/haproxy

// -u dataplaneapi 要与 haproxy.conf 里的 userlist dataplaneapi一致
```

## version VS transaction

POST  PUT DELETE 等写方法，需要再url中添加version以防冲突，如果写的version与配置文件中不符，会返回version mismatch

transaction：

​	事务，多条命令先存起来，然后一次执行。

​	不同transaction之间是隔离的

## 增加代理过程

添加backend，给backend添加server，添加frontend并bind ip port，给frontend增加use_backend配置。

1. initialize a transaction

   ```
   curl -X POST --user admin:123456 -H "Content-Type: application/json" http://114.67.72.40:5555/v2/services/haproxy/transactions?version=4
   resp：
   {"_version":4,"id":"edbf7220-5725-4f34-a1b4-cff8af344331","status":"in_progress"}
   ```

   后续请求携带id

2. add backend

   ```
    curl -X POST -u admin:123456 -H "Content-Type: application/json" "http://114.67.72.40:5555/v2/services/haproxy/configuration/backends?transaction_id=2ad5391a-5f4f-4707-8664-ff3559290f66" -d '{"adv_check":"httpchk","balance":{"algorithm":"roundrobin"},"forwardfor":{"enabled":"enabled"},"httpchk_params":{"method":"GET","uri":"/check","version":"HTTP/1.1"},"mode":"http","name":"test_backendtr"}'
   ```

3. add server to backend

   ```
   curl -X POST --user admin:123456 \
   -H "Content-Type: application/json" \
   -d '{"name": "server1", "address": "114.67.81.96", "port": 9191, "check": "enabled", "maxconn": 30, "weight": 100}' \
   "http://114.67.72.40:5555/v2/services/haproxy/configuration/servers?backend=test_backendtr&transaction_id=2ad5391a-5f4f-4707-8664-ff3559290f66"
   ```

4. add frontend

   ```
    curl -X POST -u admin:123456 -H "Content-Type: application/json" "http://114.67.72.40:5555/v2/services/haproxy/configuration/frontends?transaction_id=2ad5391a-5f4f-4707-8664-ff3559290f66" -d '{"default_backend":"test_backend","http_connection_mode":"http-keep-alive","maxconn":2000,"mode":"http","name":"test_frontendtr"}'
   ```

   

5. bind ip port to frontend

   ```
   
   curl -X POST --user admin:123456 \
   -H "Content-Type: application/json" \
   -d '{"name": "http", "address": "*", "port": 80}' \
   "http://114.67.72.40:5555/v2/services/haproxy/configuration/binds?frontend=test_frontendtr&transaction_id=2ad5391a-5f4f-4707-8664-ff3559290f66"
   ```

6. add user_backend to frontend

   backend switching rule

   ```
   curl -X POST --user admin:123456 -H "Content-Type: application/json" -d '{"id": 0, "name": "test_backendproxy","index":1}' "http://127.0.0.1:5555/v2/services/haproxy/configuration/backend_switching_rules?frontend=test_frontend&version=35"
   ```

   index: 必须是从零开始顺序递增的，否则会报out of range，如果有0 1 2 3四个，那么删除2后，变成0 1 2 三个，可以GET到index

   ```
    curl -X GET --user admin:123456 -H "Content-Type: application/json"  "http://127.0.0.1:8080/v2/services/haproxy/configuration/backend_switching_rules?frontend=test_frontend&version=32"
   {"_version":36,"data":[{"index":0,"name":"test_backendtr"},{"index":1,"name":"test_backendproxy"}]}
   ```

    ![image-20210112174956718](C:\Users\kouliping\AppData\Roaming\Typora\typora-user-images\image-20210112174956718.png)

7. commit transaction

   ```
   curl -X PUT --user admin:123456 -H "Content-Type: application/json" "http://127.0.0.1:5555/v2/services/haproxy/transactions/2ad5391a-5f4f-4707-8664-ff3559290f66"
   ```

   



# base auth

```
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	user := "admin"
	pass := "123456"
	encoded := base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
	fmt.Println(encoded)
}
输出 YWRtaW46MTIzNDU2
```



curl -u admin:123456 等价于 curl -H "Authorization:Basic YWRtaW46MTIzNDU2"  // 注意Basic之后的空格。

```
package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const pulltimeout = 5000
const pulltimes = 2

func main() {
	user := "admin"
	pass := "123456"
	encoded := base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
	header := map[string]string{
		"Authorization": "Basic " + encoded,
	}
	url := os.Args[1]
	//rbody, e, _, hs := tryGetRespBodyWithHeaders("", url, pulltimeout, pulltimes, nil)
	code, resp, hs, err := DoRequest("GET", url, "", header, nil, pulltimeout)
	fmt.Println(time.Now())
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(strings.HasSuffix(err.Error(), "connection timed out"))
	}

	headers := http.Header(hs)
	fmt.Println(string(resp))
	fmt.Println(code, err, hs, headers)
}

func tryGetRespBodyWithHeaders(host string, queryAddr string, timeout int64, tryTimes int,
	inHeaders map[string]string) ([]byte, error, int, map[string][]string) {
	var try int
	var rbody []byte
	var e error
	var headers map[string][]string
	for try < tryTimes {
		try++
		statusCode, body, hs, e := DoRequest("GET", queryAddr, host, inHeaders, nil, timeout)
		if e == nil && statusCode >= http.StatusOK && statusCode < http.StatusBadRequest {
			rbody = body
			headers = hs
			break
		}
	}
	fmt.Printf("pulled webcache page[%s], length of body(bytes)[%d], error %v", queryAddr, len(rbody), e)
	return rbody, e, try, headers
}

// reqType is one of HTTP request strings (GET, POST, PUT, DELETE, etc.)
func DoRequest(reqType string, url string, host string, headers map[string]string, data []byte, timeoutMs int64) (int, []byte, map[string][]string, error) {
	var reader io.Reader
	if data != nil && len(data) > 0 {
		reader = bytes.NewReader(data)
	}

	req, err := http.NewRequest(reqType, url, reader)
	if err != nil {
		return 0, nil, nil, err
	}

	if host != "" {
		req.Host = host
	}
	// I strongly advise setting user agent as some servers ignore request without it
	req.Header.Set("User-Agent", "Mozilla/5.0")
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	var (
		statusCode int
		body       []byte
		timeout    time.Duration
		ctx        context.Context
		cancel     context.CancelFunc
		header     map[string][]string
	)
	timeout = time.Duration(time.Duration(timeoutMs) * time.Millisecond)
	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req = req.WithContext(ctx)
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}

		defer resp.Body.Close()
		body, _ = ioutil.ReadAll(resp.Body)
		statusCode = resp.StatusCode
		header = resp.Header

		return nil
	})

	return statusCode, body, header, err
}

// httpDo issues the HTTP request and calls f with the response. If ctx.Done is
// closed while the request or f is running, httpDo cancels the request, waits
// for f to exit, and returns ctx.Err. Otherwise, httpDo returns f's error.
func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass thehe response to f.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	/*
		resp, err := client.Do(req)
						fmt.Println(resp, err)
												return err
	*/

	c := make(chan error, 1)
	go func() { c <- f(client.Do(req)) }()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c // Wait for f to return.
		return ctx.Err()
	case err := <-c:
		return err
	}
}
```



# 添加vip

```
for i in {1..100}
do
  ifconfig eth0:$i netmask 255.255.255.0 up 166.111.69.$i
done
```

如果不在同网段，需要在发起ping的机器上配置路由

```
两台机器：172.16.91.101 172.16.91.107
在91.101上增加虚拟ip，92网段的
ifconfig eth0:1 172.16.92.2 netmask 255.255.255.0 up
由于不再同一个网段需要添加路由(与vip交互的机器上）：
增加路由:route add -net 172.16.92.0 netmask 255.255.255.0 gw 172.16.91.101
注意：
添加静态路由规则的时候，需要保证gateway(gw)的IP和eth0（本机IP）在同一个网段内。否则route add报错SIOCADDRT: Network is unreachable
```



```
本机访问自己的虚拟ip，无论是否同一网段，均不需要设置路由。
```



https://blog.csdn.net/u012599988/article/details/82683440

ip addr del 10.1.1.1/24 dev ens33

ip addr add 10.1.1.1/24 dev ens33

同一网段删除会会有影响，

/sbin/sysctl net.ipv4.conf.eth0.promote_secondaries=1

echo "net.ipv4.conf.eth0.promote_secondaries=1" >>/etc/sysctl.conf



重启后丢失，需要把执行脚本写到/etc/rcc.local中，并检查/etc/rc.local和/etc/rc.d/rc.local的执行权限。





添加backend

```
curl -X POST -u admin:123456 -H "Content-Type: application/json" "http://114.67.72.40:5555/v2/services/haproxy/configuration/backends" -d `{"adv_check":"httpchk","balance":{"algorithm":"roundrobin"},"forwardfor":{"enabled":"enabled"},"httpchk_params":{"method":"GET","uri":"/check","version":"HTTP/1.1"},"mode":"http","name":"test_backend"}`
```

添加server到backend



添加frontend

```
curl -X POST -u admin:123456 -H "Content-Type: application/json" "http://114.67.72.40:5555/v2/services/haproxy/configuration/frontends" -d `{"default_backend":"test_backend","http_connection_mode":"http-keep-alive","maxconn":2000,"mode":"http","name":"test_frontend"}`
```

添加bind到frontend





frontend 怎么添加use_backend??？
每调一次post接口，version自增1
Add frontend,bind

add backend,add server

# 允许绑定非本机IP

echo 1 > /proc/sys/net/ipv4/ip_nonlocal_bind





# keepalived + haproxy



![img](https://img-blog.csdnimg.cn/20200616114902659.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTk0NzI2Nw==,size_16,color_FFFFFF,t_70)



# session

https://blog.csdn.net/weixin_45537987/article/details/106759391

## 源地址hash

根据用户IP指定rs服务器

当后端一台服务器挂了以后会造成部分session丢失

```
backend ***
```
	balance source
	```
```



## cookie识别

```
backend ***
	```
	cookie SERVERID insert indirect nocache
	server node1 ***
	server node2 ***
	```
```

如果客户端请求不带cookie，则请求被转发到任一可用服务器，并且haproxy把该服务器信息插入到响应中，如SERVERID=A，后续客户端再次请求时，携带该cookie，haproxy会将该cookie移除后转发给指定服务其。如果指定服务器不可用，则转发给其他web服务器并重新设置cookie字段。





## 基于session

haproxy将后端产生的session和后端服务器标识存在haproxy的一张表里，客户端请求时先查该表。

```
backend ***
	```
	appsession JSESSIONID len 64 timeout 5h request-learn
	server node1 ***
	server node2 ***
	```
```



1. 如果使用http1.1 keep-alive，haproxy只在第一个响应中插入cookie，且达到服务器请求中的cookie不会被删除，所以要求后端服务器对来历不明的cookie不做敏感处理，否则需要关闭keep-alive

```
   option httpclose
   ```

   

2. 如果禁用浏览器cookie功能，可以使用source负载均衡算法。

3. 使用source算法时，因为haproxy对cookie的处理具有较高优先级，所以即使客户端是动态ip，只要可以接收cookie，对转发不会有影响。

   






   ```