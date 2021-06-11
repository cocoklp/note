# curl记录响应时间：
​	新建文件

``` format.txt{
	    time_namelookup:  %{time_namelookup}\n
	       time_connect:  %{time_connect}\n
	    time_appconnect:  %{time_appconnect}\n
	   time_pretransfer:  %{time_pretransfer}\n
	      time_redirect:  %{time_redirect}\n
	 time_starttransfer:  %{time_starttransfer}\n
	                    ----------\n
	         time_total:  %{time_total}\n	
	}
	curl -i -w "@curl-format.txt" -o /dev/null -d "id=xxx&pwd=xxx" -X POST http://localhost:8080/api/login
```

curl l -s -w ‘%{time_total}’  url

# curl -w

​	诊断问题

```
curl -X POST http://10.226.193.42:8900/v1/regions/cn-north-1/wafInstanceIds/testtest810-waf/domain:listMainCfg -H "userPin:testtest810" -d '{"req":{"pageSize":100}}'  -w  "time_connect: %{time_connect}\ntime_starttransfer: %{time_starttransfer}\ntime_nslookup:%{time_namelookup}\ntime_total: %{time_total}\n" -o /dev/null
```

#  文件重定向

/dev/null
	空设备，丢弃一切写入其中的数据，但报告写入操作成功，读取会得到一个EOF
	黑洞，写入的内容会永远丢失，用于脚本和命令行
	cat $filename 2> /dev/null  :: 			
	cat $filename 1> /dev/null ： 文件不存在时会打印错误. 标准输出

0：stdin标准输入

1：stdout标准输出

2：stderr标准错误

## nohup command>/dev/null 2>&1 &

command >/dev/null

 相当于 command 1 > /dev/null， command产生了标准输出，冲顶先到/dev/null中。

command 1>a 2>&1 VS command 1>a 2>a

   2>&1 标准错误等效于标准输出

  command 1>a 2>&1 标准输出和错误都重定向到文件a中，且只打开一次文件。

  command 1>1 2>1 标准输出和错误都重定向到文件a中，打开两次文件，stdout会被stderr覆盖。

## exec

exec永久重定向。



# 时间

## 修改机器时间

```
//将时间改为10点10分
# date -s 10:10
//将日期改为2018年6月11日
# date -s 2018-06-11
//将日期改为2018年6月11日10点10分
# date -s '2018-06-11 10:10:00'
//将日期改为2018年6月11日10点10分，并覆盖硬件时间
# date -s '2018-06-11 10:10:00' && hwclock -w
```
## 查看本地系统时间

```
date
timedatectl status
------------------------
      Local time: 一 2018-07-23 10:39:31 CST
  Universal time: 一 2018-07-23 02:39:31 UTC
        RTC time: 一 2018-07-23 02:39:30			// 硬件时间
        Timezone: Asia/Shanghai (CST, +0800)
     NTP enabled: no
NTP synchronized: yes
 RTC in local TZ: no
      DST active: n/a

```
## 系统时间覆盖硬件时间

```
# clock -w
# hwclock -w 
```
## 设置硬件时间

```
//设置硬件时间为18年6月11日10点10分0秒
# hwclock --set --date '2018-06-11 10:10:00'
```
重复：
 for ((i=1;i<=100;i++)); do go run offset.go; done

 curl -g -6 -H"Host:waf.ipv6.com" http://[2402:db40:1570:11c3:0000:0000:0000:0025]:80

 curl支持SNI
curl -k -v 'https://visx.com:8443/' --resolve "visx.com:8443:127.0.0.1"

Linux下sz下载文件超过4G
	scp
	ftp
	nc
	icmp
	文件切割：
		文件切割
			cat xxx.tar |split  -b 2G - xxx.tar.
			会按2G文件进行切割，切割后的文件 xxx.tar.aa xxx.tar.ab xxx.tar.ac
		下载到本地
			sz xxx.tar.a* 
		合并
			cat xxx.tar.a* > xxx.tar
			哈希验证正确性

# 查看centos版本

cat /etc/redhat-releaseuname -a



# 自定义systemctl服务

https://www.cnblogs.com/fire909090/p/11303500.html

