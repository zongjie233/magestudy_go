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

	// 获取user中嵌套的Car结构体其中成员变量
	speedField := typeUser.FieldByIndex([]int{3, 0})
	fmt.Println(speedField.Name, speedField.IsExported())
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

func getMethod() {
	typeUser := reflect.TypeOf(Bird{})
	methodNum := typeUser.NumMethod()
	for i := 0; i < methodNum; i++ {
		method := typeUser.Method(i) // 只能获取可导出的方法
		fmt.Printf("name:%s type:%s isEx %v\n", method.Name, method.Type, method.IsExported())
	}
}

func getStructrelation() {
	plan := Plane{}
	bird := Bird{}
	frog := Frog{}
	planeType := reflect.TypeOf(plan)
	birdType := reflect.TypeOf(bird)
	fmt.Println(planeType)
	fmt.Println(birdType)
	frogType := reflect.TypeOf(frog).Elem()
	// 查看结构体是否实现指定接口
	fmt.Println(frogType.Implements(frogType))

}

func valueOtherconversion() {
	iValue := reflect.ValueOf(1)
	sValue := reflect.ValueOf("hello")
	userValue := reflect.ValueOf(&User{
		Name: "hs",
		Sex:  1,
		Age:  19,
	})
	fmt.Println(iValue)
	fmt.Println(sValue)
	fmt.Println(userValue)

	iType := iValue.Type()
	sType := sValue.Type()
	userType := userValue.Type()

	fmt.Println(iValue.Kind() == iType.Kind(), iValue.Kind() == reflect.Int)
	fmt.Println(sValue.Kind() == sType.Kind(), sValue.Kind() == reflect.String)
	fmt.Println(userValue.Kind() == userType.Kind(), userValue.Kind() == reflect.Ptr)

	userValue2 := userValue.Elem() // 与Addr是互逆操作，Elem是在解析指针
	fmt.Println(userValue2.Kind() == reflect.Struct)
	userValue3 := userValue2.Addr()
	fmt.Println(userValue3.Kind() == reflect.Ptr)
}

func main() {
	//get_field()
	//memoly_aligh()
	//getMethod()
	//getStructrelation()
	valueOtherconversion()
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
