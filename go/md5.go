package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

}

/*
	占用内存最多，把整个文件读入内存
*/
func md5sum1(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5.Sum(data))
}

/*
	占用内存较少，一般情况下io.Copy() 每次分配32*1024字节的内存
*/
func md5sum2(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

/*
	bufio.Reader 默认创建4096字节的buffer
*/
func md5sum3(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	h := md5.New()
	_, err = io.Copy(h, r)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

/*
	写文件的同时计算md5
*/
func md5sum4(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 创建文件
	dest, err := os.Create("filename")
	if err != nil {
		return "", err
	}
	hMd5 = md5.New()
	// multiwriter 多实例写入，同时写入多个buffer
	multiWriter = io.MultiWriter(hMd5, respFile)
	_, err = io.Copy(multiWriter, resp.Body)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil

}
