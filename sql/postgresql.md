# 安装

https://www.postgresql.org/download/linux/redhat/

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