压缩解压缩：
https://blog.csdn.net/weixin_44901564/article/details/99682926
https://blog.csdn.net/freeking101/article/details/51480295
解压：
1.	tar –xvf file.tar         // 解压 tar包  
2.	tar -zxvf file.tar.gz     // 解压tar.gz  
3.	tar -jxvf file.tar.bz2    // 解压 tar.bz2  
4.	tar –Zxvf file.tar.Z      // 解压tar.Z 
压缩：
1.	tar –cvf jpg.tar *.jpg     // 将目录里所有jpg文件打包成tar.jpg  
2.	tar –czf jpg.tar.gz *.jpg  // 将目录里所有jpg文件打包成jpg.tar后，并且将其用gzip压缩，生成一个gzip压缩过的包，命名为jpg.tar.gz  
3.	tar –cjf jpg.tar.bz2 *.jpg // 将目录里所有jpg文件打包成jpg.tar后，并且将其用bzip2压缩，生成一个bzip2压缩过的包，命名为jpg.tar.bz2  
4.	tar –cZf jpg.tar.Z *.jpg   // 将目录里所有jpg文件打包成jpg.tar后，并且将其用compress压缩，生成一个umcompress压缩过的包，命名为jpg.tar.Z

tgz文件：
tar zxvf nginx-11.8.tgz 解压
tar czvf nginx-11.8.tgz nginx-11.8 压缩
tar参数
c 新建打包文件，cv一起使用查看过程中打包文件名
x 解压文件，-C 指定解压到的文件目录
f 指定要处理的文件
j bzip2方式压缩或解压，以.tar.br2为后缀
z gzip方式，.tar.gz后缀
v 压缩或解压过程，显示
t 查看打包文件中的内容
u 更新打包文件中的内容
p 保留绝对路径
P 保留数据原来的权限和属性



1、把/home目录下面的mydata目录压缩为mydata.zip
zip -r mydata.zip mydata #压缩mydata目录
2、把/home目录下面的mydata.zip解压到mydatabak目录里面
unzip mydata.zip -d mydatabak
3、把/home目录下面的abc文件夹和123.txt压缩成为abc123.zip
zip -r abc123.zip abc 123.txt
4、把/home目录下面的wwwroot.zip直接解压到/home目录里面
unzip wwwroot.zip
5、把/home目录下面的abc12.zip、abc23.zip、abc34.zip同时解压到/home目录里面
unzip abc\*.zip
6、查看把/home目录下面的wwwroot.zip里面的内容
unzip -v wwwroot.zip
7、验证/home目录下面的wwwroot.zip是否完整
unzip -t wwwroot.zip
8、把/home目录下面wwwroot.zip里面的所有文件解压到第一级目录
unzip -j wwwroot.zip
https://www.cnblogs.com/wangkongming/p/4305962.html


zip -q -r html.zip /home/html


[kouliping@A01-R04-I17-130-06CEKFK bin]$ sed -i 's/^M//g' ./supervise.logs-manager
^M输入的方法是ctrl+v然后m
[kouliping@A01-R04-I17-130-06CEKFK bin]$ cat -A supervise.logs-manager
查找文件
find / -name filename
locate filename


查看进程总数
ps -ef |grep ping |wc -l



protoc
安装下载
下载源码：
https://github.com/google/protobuf/releases/tag/v3.6.1
解压后得到 bin、include、readme
go get github.com/golang/protobuf/proto


得到bin和src，拷贝到linux中安装protoc
修改gobin为src所在路径
生成配置文件：
./protoc --go_out=. *.proto



查看磁盘空间
cannot create temp file for here-document: No space left on device
df –sh /*  // 逐级查看占用



环境变量
打开cmd，set PATH=%PATH%;**


grep -v 不包含
grep常用用法
[root@www ~]# grep [-acinv] [--color=auto] '搜寻字符串' filename
选项与参数：
-a ：将 binary 文件以 text 文件的方式搜寻数据
-c ：计算找到 '搜寻字符串' 的次数
-i ：忽略大小写的不同，所以大小写视为相同
-n ：顺便输出行号
-v ：反向选择，亦即显示出没有 '搜寻字符串' 内容的那一行！
--color=auto ：可以将找到的关键词部分加上颜色的显示喔！



## syslog
local0~local7保留给本机用户使用
在/etc/rsyslog.conf中增加配置
local2.*  /var/log/test.log
然后写 logger -p local2.info "hello world"（-p 指定优先级）,查看/var/log/test.log会看到 hello world的日志



kill杀死多个进程
 ps -ef|grep LOCAL=NO|grep -v grep|cut -c 9-15|xargs kill -9

