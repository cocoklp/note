curl记录响应时间：
	新建文件 format.txt{
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

curl -w
	诊断问题
	curl -X POST http://10.226.193.42:8900/v1/regions/cn-north-1/wafInstanceIds/testtest810-waf/domain:listMainCfg -H "userPin:testtest810" -d '{"req":{"pageSize":100}}'  -w  "time_connect: %{time_connect}\ntime_starttransfer: %{time_starttransfer}\ntime_nslookup:%{time_namelookup}\ntime_total: %{time_total}\n" -o /dev/null

/dev/null
	空设备，丢弃一切写入其中的数据，但报告写入操作成功，读取会得到一个EOF
	黑洞，写入的内容会永远丢失，用于脚本和命令行
	cat $filename 2> /dev/null  :: 			
	cat $filename 1> /dev/null ： 文件不存在时会打印错误. 标准输出

修改机器时间
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
查看本地系统时间
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
系统时间覆盖硬件时间
```
# clock -w
# hwclock -w 
```
设置硬件时间
```
//设置硬件时间为18年6月11日10点10分0秒
# hwclock --set --date '2018-06-11 10:10:00'
```