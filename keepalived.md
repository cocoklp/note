# 安装

https://blog.csdn.net/bbwangj/article/details/80346428

 yum -y install openssl-devel

```
1. [root@master src]# pwd
2. /usr/local/src
3. [root@master src]# wget http://www.keepalived.org/software/keepalived-2.0.7.tar.gz
4. [root@master src]# tar xvf keepalived-2.0.7.tar.gz
5. [root@master src]# cd keepalived-2.0.7
6. [root@master keepalived-2.0.7]# ./configure --prefix=/usr/local/keepalived
7. [root@master keepalived-2.0.7]# make && make install
```



## 启动 & 初始化

```
# keepalived启动脚本变量引用文件，默认文件路径是/etc/sysconfig/，也可以不做软链接，直接修改启动脚本中文件路径即可（安装目录下）
[root@localhost /]# cp /usr/local/keepalived/etc/sysconfig/keepalived  /etc/sysconfig/keepalived 
 
# 将keepalived主程序加入到环境变量（安装目录下）
[root@localhost /]# cp /usr/local/keepalived/sbin/keepalived /usr/sbin/keepalived
 
# keepalived启动脚本（源码目录下），放到/etc/init.d/目录下就可以使用service命令便捷调用
[root@localhost /]# cp /usr/local/src/keepalived-2.0.7/keepalived/etc/init.d/keepalived  /etc/init.d/keepalived
 
# 将配置文件放到默认路径下
[root@localhost /]# mkdir /etc/keepalived
[root@localhost /]# cp /usr/local/keepalived/etc/keepalived/keepalived.conf /etc/keepalived/keepalived.conf
```



```
加为系统服务：chkconfig --add keepalived
开机启动：chkconfig keepalived on
查看开机启动的服务：chkconfig --list
启动、关闭、重启service keepalived start|stop|restart
```







