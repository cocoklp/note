package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Sslmeta struct {
	Ssl  string
	User string
}

func main() {
	mList2 := map[string]interface{}{
		"Ssl":  "klp1",
		"User": "klpklp1",
	}
	var ssls *Sslmeta
	mapToStruct(mList2, &ssls)
}

func mapToStruct(mList map[string]interface{}, model interface{}) (err error) {
	val := reflect.Indirect(reflect.ValueOf(model))
	valof := reflect.ValueOf(model)

	fmt.Println(valof.Type(), reflect.ValueOf(model), val.Type(), reflect.ValueOf(val))

	typ := val.Type()

	mVal := reflect.Indirect(reflect.New(typ.Elem().Elem())).Addr()
	for key, val := range mList {
		err = setField(mVal.Interface(), key, val)
		if err != nil {
			return err
		}
	}
	val = reflect.Append(val, mVal)

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
	fmt.Println("fieldtype", structFieldType, val)
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
