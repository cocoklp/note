# 安装

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

