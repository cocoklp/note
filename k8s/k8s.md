* [介绍说明](#介绍说明)
	* [发展历史](##发展历史)
	* [框架](##kubernetes框架)
	* [关键字](##关键字)
* [基本概念](#基本概念)
    * [pod](##pod)
    * [控制器类型](##控制器类型)
    * [网络通讯模式](##网络通讯模式)
* [k8s集群](#k8s集群)
    * [构建k8s](##构建k8s)
* [资源清单](#资源清单)
    * [什么是资源](##什么是资源)
    * [资源清单的语法](##资源清单的语法)
    * [编写pod](##编写pod)
    * [pod生命周期](##pod生命周期)
* [pod控制器](#pod控制器)
    * [什么是控制器](##什么是控制器)
    * [控制器类型说明](##控制器类型说明)
* [服务SVC发现](#服务SVC发现)
    * [SVC原理及构建方式](##SVC原理及构建方式)
* [无状态服务](#无状态服务)
    * [SVC原理及构建方式](##SVC原理及构建方式)
* [存储](#存储)
    * [多种存储类型特点](##多种存储类型特点)
    * [存储方案的选择](##存储方案的选择)
* [调度器](#调度器)
    * [调度器原理](##调度器原理)
    * [调度实践](##调度实践)
* [集权安全机制](#集群安全机制)
    * [认证](##认证)
    * [鉴权](##鉴权)
    * [访问控制原理及流程](##访问控制原理及流程)
* [HELM](#HELM)
    * [原理](##原理)
    * [模板自定义](##模板自定义)
    * [部署](##部署)
* [运维](#运维)
    * [kubeadm源码修改](##kubeadm源码修改)
    * [kubeadm高可用构建](##kubeadm高可用构建)


## 发展历史
#服务发现
	pod不能直接提供给客户端访问（集群内部是私有ip），通过服务发现把服务暴露给客户端。客户端根据暴露的地址+端口来访问。
	统一的访问入口。


#无状态服务
	有状态服务： DBMS。节点摘除后再加入不能继续正常服务。
	无状态服务： LVS APACHE（单个节点不需要存储数据）。

#调度器
	k8s把容器/pod调度到对应的节点。
#HELM
	linux的yum安装工具，通过命令部署服务。

## kubeadm源码修改
	证书


## pod

#自主式pod

#控制器管理的pod
    ReplicatoinController（RC）： 
    保证容器副本数为期望值。有pod退出会创建新的pod。多了会回收。
    ReplicationSet（RS）：
    同RC，支持集合式的selector。不支持rolling-update
    Deployment：
    自动管理RS。支持rolling-update。不支持pod创建。
    deployment创建后会创建一个rs，再创建pod。
    HPA（horizontal pod autoscaling）
    仅适用于deployment和rs。监控cpu利用率，当cpu>80%时，创建新的pod，最多创建10个。利用率变低后会回收，最少剩两个。
    StatefulSet：
    解决有状态服务的问题。
    稳定的持久化存储。
    稳定的网络标识（新pod替代就pod时，podname和hostname不变）
    有序部署：前一个pod处于running或ready状态才会创建新的。服务依赖有启动顺序。由下向上创建，由上向下回收。
    有序回收：有序回收
    daemonset：（守护进程）
    确保全部或一些node上运行一个pod副本。
    job，corn job
    job：批处理任务，仅执行一次的任务。
    corn job： 基于时间。




borg架构：
	borgmaster 多个，奇数个
	borglet  多个
	 
组件：
master
    ```
    apiserver：
        所有服务访问的统一入口
    controller manage：
        维护副本的期望数目
    scheduler： 
        调度器，接受任务选择合适的节点进行分配任务
        (节点资源足够)
    etcd:
        kv存储，存储k8s集群的所有重要信息（持久化）
    ```
node
```
kubelet:
    直接与容器交互实现容器的生命周期管理
kubeproxy:
    写入规则至ip tables/ipvs 实现服务映射访问
```
coredns
    为集群中的svc创建A记录
dashbord
    b/s（Browser/Server）结构，给k8s集群提供bs结构的访问体系。
ingress controller：
    官方实现四层代理，ingress实现七层代理。根据域名实现负载均衡。
federation：
    跨集群中心多k8s统一管理功能
prometheus：
    提供一个k8s集群的监控能力
elk：
    提供k8s集群日志统一分析接入平台


网络通讯模式
    所有pod都在一个可以直接联通的扁平的网络空间中。所有的pod都可以通过对方的ip“直接”到达。
    GCE：
    同一pod里的多个容器： 共享 pause的协议栈，通过localhost可访问
    各pod之间的通讯：
        在同一台机器： 由docker0网桥直接转发请求到pod
        不在： flannel
    
    pod与service： LVS
    pod到外网：
        查找路由表。
    外网访问pod：
        service。

    Flannel： 集群中不同节点主机创建的docker容器都具有全集群唯一的虚拟IP地址，并且在ip地址之间建立一个覆盖网络，将数据包原封不动地传递到目标容器内。
    etcd flannel
    etcd 存储管理 flannel 可分配的ip地址段资源。
    监控etcd中 每个pod的实际地址，并在内存中建立维护 pod 节点路由表。

安装集群

修改主机名和host文件
hostnamectl set-hostname k8s-master
安装依赖包
yum install -y conntrack ntpdate ntp ipvsadm ipset jq iptables curl sysstat libseccomp wget vim net-tools git
设置防火墙
systemctl stop firewalld && systemctl disable firewalld
yum -y install iptables-services && systemctl start iptables && systemctl enable iptables && iptables -F && service iptables save

关闭setlinux  // 防止容器运行在虚拟内存
swapoff -a && sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab
setenforce 0 && sed -i 's/^SELINUX=.*/SELINUX/disabled/' /etc/selinux/config

调整内核参数
cat > kubernetes.conf <<EOF
net.bridge.bridge-nf-call-iptables=1
net.bridge.bridge-nf-call-ip6tables=1
net.ipv4.ip_forward=1
net.ipv4.tcp_tw_recycle=0
vm.swappiness=0
vm.overcommit_memory=1
vm.panic_on_oom=0
fs.inotify.max_user_instances=8192
fs.inotify.max_user_watches=1048576
fs.file-max=52706963
fs.nr_open=52706963
net.ipv6.conf.all.disable_ipv6=1
net.netfilter.nf_conntrack_max=2310720
EOF
modprobe br_netfilter
cp kubernetes.conf /etc/sysctl.d/kubernetes.conf
sysctl -p /etc/sysctl.d/kubernetes.conf

网络参数解释
https://blog.csdn.net/wangxizhen123/article/details/72850395

调整系统时区
# 将时区设置为 中国/上海
timedatectl set-timezone Asia/Shanghai
# 将当前的UTC时间写入硬件时钟
timedatectl set-local-rtc 0
# 重启依赖于系统时间的服务
systemctl restart rsyslog
systemctl restart cornd
关闭系统不需要的服务
systemctl stop postfix && systemctl disable postfix
设置rsyslogd 和 systemd journald
mkdir /var/log/journal  # 持久化保存日志的目录
mkdir /etc/systemd/journald.conf.d
cat > /etc/systemd/journald.conf.d/99-prophet.conf << EOF
[Journal]
# 持久化保存到磁盘
Storage=persistent
# 压缩历史日志
Compress=yes
SyncIntervalSec=5m
RateLimitInterval=30s
RateLimitBurst=1000

# 最大占用空间 10G
SystemMaxUse=10G

# 单日志文件最大 200M
SystemMaxFileSize = 200M

# 日志保存时间2周
MaxRetentionSec=2week

#不将日志转发到syslog
ForwardToSyslog=no
EOF
systemctl restart systemd-journald

升级系统内核为4.44
自带的3.10.x内核存在bug导致kubenetes和docker不稳定
`
查看内核版本命令 uname -r
`
rpm -Uvh http://www.elrepo.org/elrepo-release-7.0-3.el7.elrepo.noarch.rpm
#安装完成后检查 /boot/grub2/grub.cfg 中对应内核 menuentry是否包含initrd16配置，如果没有再安装一次。
yum --enablerepo=elrepo-kernel install -y kernel-lt
#设置开机从新内核启动
grub2-set-default "CentOS Linux(4.4.182-1.el7.elrepo.x86_64) 7 (Core)"


`
查看当前内核版本
awk -F\' '$1=="menuentry " {print i++ " : " $2}' /etc/grub2.cfg
0 : CentOS Linux (3.10.0-1127.19.1.el7.x86_64) 7 (Core)
1 : CentOS Linux (4.4.240-1.el7.elrepo.x86_64) 7 (Core)
2 : CentOS Linux (3.10.0-957.el7.x86_64) 7 (Core)
3 : CentOS Linux (0-rescue-7deb1fe2d78d4280b703c56798959e4b) 7 (Core)
设置默认版本
grub2-set-default 1  // 这个数字是awk的结果
`

kube-proxy开启ipvs的前置条件
modprobe br_netfilter
cat > /etc/sysconfig/modules/ipvs.modules <<EOF
#!/bin/bash
modprobe -- ip_vs
modprobe -- ip_vs_rr
modprobe -- ip_vs_wrr
modprobe -- ip_vs_sh
modprobe -- nf_conntrack_ipv4
EOF
chmod 755 /etc/sysconfig/modules/ipvs.modules && bash /etc/sysconfig/modules/ipvs.modules && lsmod |grep -e ip_vs -e conntrack_ipv4

安装docker软件
yum install -y yum-utils device-mapper-persistent-data lvm2

镜像仓库
yum-config-manager \
  --add-repo \
  http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
更新系统，安装docker ce
yum update -y && yum install -y docker-ce
此时重启，又回到了默认的3.10.0版本。需要重新设置默认版本再重启
grub2-set-default "CentOS Linux(4.4.182-1.el7.elrepo.x86_64) 7 (Core)" && reboot

启动docker并开机自启
systemctl start docker
systemctl enable docker


创建 /etc/docker 目录
mkdir /etc/docker

配置daemon
cat > /etc/docker/daemon.json << EOF
{
    "exec-opts":["native.cgroupdriver=systemd"],
    "log-driver":"json-file",
    "log-opts":{
        "max-size":"100m"
    }
}
EOF
mkdir -p /etc/systemd/system/docker.service.d
重启docker服务
systemctl daemon-reload && systemctl restart docker && systemctl enable docker

安装kubeadm（主从配置）
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=http://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=http://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg
http://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF

yum -y install kubeadm-1.15.1 kubectl-1.15.1 kubelet-1.15.1
systemctl enable kubelet.service


安装镜像
https://zhuanlan.zhihu.com/p/46341911

初始化主节点
vim kubeadm-config.yaml 
localAPIEndpoint:
  advertiseAddress: 192.168.0.35
kubernetesVersion: v1.15.1
networking:
  dnsDomain: cluster.local
  podSubnet: "10.244.0.0/16"
  serviceSubnet: 10.96.0.0/12
scheduler: {}
apiVersion: kubeproxy.config.k8s.io/v1alpha1
kind: KubeProxyConfiguration
featureGates:
  SupportIPVSProxyMode: true
mode: ipvs

kubeadm init --config=kubeadm-config.yaml --experimental-upload-certs | tee kubeadm-init.log
 mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
// 修改用户权限 export KUBECONFIG=/etc/kubernetes/admin.conf



部署网络
mkdir install-k8s
mv kubeadm-init.log kubeadm-config.yaml install-k8s/
cd install-k8s/
mkdir core
mv * core/
mkdir plugin
cd plugin/
mkdir flannel
cd flannel/
wget https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
kubectl create -f kube-flannel.yml
kubectl get pod -n kube-system

kubectl 命令行工具
kubectl 命令 -n namespace  所有组件默认安装在 kube-system， 通过-n指定，不指定的话用default

node节点加入集群
kubeadm join 192.168.0.35:6443 --token 5ogo68.y3oigcb8zpdmpqyw     --discovery-token-ca-cert-hash sha256:67fa3860b6dffc9d96a2aea4be8a0de20950ea76e6d32038bd2a83e95691c762
如果卡住，或者报错
`
error execution phase preflight: couldn't validate the identity of the API Server: abort connecting to API servers after timeout of 5m0s
`
原因是token过期，重新生成一份再重新join即可。
kubeadm token create
 openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'

