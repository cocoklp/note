* [表2结果作为表1条件](#表2结果作为表1条件)
* [两个表分页](#两个表分页)
* [导入导出数据](#导入导出数据)

# 表2结果作为表1条件
select * from tb1 where tb1_col1 in (select tb2_col1 from tb2 as s where tb2_col2 ='value');

# 两个表分页
## union
连接两个以上的select语句的结果组合到一个结果集合中，多个select语句会删除重复数据
select user_pin from instance_info where valid =1 union select user_pin from instance_usage_info where valid =1 limit 1,3;
## union all
包含重复数据

# 导入导出数据
导出
mysqldump -hhost -Pport -upin -ppassword database tablename --skip-lock-tables > filename
// tablename为空则备份整个db
里面是drop table， create table， insert into等语句，备份表结构和数据
mysqldump -hhost -Pport -upin -ppassword database tablename --skip-lock-tables > filename
只导出表结构
导入
mysql -hhost -Pport -upin -ppassword database tablename --skip-lock-tables < filename


# sql 没有bool型，会自动转为tinyint存储，0-false，1-true

# offset
	随着offset增大，执行效率会下降
	select * from test where val=4 limit 300000,5; 
	执行过程
		查询到索引叶子节点数据
		根据叶子节点上的主键去聚簇索引上查询需要的全部字段值
		然后过滤掉前 300000 条数据，取出最后5条
	select * from test a inner join (select id from test where val=4 limit 300000,5) b on a.id=b.id
			

varchar() 存的是字符，比如varchar(1)可以存一个汉字/一个英文字母