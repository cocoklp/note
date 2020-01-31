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