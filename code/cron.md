定时任务

秒(0-59 * / , -)

 分(0-59 * / , -)

 时(0-23 * / , -) 

日(1-31 * / , -) 

月(1-21 or jan-dec) 

星期(0-6 or sun-sat)

字母不区分大小写

**匹配该字段所有值，比如第五个字段使用*表示每个月

/ 增长间隔，min字段值为 3-59/15表示每小时第三分钟开始执行一次，后续每15分钟执行一次。

, 枚举，星期字段 mon,wed,fri表示周一 周三 周五。

-表示一个范围，比如小时字段为9-17表示每个小时，包括9 和 17

?表示不指定值，可用于替代*

```
每隔5秒执行一次：*/5 * * * * ?
每隔1分钟执行一次：0 */1 * * * ?
每天23点执行一次：0 0 23 * * ?
每天凌晨1点执行一次：0 0 1 * * ?
每月1号凌晨1点执行一次：0 0 1 1 * ?
在26分、29分、33分执行一次：0 26,29,33 * * * ?
每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?
```



新版本的cron不再默认支持秒，需要cron.WithSeconds()方法定义

```
package main

import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {
	fmt.Println("cron running:")
	i := 0
	c := cron.New(cron.WithSeconds())
	spec := "*/1 * * * * *"
	c.AddFunc(spec, func() {
		i++
		fmt.Println("cron running:", i)
	})
	go c.Start()
	fmt.Println("cron start")

	select {}
	fmt.Println("cron end")
}
```



