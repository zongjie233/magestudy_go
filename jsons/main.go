package main

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type User struct {
	Name string
	Age  int
	Sex  byte `json:"gender"`
}

func Marshal(v interface{}) ([]byte, error) {
	value := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)
	bf := bytes.Buffer{}
	switch typ.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Uint16, reflect.Uint32, reflect.Uint8:
		return []byte(fmt.Sprintf("%v", value.Interface())), nil
	case reflect.String:
		return []byte(value.String()), nil
	case reflect.Bool:
		return []byte(fmt.Sprintf("%t", value.Bool())), nil
	case reflect.Float32, reflect.Float64:
		return []byte(fmt.Sprintf("%f", value.Float())), nil
	case reflect.Struct: // 结构体序列化
		bf.WriteByte('{')
		if value.NumField() > 0 {
			for i := 0; i < value.NumField(); i++ {
				fieldValue := value.Field(i)
				fieldType := typ.Field(i)
				if fieldType.IsExported() {
					name := fieldType.Name
					if len(fieldType.Tag.Get("json")) > 0 {
						name = fieldType.Tag.Get("json")
					}
					bf.WriteByte('"')
					bf.WriteString(name)
					bf.WriteByte('"')
					bf.WriteByte(':')
					if bs, err := Marshal(fieldValue.Interface()); err == nil {
						bf.Write(bs)
					} else {
						return nil, err
					}
					bf.WriteByte(',')
				}
			}
			bf.Truncate(len(bf.Bytes()) - 1)
		}
		bf.WriteByte('}')
		return bf.Bytes(), nil
	default:
		return nil, fmt.Errorf("暂时不支持这种数据类型")
	}

}

// 反序列化
func UnMarshal(v interface{}, data []byte) error {
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
	case reflect.Uint, reflect.Uint8, reflect.Uint32:
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
	case reflect.Struct:
		if s[0] == '{' && s[len(s)-1] == '}' {
			arr := strings.Split(s[1:len(s)-1], ",")
			if len(arr) > 0 {
				fieldCount := typ.NumField()
				tag2Fidle := make(map[string]string, fieldCount)
				for i := 0; i < fieldCount; i++ {
					fieldType := typ.Field(i)
					name := fieldType.Name
					if len(fieldType.Tag.Get("json")) > 0 {
						name = fieldType.Tag.Get("json")
					}
					tag2Fidle[name] = fieldType.Name
				}
				for _, ele := range arr {
					brr := strings.SplitN(ele, ":", 2) // 声明只要两个部分，出去第一部分，其他均归到第二部分
					tag := brr[0]
					if tag[0] == '"' && tag[len(tag)-1] == '"' {
						tag = tag[1 : len(tag)-1] //gender
						if fieldName, exists := tag2Fidle[tag]; exists {
							fieldValue := value.FieldByName(fieldName)
							fieldType := fieldValue.Type()
							if fieldValue.Kind() != reflect.Ptr {
								fieldValue = fieldValue.Addr()
								if err := UnMarshal(fieldValue.Interface(), []byte(brr[1])); err != nil {
									return err
								}
							} else {
								newValue := reflect.New(fieldType.Elem()) // 反射中创建一个对象 等价于 user:=new(User)
								if err := UnMarshal(fieldValue.Interface(), []byte(brr[1])); err != nil {
									return err
								} else {
									//value.FieldByName(fieldName).Set(newValue)
									fieldValue.Set(newValue)
								}
							}
						} else {

						}
					}
				}
			}
		} else {
			return fmt.Errorf("json格式不对:%s", s)
		}
	default:
		return fmt.Errorf("暂时不支持这种数据类型")
	}
	return nil
}

func main() {
	//var v interface{}
	//v = 7.5
	//if data, err := Marshal(v); err == nil {
	//	fmt.Println(string(data))
	//	var a float32
	//	if e := unMarshal(&a, data); e == nil {
	//		fmt.Println(a)
	//	} else {
	//		fmt.Println(e)
	//	}
	//}
	user := User{
		Name: "hs",
		Age:  18,
		Sex:  1,
	}
	if data, err := Marshal(user); err == nil {
		fmt.Println(string(data))
		var u2 User
		if err := UnMarshal(&u2, data); err == nil { // UnMarshal第一个参数必须接受指针
			fmt.Println(u2.Name, u2.Age, u2.Sex)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}
