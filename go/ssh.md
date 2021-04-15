# 远程写文件

```
package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("vim-go")
	sftpClient, err := connect("root", "2020.08.29Hl", "114.67.71.44", 22)
	if err != nil {
		panic("connetc fail" + err.Error())
	}
	defer sftpClient.Close()
	remote := os.Args[1]
	localFilePath := "/etc/haproxy/haproxy.cfg"
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		panic("open file fail " + err.Error())
	}

	dstFile, err := sftpClient.Create(remote)
	if err != nil {
		panic("create client fail " + err.Error())
	}
	defer dstFile.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}

	fmt.Println("copy file to remote server finished!")
}

func connect(user, password, host string, port int) (*sftp.Client, error) {
	config := &ssh.ClientConfig{
		User:    user,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	config.Auth = []ssh.AuthMethod{ssh.Password(password)}

	// connet to ssh
	addr := fmt.Sprintf("%s:%d", host, port)

	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, err
	}

	// create sftp client
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return nil, err
	}

	return sftpClient, nil
}
```

