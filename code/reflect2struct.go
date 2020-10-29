package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	mList := []map[string]interface{}{
		{"Id": 213, "Name": "zhaoliu", "Sex": "男"},
		{"Id": 56, "Name": "zhangsan", "Sex": "男"},
		{"Id": 7, "Name": "lisi", "Sex": "女"},
		{"Id": 978, "Name": "wangwu", "Sex": "男"},
	}

	type User struct {
		Id   int
		Name string
		Sex  string
	}
	users := []*User{}

	mapToStruct(mList, &users)
	fmt.Printf("users: %+v\n", users)
}

func mapToStruct(mList []map[string]interface{}, model interface{}) (err error) {
	val := reflect.Indirect(reflect.ValueOf(model))
	typ := val.Type()
	for _, r := range mList {
		mVal := reflect.Indirect(reflect.New(typ.Elem().Elem())).Addr()
		for key, val := range r {
			err = setField(mVal.Interface(), key, val)
			if err != nil {
				return err
			}
		}
		val = reflect.Append(val, mVal)
	}
	DeepCopy(model, val.Interface())
	return err
}

//用map的值替换结构的值
func setField(obj interface{}, name string, value interface{}) error {
	// 将首字母转换为大写
	sl := strings.Split(name, "")
	sl[0] = strings.ToUpper(sl[0])
	name = strings.Join(sl, "")

	structValue := reflect.ValueOf(obj).Elem() //结构体属性值
	//fmt.Printf("structValue: %+v\n", structValue)
	structFieldValue := structValue.FieldByName(name) //结构体单个属性值
	//fmt.Printf("structFieldValue: %+v\n", structFieldValue)
	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type() //结构体的类型
	val := reflect.ValueOf(value)              //map值的反射值

	if structFieldType != val.Type() {
		return errors.New("type is err")
	}

	structFieldValue.Set(val)
	return nil
}

func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
