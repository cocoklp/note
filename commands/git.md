Git: fatal: Pathspec is in submodule
    git rm --cached directory
    git add directory



git blame name

git基本命令
git clone
获取代码  //svn checkout
eg: 获取URL为http://git.jd.com/cloudwaf/wafapi/tree/dev的代码
git clone http://git.jd.com/cloudwaf/wafapi.git

git status
当前路径下文件状态，svn diff
新增文件为untracked

git add
添加内容到下一次提交中
开始跟踪新文件，或把已跟踪的文件放到暂存区
git add 文件名：开始跟踪一个文件。
git add 目录：递归跟踪该目录下所有文件

git commit
提交，svn commit
git commit –a：自动把所有跟踪过的文件暂存一并提交，从而跳过git add 步骤

git rm

git log
git log –p 可以查看log的详细信息



git pull
自动抓取并合并远程分支到当前分支
Svn update

git push
git push：将本地库中的最新信息发送给远程库
git commi：将本地修改过的文件提交到本地库


git branch
git branch klp：自己新建分支klp


注册表修改环境变量GOPATH：
运行-regedit
\HKEY_USERS\S-1-5-21-1713849901-2797640346-4150151575-784095\Environment


git reset
Add 多了，用go reset回退
git reset HEAD




初始化
进入远程文件夹，git --bare init –shared
Git clone
Touch 文件名
Git add
Git commit

receive.denyCurrentBranch’ Configuration variable to ‘refuse
git --bare init --share



git rm filename
删除文件
git checkout 新分支报错
$ git checkout test
error: pathspec 'test' did not match any file(s) known to git.


Merge：
修改要merge到test分支
在test分支，Git merge dev，然后git push origin test
上一个git命令中止，无法继续git
今天写完代码git commit -m 一下发现提示
Another git process seems to be running in this repository, e.g.
an editor opened by 'git commit'. Please make sure all processes 
are terminated then try again. If it still fails, a git process 
may have crashed in this repository earlier: 

remove the file manually to continue.
大致意思就是你已经提交一了一个commit，该进程卡在这里了，所以无法继续提交。
上网百度了一下
使用$ rm -f ./.git/index.lock 命令，结束该进程。
即可继续提交。
放弃本地所有修改
尚未add

$ git checkout – filename
$ git checkout . 放弃所有修改

已经使用了  git add 缓存了代码。
可以使用  git reset HEAD filepathname （比如： git reset HEAD readme.md）来放弃指定文件的缓存，放弃所以的缓存可以使用 git reset HEAD . 命令。
此命令用来清除 git  对于文件修改的缓存。相当于撤销 git add 命令所在的工作。在使用本命令后，本地的修改并不会消失，而是回到了如（一）所示的状态。继续用（一）中的操作，就可以放弃本地的修改。
 
三，
已经用 git commit  提交了代码。
可以使用 git reset --hard HEAD^ 来回退到上一次commit的状态。此命令可以用来回退到任意版本：git reset --hard  commitid 
你可以使用 git log 命令来查看git的提交历史。git log 的输出如下,之一这里可以看到第一行就是 commitid：

查看用户名

配置用户名
$ git config --global user.name "weiguobing" 
$ git config --global user.email "guobingwei@aliyun.com"
git config --global --list


pull冲突时：
git stash

强制放弃拉取远程分支
git fetch --all  
git reset --hard origin/master 
git pull

撤销某次commit
git log 找到要撤销的id
git reset id 撤销某次commit但不对代码的修改进行撤销
git reset -hard id 撤销某次commit并将代码回复到前一次commit的版本
注意：id为回退到的版本，如本次commit id 为id1，上一次为id2，则应该为git reset id2
修改commit的message
https://blog.csdn.net/sodaslay/article/details/72948722
git commit –amend
在打开的文件中修改message
git push origin branch
新建分支
git checkout –b local_branch_name
git push origin localbranch_name:remote_branch_name

rebase & reset
到自己的开发分支
Git rebase dev，如有conflic，解决完后git rebase –continue
查看dev的git log，得到最新的commit id
在自己的分支git reset id
然后把修改重新add commit push
http://gitbook.liuhui998.com/4_2.html

分支
删除远程分支后，本地git branch -a仍存在已删除的分支
使用命令 git remote show origin，可以查看remote地址，远程分支，还有本地分支与之相对应关系等信息。此时我们可以看到那些远程仓库已经不存在的分支，根据提示，使用 git remote prune origin 命令：删除远程已不存在的分支
https://blog.csdn.net/zhaoguanghui2012/article/details/75127446
删除本地分支：
git branch –d branchname
删除远程分支：
git push origin --delete <BranchName>
https://blog.csdn.net/qq_32452623/article/details/54340749

git merge
将两个或以上的开发历史合并
从指定的commit合并到当前分支
1 git pull = git fetch + git merge
2 从一个分支到另一个分支的合并
git merge –abort
	抛弃合并过程并尝试重建合并前的状态，但是当合并开始时如存在未commit的文件，则无法重现。
	如果commit了，git merge或git pull冲突时，可以这样回退merge

强制覆盖
方法1git push origin test:master -f           //将test分支强制（-f）推送到主分支master
// 注意看清顺序，master被覆盖
方法2（假设当前位于test分支）
git checkout master                          //将当前分支切换到主分支
git reset -hard test                            //将主分支重置为test分支
git push origin master -force             //将重置后的master分支强制推送到远程仓库


参数：
https://blog.csdn.net/t3/article/details/77069523
代码
下载代码
Git clone
git clone http://git.jd.com/cloudwaf/mid.git
下载下来的是主线，需要切到branch
cd wafapi
git checkout dev
即可
Ssh方式
git clone git@git.jd.com:jdlbproduct/lbconfcenter.git
编译代码
cd mid
./build.sh midserver/main.go

上传代码
Git status 查看状态，修改的什么文件，branch分支
Git add：如果都是自己改的，可直接git add .，否则需要git add 文件名将所有文件add
Git commit –m “log”：本地代码上传到本地库
git config --global user.name "kouliping"
git config --global user.email kouliping@jd.com
git push origin klp：本地库到远程库
 

# lbconfcenter/models
building/src/lbconfcenter/models/admin_m.go:153:14: undefined: common.XEngPoll
building/src/lbconfcenter/models/admin_m.go:158:12: undefined: common.XEngPoll
building/src/lbconfcenter/models/admin_m.go:163:12: undefined: common.XEngPoll
building/src/lbconfcenter/models/admin_m.go:170:8: undefined: idcMeta
building/src/lbconfcenter/models/admin_m.go:172:14: undefined: common.XEngPoll
building/src/lbconfcenter/models/admin_m.go:173:2: too many arguments to return
building/src/lbconfcenter/models/admin_m.go:178:14: assignment mismatch: 2 variables but 1 values
building/src/lbconfcenter/models/admin_m.go:211:2: too many arguments to return
	have ([]proto.ConfdirMeta, error)
	want ()
building/src/lbconfcenter/models/admin_m.go:256:21: cannot refer to unexported name proto.confdirMeta
building/src/lbconfcenter/models/admin_m.go:256:21: undefined: proto.confdirMeta
building/src/lbconfcenter/models/admin_m.go:256:21: too many errors
[kouliping@A01-R04-I17-130-06CEKFK lbconfcenter]$







