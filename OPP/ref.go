package main

import (
	"fmt"
	"reflect"
)

func tpy() {

	typeI := reflect.TypeOf(1)
	typeS := reflect.TypeOf("hello")
	fmt.Println(typeI)
	fmt.Println(typeS)
}

func get_field() {

	typeUser := reflect.TypeOf(User{})
	fieldNum := typeUser.NumField()
	for i := 0; i < fieldNum; i++ {
		field := typeUser.Field(i)
		fmt.Printf("num:%d--- Name:%s---  an:%t--- off:%d--- type:%s--- Isexported:%t--- json:%s--- place:%s--- xml:%s\n",
			i,
			field.Name,
			field.Anonymous,
			field.Offset,
			field.Type,
			field.IsExported(),
			field.Tag.Get("json"),
			field.Tag.Get("place"),
			field.Tag.Get("xml"))
	}

	if nameField, ok := typeUser.FieldByName("Name"); ok {
		fmt.Println(nameField.IsExported())
	}

}

func memoly_aligh() {
	type A struct {
		sex    bool   // 1B offset 0
		weight uint16 // 2B 需要在2的倍数上进行填充偏移，所以offset为2
		addr   byte   // 1B off 4
		age    int64  // 8B off 8
	}
	t := reflect.TypeOf(A{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%d---  Name:%s Offset:%d \n", i, field.Name, field.Offset)
	}

}

func main() {
	//get_field()
	memoly_aligh()
}

/*
FieldName	FieldType	FieldSize	Offset
sex			bool		1 byte		0
padding		uint8		1 byte		1
weight		uint16		2 bytes		2
addr		byte		1 byte		4
padding		[3]uint8	3 bytes		5
age			int64		8 bytes		8
*/
