package main

import (
	"fmt"
	"strings"

	uuid "github.com/github.com/satori/go.uuid"
)

func RandStr() string {
	uuidO, _ := uuid.NewV4()
	uid := uuidO.String()
	s := strings.Split(uid, "-")
	n := len(s)
	return s[n-1]
}
func main() {
	fmt.Println(RandStr())
}
