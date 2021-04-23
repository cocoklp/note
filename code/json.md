```
import("encoding/json")
type Student struct {
    StudentId      string `json:"sid"`  // 使用sid作为key
    StudentName    string 				// 使用field名作为key，即StudentName
    StudentClass   string `json:"class,omitempty"`  // 使用class作为key，且class为空时忽略该字段
    StudentTeacher string `json:"-"`  // 直接忽略该字段
}
```



Unmarshal 是怎么找到结构体中对应的值呢？比如给定一个 JSON key Filed，它是这样查找的：
    首先查找 tag 名字（关于 JSON tag 的解释参看下一节）为 Field 的字段
    然后查找名字为 Field 的字段
    最后再找名字为 FiElD 等大小写不敏感的匹配字段。
    如果都没有找到，就直接忽略这个 key，也不会报错。这对于要从众多数据中只选择部分来使用非常方便



结构体嵌套的情况

```
package main

import (
	"encoding/json"
	"fmt"
)

type AppPostInfo struct {
	AuthInfo `json:"authInfo"`
	CapInfo  []CapPostInfo `json:"capInfo"`
}

type AuthInfo struct {
	Credential `json:"credential"`
}

type Credential struct {
	AccessKeyId string `json:"accessKeyId"`
	SecretKey   string `json:"secretKey"`
}

type CapPostInfo struct {
	OperationType string `json:"operationType"`
	GroupId       string `json:"groupId"`
	Domain        string `json:"domain"`
}

type AppPostInfo1 struct {
	AuthInfo
	CapInfo []CapPostInfo
}

type AuthInfo1 struct {
	Credential
}

type Credential1 struct {
	AccessKeyId string `json:"accessKeyId"`
	SecretKey   string `json:"secretKey"`
}

type CapPostInfo1 struct {
	OperationType string `json:"operationType"`
	GroupId       string `json:"groupId"`
	Domain        string `json:"domain"`
}

func main() {
	{
		typ := AppPostInfo{}
		data, _ := json.Marshal(typ)
		fmt.Println(string(data))
	}
	{
		typ := AppPostInfo1{}
		data, _ := json.Marshal(typ)
		fmt.Println(string(data))
	}
}

```



