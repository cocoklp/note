package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	fmt.Println("vim-go")
	src := os.Args[1]
	dst := os.Args[2]
	fmt.Println(CpFileAtomic(src, dst, 0755))
}

func WriteFileAtomic(filename string, data []byte, perm os.FileMode) error {
	dir, name := path.Split(filename)
	f, err := ioutil.TempFile(dir, name)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err == nil {
		err = f.Sync()
	}
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	if permErr := os.Chmod(f.Name(), perm); err == nil {
		err = permErr
	}
	if err == nil {
		err = os.Rename(f.Name(), filename)
	}
	// Any err should result in full cleanup.
	if err != nil {
		os.Remove(f.Name())
	}
	return err
}

func CpFileAtomic(srcFile string, destFile string, perm os.FileMode) error {
	data, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return err
	}
	return WriteFileAtomic(destFile, data, perm)
}
func CmdRunWithTimeout(cmd *exec.Cmd, timeout time.Duration) (error, bool) {
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
		// timeout
		if err = cmd.Process.Kill(); err != nil {
			glog.Errorf("failed to kill: %s, error: %s", cmd.Path, err)
		}
		go func() {
			<-done // allow goroutine to exit
		}()
		glog.Infof("process: `%s %s` killed", cmd.Path, strings.Join(cmd.Args, " "))
		//fmt.Printf("process:`%s %s` killed\n", cmd.Path, strings.Join(cmd.Args, " "))
		return err, true
	case err = <-done:
		return err, false
	}
}

func CmdCopy(src, dest string) (string, error) {
	if !dir.IsExist(src) {
		return "", fmt.Errorf("src [%s] is not exist.", src)
	}
	if dir.IsExist(dest) {
		if err := os.RemoveAll(dest); err != nil {
			return "", err
		}
	}

	cmdStr := fmt.Sprintf("cp -r %s %s", src, dest)
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Start() // attention!

	var isTimeout bool
	timeout := 60
	err, isTimeout := CmdRunWithTimeout(cmd, time.Duration(timeout)*time.Second)
	cmdOutput := "stdout:[" + stdout.String() + "]stderr:[" + stderr.String() + "]"

	if err != nil {
		return cmdOutput, err
	}
	if isTimeout {
		return cmdOutput, fmt.Errorf("err: Exec timeout %s", cmdStr)
	}

	return cmdOutput, nil
}

package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	//"io/ioutil"
	//"bufio"
	//"io"
	"os"
	"time"
)

var (
	user     = "root"
	password = "cmcc12#$"
	host     = "172.30.204.107"
	port     = "22"
)

var (
	HaRemoteConfFile = "/root/klp/haproxy.cfg"
	HaLocalConfFile  = "haproxy.cfg"
)

func main() {
	err := SyncConfFileForBackup()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}

func SyncConfFileForBackup() error {
	session, err := NewSftpSession(user, password, host, port)
	if err != nil {
		return err
	}
	defer session.Close()
	err = TransFile(session, HaLocalConfFile, HaRemoteConfFile)
	return err
}

func NewSftpSession(sshUser, sshPassword, sshHost, sshPort string) (*sftp.Client, error) {
	sshClient, err := newSshClient(sshUser, sshPassword, sshHost, sshPort)
	if err != nil {
		return nil, err
	}
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return nil, err
	}
	return sftpClient, nil
}

func TransFile(client *sftp.Client, localFile, remoteFile string) error {
	srcFile, err := os.Open(localFile)
	if err != nil {
		return fmt.Errorf("open file fail %s", err.Error())
	}

	defer srcFile.Close()

	dstFile, err := client.Create(remoteFile)
	if err != nil {
		return fmt.Errorf("create client fail %s", err.Error())
	}
	defer dstFile.Close()

	for {
		buf := make([]byte, 1024)
		n, _ := srcFile.Read(buf)
		fmt.Printf("[%d][%s]\n", n, string(buf))
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}

	/*
		ff, err := ioutil.ReadAll(srcFile)
		if err != nil {
			return fmt.Errorf("ReadAll fail %s", err.Error())
		}
		dstFile.Write(ff)
	*/
	/*
		rd := bufio.NewReader(srcFile)
		for {
			line, err := rd.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			dstFile.Write([]byte(line))
		}
	*/
	return nil
}

func newSshClient(sshUser, sshPassword, sshHost, sshPort string) (*ssh.Client, error) {
	//创建sshp登陆配置
	config := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%s", sshHost, sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		fmt.Println("创建ssh client 失败 [%s]", err.Error())
		return nil, err
	}
	return sshClient, nil
}
