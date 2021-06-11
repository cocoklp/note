# 安装

## 在线安装

https://www.postgresql.org/download/linux/redhat/

安装完设置

sudo -u postgres psql

ALTER USER postgres WITH PASSWORD 'postgres';  设置密码

允许远程访问

vim /var/lib/pgsql/13/data/postgresql.conf

listen_addresses = '*'

/var/lib/pgsql/13/data/pg_hba.conf

在该配置文件的host all all 127.0.0.1/32 md5行下添加以下配置，或者直接将这一行修改为以下配置
host all all 0.0.0.0/0 md5
如果不希望允许所有IP远程访问，则可以将上述配置项中的0.0.0.0设定为特定的IP值。

然后systemctl restart postgresql-13



## 离线安装

https://www.cnblogs.com/sunalways/p/13817939.html

下载地址：

https://www.postgresql.org/download/

下拉选择direct rpm download

![image-20210512114201238](C:\Users\kouliping\AppData\Roaming\Typora\typora-user-images\image-20210512114201238.png)

选择版本和系统

![image-20210512114339006](C:\Users\kouliping\AppData\Roaming\Typora\typora-user-images\image-20210512114339006.png)



![image-20210512114402297](C:\Users\kouliping\AppData\Roaming\Typora\typora-user-images\image-20210512114402297.png)

下载以上rpm包，然后rpm -ivh安装

```
rpm -ivh postgresql13-libs-13.2-1PGDG.rhel7.x86_64.rpm 
rpm -ivh postgresql13-13.2-1PGDG.rhel7.x86_64.rpm 
rpm -ivh postgresql13-server-13.2-1PGDG.rhel7.x86_64.rpm 
```

初始化

```
/usr/pgsql-13/bin/postgresql-13-setup initdb
```

启动

```
systemctl enable postgresql-13
systemctl start postgresql-13
```

修改ip和port信息

```
vi /var/lib/pgsql/12/data/postgresql.conf
# *表示监听所有的ip信息，也可以使用localhost、127.0.0.1等
listen_addresses = '*'  
# 默认的监听端口为5432，也可以换为其它的
port = 5432
```



 sudo -u postgres psql

登陆后创建用户和db

 create database mepm;

create user mepm with password '123456';



# 权限



# 命令

## 登录

```
psql -U username -d dbname -h hostip -p port
```

## 用户

创建用户：

create user mepm with replication createdb login password '123456';

查看用户

\du

# 数据库

create database dbname;

\l 查看所有数据库

\c dbname 切换数据库



##表

\d 查看当前db的表

\d tablename table所有字段

\d+ tablename 表的基本情况，包括索引

