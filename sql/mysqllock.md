# mysql 锁
`
	mysql的锁可以分为服务层实现的锁，例如Lock Tables、全局读锁、命名锁、字符锁，或者存储引擎的锁，例如行级锁。
`

服务级别的锁：
===

表锁（MyISAM）
=====
`
lock table table_name read
	不阻塞其他线程读，阻塞其他线程写
	不允许对标进行更新、插入、删除，不允许访问未被锁住的表   ？？ 
lock table table_name write
	阻塞其他线程的读和写
	允许访问其他表
unlock tables;
`

全局读锁
=====
`
flush tables with read lock;
阻塞一切操作，用于数据库备份
`

命名锁
======
`
	命名锁是一种表锁，服务器创建或者删除表的时候会创建一个命名锁。如果一个线程LOCK TABLES，另一个线程对被锁定的表进行重命名，查询会被挂起，通过show open tables可以看到两个名字(新名字和旧名字都被锁住了)。
`

字符锁
=====
`
	字符锁是一种自定义锁，通过SELECT GET_LOCK("xxx",60)来加锁 ，通过release_lock()解锁。假设A线程执行get_lock("xxx"，60)后执行sql语句返回结果为1表示拿到锁，B线程同样通过get_lock("xxx"，60)获取相同的字符锁，则B线程会处理阻塞等待的状况，如果60秒内A线程没有将锁释放，B线程获取锁超时就会返回0，表示未拿到锁。使用get_lock()方法获取锁，如果线程A调用了两次get_lock(),释放锁的时候也需要使用两次release_lock()来进行解锁。
`

InnoDB锁
=====
`
单条索引记录上加锁，record lock锁住的永远是索引，而非记录本身，即使该表上没有任何索引，那么innodb会在后台创建一个隐藏的聚集主键索引，那么锁住的就是这个隐藏的聚集主键索引。所以说当一条sql没有走任何索引时，那么将会在每一条聚集索引后面加X锁，这个类似于表锁，但原理上和表锁应该是完全不同的。
`

锁类别
=======
`
	共享锁(S): 允许事务读一行，阻塞其他事务对该数据进行修改
	排他锁(X): 允许事务去读取更新数据，阻塞其他事务对数据进行查询或修改
	意向共享锁(IS): 事务给一条数据加S锁时，先给整个表加IS锁，成功后才能加S锁
	意向排他锁(IX): 事务给一条数据加X锁时，先给整个表加IX锁，成功后才能加X锁
	/*行锁虽然很赞，但是还有一个问题，如果一个事务对一张表的某条数据进行加锁，这个时候如果有另外一个线程想要用LOCK TABLES进行锁表，这时候数据库要怎么知道哪张表的哪条数据被加了锁，一张张表一条条数据去遍历是不可行的。InnoDB考虑到这种情况，设计出另外一组锁，意向共享锁（IS）和意向排他锁(IX)*/
	意向锁之间兼容，但是与SX冲突
`
锁等待信息
=======
`
	Information_schema.processList
	Information_schema.innodb_lock_waits
	Information_schema.innodb_trx
	Information_schema.innodb_locks
`
索引和锁
=======
`
	添加行锁时，通过索引添加，如果查询没有用到索引，会使用表锁
	即使使用索引，也可能会锁住一些不需要的数据
`
锁算法
=======
`
	record lock： 单条记录
	gap lock： 间隙锁，锁定一个范围，不包括记录本身
		// 可重入Gap locks in InnoDB are “purely inhibitive”, which means they only stop other transactions from inserting to the gap. Thus, a gap X-lock has the same effect as a gap S-lock.
		// 表中有11 13 20 10， delete 21，会给20~正无穷加锁，防止插入数据，当两个线程都删除再插入就会死锁
		// https://www.cnblogs.com/fan-yuan/p/7918740.html
	next-key lock： record lock + gap lock
	mysql官网介绍：
	https://dev.mysql.com/doc/refman/5.7/en/innodb-transaction-isolation-levels.html
`

乐观锁和悲观锁：
===
`
悲观锁：
	指悲观的认为，需要访问的数据随时可能被其他人访问或者修改。因此在访问数据之前，对要访问的数据加锁，不允许其他其他人对数据进行访问或者修改。上述讲到的服务器锁和InnoDB锁都属于悲观锁。
乐观锁：
	指乐观的认为要访问的数据不会被人修改。因此不对数据进行加锁，如果操作的时候发现已经失败了，则重新获取数据进行更新（如CAS），或者直接返回操作失败。
`

举例
===
insert 引起的死锁
=====
`
https://blog.csdn.net/varyall/article/details/80219459
`
delete and insert 死锁
=====
`
t1 delete 没数据
t2 delete 没数据
t1 insert 锁等待
t2 insert deadlock
delete 的条件为唯一索引
操作的是不同数据
delete后，因为记录不存在且隔离级别为RR，t1和t2都会加一个gap锁，范围相同，此时插入会获取插入意向锁，t1等待t2释放gap，t2等待t1释放导致死锁
如果delete的数据存在，只会加行锁，不会死锁
`

参考
=====
`
https://blog.csdn.net/Donald_Draper/article/details/88307383
各种情况加锁分析	
http://hedengcheng.com/?p=771#_Toc374698318
https://my.oschina.net/hebaodan/blog/1835966
`