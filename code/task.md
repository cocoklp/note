# 任务队列

处理lb的增删改



## golang channel

```
type taskInfo struct{
	Operation string
	Data interface{}
}

taskChan := make(chan taskInfo,100)
producer:
	controller流程中，taskChan <- task
worker & consumer
	for task := range taskChan {
    	// analyze operation and data then call dataplaneapi
    }
```

优：实现简单

缺：重启丢失数据。任务不能修改(比如增加了一个lb规则，还未真正下发时再次修改改规则，则会首先顺序下发两次)

## redis

使用redis的lpush rpop实现

```
producer:
	controller流程中，redisPool.Do("lpush","tasklist",appName)
	redis中appName为key计数加一
worker & consumer:	
	if tasklist非空
	task = redisPool.Do("rpoplpush",‘tasklist’, ‘taskdoing’)
	redis中appName为key计数减一，如果减一后非零，说明队列中有该appName的新任务，该任务跳过。如果为零，根据appName查数据库，得到最终配置，下发
	redisPool.Do("lrem",1,task)
	检查taskdoing，如果有内容说明任务冲tasklist取出后没有执行完毕就重启了，此时把taskdoing的内容写回tasklist，appName为key技术加以，等待后续执行
```

优点：重启不丢失任务，配置多次修改，理想情况只下发最后一次

缺点：引入redis，实现相对复杂。配置多次修改仍有小概率多次下发配置。

## pgsql

```
db中存储全量配置，并增加update_time列
producer：
	insert/update db,插入/更新配置，updatetime记录最后一次修改时间
consumer：
	定时器，读取status为未执行且update_time最小的数据
	do something
	update status，此时需要判断update_time是否为之前取出的，如果不一致则不更新status
```

优点：重启不丢失任务，配置多次修改，理想情况只下发最后一次，实现相对简单

缺点：pgsql 可能会有性能瓶颈。如果读出数据后，执行前再次修改配置，仍会导致下发多次配置。