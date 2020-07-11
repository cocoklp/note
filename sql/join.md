join
https://www.cnblogs.com/fudashi/p/7491039.html
笛卡尔积，将A表的每一条记录和B表的每一条记录强行拼在一起。A有n条，B有m条，则笛卡尔积有m*n条

select *From user_test, user;
inner join:
	// 从结果中挑出 on条件成立的部分
	select *From user_test join user on user_test.id=user.id;

left join:
	// 两表交集外加左表剩余部分
	select *From user_test left join user on user_test.id=user.id;
right join:
	// 两表交集外加右表剩余部分
	select *From user_test right join user on user_test.id=user.id;
outer join:
	// 两表并集
	select *From user_test right join user on user_test.id=user.id
	union
	select *From user_test left join user on user_test.id=user.id;
using:
	MySQL中连接SQL语句中，ON子句的语法格式为：table1.column_name = table2.column_name。当模式设计对联接表的列采用了相同的命名样式时，就可以使用 USING 语法来简化 ON 语法，格式为：USING(column_name)。
所以，USING的功能相当于ON，区别在于USING指定一个属性名用于连接两个表，而ON指定一个条件。另外，SELECT *时，USING会去除USING指定的列.