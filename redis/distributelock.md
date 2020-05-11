
参考： http://www.redis.cn/articles/20181020004.html

基本思想
	获取锁之前查询一下该锁为key的value是否存在，存在则说明被锁了，不存在则为该key设置一个value，即可锁定
	为避免获取锁后程序宕机导致无法释放，需给锁加过期时间
	伪码
	`
		func Lock(key string, timeout string)bool{
			if (exist(key)){
				return false
			}else{
				set(key, timeout)
			}
			return true
		}
	`
	`
		func myFunc(){
			lock := Lock(key)
			if lock{
				doSomething()
				Unlock(key)
			}
		}
	`

问题一： 锁不一定是排他锁，而set非原子操作导致多个进程获取到锁并set
`
SET my_key my_value NX PX milliseconds
NX表示只有当key不存在时才会设置
`
问题二： 释放其他客户端获取的锁
eg:
```
设置的超时时间时2秒，A执行了3s才结束，此时锁已释放后被B获取，A继续执行unlock导致释放了B的锁
一个很简单的方法是，我们设置key的时候，将value设置为一个随机值r，当释放锁，也就是删除key的时候，不是直接删除，而是先判断该key对应的value是否等于先前设置的随机值，只有当两者相等的时候才删除该key，由于每个客户端产生的随机值是不一样的，这样一来就不会误释放别的客户端申请的锁了。通过lua脚本保证操作的原子性
if redis.call("get",KEYS[1]) == ARGV[1] then
    return redis.call("del",KEYS[1])
else
    return 0
end
```