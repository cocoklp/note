

# 虚拟机

1. 安装完成后，网络模式设置为nat模式
2. 查看主机 ip配置，cmd 中ipconfig，关注以太网的ipaddr netmask gateway，据此配置虚拟机
3. ifconfig查看网卡，ens-33

动态获取，网关2结尾
https://segmentfault.com/q/1010000002417223

```
vim /etc/sysconfig/network-scripts/ifcfg-ens-33
TYPE="Ethernet"
PROXY_METHOD="none"
BROWSER_ONLY="no"
BOOTPROTO="static"  # dhcp 修改为static
DEFROUTE="yes"
IPV4_FAILURE_FATAL="no"
IPV6INIT="yes"
IPV6_AUTOCONF="yes"
IPV6_DEFROUTE="yes"
IPV6_FAILURE_FATAL="no"
IPV6_ADDR_GEN_MODE="stable-privacy"
NAME="ens33"
UUID="9e251664-9029-47fb-b2d3-ab3f13257ec4"
DEVICE="ens33"
ONBOOT="yes"
IPADDR=10.2.43.205
NETMASK=255.255.255.0
GATEWAY=10.2.43.254
DMS1=8.8.8.8
```

4. 刷新网络配置

   service network restart

   nmcli c reload ens-33

   如果ifconfig没有ens，ifconfig -a 找到，然后ifconfig ens-33 up，此时只有v6地址没有v4，然后ifconfig ens-33 ip，nmcli c reload ens-33

5. 配置代理

   firefox->preferences->network proxy 如果找不到代理配成ipport格式

   对于centos7，还需要在终端配置

   vim /etc/enviroment 输入代理然后source刷新，env查看代理。同样如果不成功(访问时提示can not resolv，则配置成ipport格式)

   

虚拟机git clone代码要删除代理配置

```
unset http_proxy
```

https://blog.csdn.net/redkeyer/article/details/96125735

# vim





# go源码安装

下载源码，解压到/usr/local目录

修改~/.bash_profile，配置GOPATH GOROOT PATH，然后source 文件生效即可

# vim安装

下载vim tar.gz安装包

解压tar -zxvf  v8.1.1766.tar.gz 

下载后安装依赖

rpm -ivh ncurses-devel-5.9-14.20130511.el7_4.x86_64.rpm

安装配置

cd vim-8.1.1766/ ./configure --prefix=/usr/local&&make && make install

配置指向

alias vim='/usr/local/bin/vim' echo "alias vim='/usr/local/bin/vim' " >> ~/.bashrc

显示行号： vim ~/.vimrc 添加set nu即可





# vim-go

bundle vimrc等文件 //传到git

vim 进入编辑器， :PluginInstall



go install golang.org/x/tools/cmd/guru/
  go install golang.org/x/tools/gopls/
  go install github.com/davecgh/reftools/cmd/fillstruct/
    go install github.com/fatih/motion/
  go install github.com/rogpeppe/godef
  go install github.com/kisielk/errcheck/
  go install github.com/go-delve/delve/cmd/dlv/
  go install golang.org/x/tools/cmd/gorename/
  go install github.com/koron/iferr

go install 后的二进制在 $GOPATH/bin目录下，需要把改目录配置到PATH中

GoInstallBinaries



#安装docker

离线安装包

rpm安装后，service docker start

scp docker-compose /usr/local/bin

chmod +x /usr/local/bin/docker-compose

ln -s /usr/local/bin/docker-compose /use/bin/docker-compose

安装harbor

https://goharbor.io/docs/2.1.0/install-config/
<<<<<<< HEAD



# 添加快捷方式到桌面

http://www.zh-go.com/index.php/2019/10/20/centos-7-%E5%AE%89%E8%A3%85goland/

vim /root/桌面/goland.desktop
内容如下：

```go
[root@zdwork ~]# cat 桌面/goland.desktop
[Desktop Entry]
Version=1.0
Encoding=UTF-8
Name=JetBrains GoLand 2019.2.3
Type=Application
Terminal=false
Name[en_US]=JetBrains GoLand 2019.2.3
Exec=/usr/local/src/GoLand-2019.2.3/bin/goland.sh
Comment[en_US]=GoLand-2019.2.3
Comment=GoLand-2019.2.3
GenericName[en_US]=
Icon=/usr/local/src/GoLand-2019.2.3/bin/goland.png
[root@zdwork ~]#
```
=======
>>>>>>> 8506f591f8c81b37c2c45be48caa3bef22dd9815
