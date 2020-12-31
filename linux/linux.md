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