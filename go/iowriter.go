package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

/*
   使用go下载文件
*/
func download(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 读取到内存
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	ioutil.WriteFile("filename", data, 0644)
}

func downloadlarge(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 创建文件
	dest, err := os.Create("filename")
	if err != nil {
		return err
	}
	// 两个文件指针之间的内容拷贝，省去了写入内容然后由内存写入文件的过程
	_, err := io.Copy(dest, resp.Body)
	if err != nil {
		return err
	}
}
