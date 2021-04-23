Mysql
登陆数据库
mysql -h172.28.13.159 -P3306 -uroot -p1qaz@WSX
use jd_lb_db_test2
show tables  // 显示当前数据库都有哪些表
create创建表
drop table table-name删除表





授权
 grant all PRIVILEGES on db_name.* to 'username'@'xxx.xxx.xx.x' identified by 'password' WITH GRANT OPTION;
上面的语句表示将数据库 db_name 的所有权限授权给 username 这个用户，允许 username 用户在 xxx.xxx.xx.x 这个 IP 进行远程登陆，并设置 username 用户的密码为 password。
分析参数：
all PRIVILEGES 表示赋予所有的权限给指定用户，这里也可以替换为赋予某一具体的权限，例如：select,insert,update,delete,create,drop 等，具体权限间用“,”半角逗号分隔。
db_name.* 表示上面的权限是针对于哪个表的，db_name指的是数据库名称，后面的 * 表示对于所有的表，由此可以推理出：对于全部数据库的全部表授权为“*.*”，对于某一数据库的全部表授权为“数据库名.*”，对于某一数据库的某一表授权为“数据库名.表名”。
username表示你要给哪个用户授权，这个用户可以是存在的用户，也可以是不存在的用户。
xxx.xxx.xx.x 表示允许远程连接的 IP 地址，你的IP，如果想不限制链接的 IP 则设置为“%”即可。
password 为用户username的密码。
最后执行了上面的语句后，一般都会立即生效，返回值如下：
Query OK, 0 rows affected (0.01 sec)

如果没有上面的语句那么请执行下面的命令，即可立即生效。
Mysql> flush privileges



数据
1.	更新记录：
Update 表名 set 字段=** where 条件;
2.	插入记录
insert into ssl_meta (`ssl`,not_before,not_after) value("klp1","2006-09-22","2009-03-12");
// 以分号结尾

3.	显示数据
select *from table name where 筛选条件

4.	插入数据
insert into ssl_meta (`ssl`,not_before,not_after) value("klp1","2006-09-22","2009-03-12");
insert into命令用于向表中插入数据。

insert into命令格式：insert into <表名> [(<字段名1>[,..<字段名n > ])] values ( 值1 )[, ( 值n )];
注：ssl是关键字，需要加反引号
复制
https://www.cnblogs.com/lxboy2009/p/7234535.html
表结构
删除某field
MySQL [jd_lb_db_test2]> alter table vip_meta drop column utime, drop route_cluster;
Query OK, 0 rows affected (0.18 sec)
Records: 0  Duplicates: 0  Warnings: 0

显示表结构
MySQL [jd_lb_db_test2]> desc vip_meta;

增加某field
MySQL [jd_lb_db_test2]> alter table vip_meta add conf_dir varchar(20) NOT NULL ;

修改fieldming

alter table test rename test1; --修改表名

alter table test add  column name varchar(10); --添加表列

alter table test drop  column name; --删除表列

alter table test modify address char(10) --修改表列类型
||alter table test change address address  char(40)


alter table test change  column address address1 varchar(30)--修改表列名
https://www.cnblogs.com/jiangxiaobo/p/6110679.html

修改键属性：
alter table port_meta  drop key port_UNIQUE;

修改表名
RENAME TABLE old_table_name TO new_table_name;


Mysql备份
mysqldump -hm2836m.mysql.jddb.jcloud.com -P3358 -uwaf_mid_rw -pJdnWRQYEq2OqeVAFJAtx waf_mid eslog_download_meta  --skip-lock-tables > eslogbak.sql;	
	mysql -hm2836m.mysql.jddb.jcloud.com -P3358 -uwaf_mid_rw -pJdnWRQYEq2OqeVAFJAtx waf_mid < eslogbak.sql
文件里是drop create insert语句
加条件的导出
--where='id < 10'


只导出表结构不导出数据：
mysqldump　-d　-A　--add-drop-table　-uroot　-p　>xxx.sql 

只导出数据不导出结构
mysqldump　-d　-A　--add-drop-table　-uroot　-p　>xxx.sql 

Excel导入到mysql
load data local infile "/export/home/kouliping/files/region.csv" into table `dn_region` FIELDS TERMINATED BY ',' LINES TERMINATED BY '\r\n';
文件里要有id


Mysql切换调整字段顺序，不改变索引哦，完美解决方案！
mysql> alter table student modify id int(10) unsigned auto_increment first;
mysql> alter table student modify name varchar(10) after id;
alter table qq_admin add unique key(admin_name);

拷贝数据库
跨服务器：
服务器A,B都安装了mysql，然后需要把服务器A上的数据库DB1都拷贝到服务器B上。
1st：在服务器A上操作
mysqldump -u 用户名 -p 密码 数据库名 [表名] > 导出的文件名
	// 不输入表名参数则把该数据库下的所有表都拷贝
2nd: 在服务器B上，scp 把导出的文件拷贝过来
3rd： 服务器B上把数据导入数据库
mysql –u 用户名 p 密码
mysql>use 数据库
mysql> source temp.sql;// 文件名

字符集

show variables like '%char%';
php不加set names `charname`时所使用的编码
show global variables like '%char%'; 
字符集：character set，不同字符集规定了不同的字符的编码方式，一个character set 是一组符号和编码eg：ascii 编码方式是7bit表示一个字符，非英文不能使用ascii编码。Utf8对世界上所有的语言文字做了编码，
字符集级别：
服务器
数据库
数据表
表列
自下向上继承
最终使用字符集的是存储字符的列，服务器级 数据库级 数据表级 的字符集都是列的字符集的默认选项
1.	在配置文件my.cnf中 [client] 增加 default-character-set =utf8 ,会立即对本机上的新创建连接生效

2.在配置文件my.cnf中 [mysqld] 增加 default-character-set =utf8 ，待mysqld服务重新启动后生效

3. 执行SET语句修改字符集，对非本机新创建的连接也会生效

SET GLOBAL character_set_clinet=utf8;
SET GLOBAL character_set_connection=utf8;
SET GLOBAL character_set_database=utf8;
SET GLOBAL character_set_results=utf8;
SET GLOBAL character_set_server=utf8;

4.对于之前的连接线程，则没有办法，除非他们自己设置为utf8 或者等待其断开重新连接
修改databases名称
https://blog.csdn.net/tjcwt2011/article/details/79451764
删除database
 drop database databasename

设置字符集
set character_set_server=utf8;
set character_set_database=utf8;
show variables like '%char%';

