package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func Marshal(v interface{}) ([]byte, error) {
	value := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)
	switch typ.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Uint16, reflect.Uint32:
		return []byte(fmt.Sprintf("%v", value.Interface())), nil
	case reflect.String:
		return []byte(value.String()), nil
	case reflect.Bool:
		return []byte(fmt.Sprintf("%t", value.Bool())), nil
	case reflect.Float32, reflect.Float64:
		return []byte(fmt.Sprintf("%f", value.Float())), nil
	default:
		return nil, fmt.Errorf("暂时不支持这种数据类型")
	}

}

// 反序列化
func unMarshal(v interface{}, data []byte) error {
	value := reflect.ValueOf(v)
	typ := value.Type() // 等价于 typ := reflect.TypeOf(v)
	if typ.Kind() != reflect.Ptr {
		return errors.New("v must be pointer")
	}
	typ = typ.Elem() //解析指针，变成非指针
	value = value.Elem()

	s := string(data)
	switch typ.Kind() {
	case reflect.Int:
		if i, err := strconv.ParseInt(s, 10, 64); err == nil {
			value.SetInt(i)
		} else {
			return err
		}
	case reflect.Uint:
		if i, err := strconv.ParseUint(s, 10, 64); err == nil {
			value.SetUint(i)
		} else {
			return err
		}
	case reflect.Bool:
		if b, err := strconv.ParseBool(s); err == nil {
			value.SetBool(b)
		} else {
			return err
		}
	case reflect.Float32:
		if f, err := strconv.ParseFloat(s, 32); err == nil {
			value.SetFloat(f)
		} else {
			return err
		}
	case reflect.String:
		value.SetString(s)
	default:
		return fmt.Errorf("暂时不支持这种数据类型")
	}
	return nil
}

func main() {
	var v interface{}
	v = 7
	data, err := Marshal(v)
	if err == nil {
		fmt.Println(string(data))
		var a int
		unMarshal(&a, data)
		fmt.Println(a)
	}
}
