package main

import (
	"fmt"
	"io/ioutil"
	"log"

	//"os"
	"os/exec"
	"runtime"
	"time"
)

func ping() {
	ip := "127.0.0.1"
	cmd := exec.Command("ping", "-c5", "-w1", ip)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(opBytes))
}

func main() {
	fmt.Println(runtime.GOARCH)
	err1 := fmt.Errorf("err1")
	err2 := fmt.Errorf("err2")
	errret := fmt.Errorf("%s", fmt.Sprintf("%s %s", err1.Error(), err2.Error()))
	fmt.Println(err1, err2, errret)
	count := 0
	for {
		ping()
		count++
		fmt.Println(count, "~~~", time.Now())
		time.Sleep(time.Duration(2) * time.Second)
	}
}
