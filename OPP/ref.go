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

func valueIsempty() {
	var ifc interface{}
	v := reflect.ValueOf(ifc)
	fmt.Printf("v持有真实的值 %t\n", v.IsValid())
	ifc = 8
	v = reflect.ValueOf(ifc)
	fmt.Printf("v持有真实的值 %t\n", v.IsValid())

	var user *User = nil
	v = reflect.ValueOf(user)
	if v.IsValid() {
		fmt.Printf("v持有值为nil %t\n", v.IsNil()) // 调用IsNil之前，必须保证isValid()true，否则panic
	} else {
		fmt.Printf("v没有持有值\n")
	}
}

// 可寻址
func addressable() {
	v1 := reflect.ValueOf(1)
	var x int
	v2 := reflect.ValueOf(x)
	v3 := reflect.ValueOf(&x)
	v4 := v3.Elem()
	fmt.Printf("v1 is addressable %t\n", v1.CanAddr())
	fmt.Printf("v2 is addressable %t\n", v2.CanAddr())
	fmt.Printf("v3 is addressable %t\n", v3.CanAddr())
	fmt.Printf("v4 is addressable %t\n", v4.CanAddr())
}

func changeValue() {
	// 通过反射修改Int
	var i int = 10
	fmt.Printf("address of i is %p \n", &i)
	iValue := reflect.ValueOf(&i)
	if iValue.CanAddr() {
		iValue.SetInt(8)
		fmt.Printf("1: i=%d\n", i)
	}
	iValue2 := iValue.Elem()
	if iValue2.CanAddr() {
		iValue2.SetInt(8)
		fmt.Printf("2: i=%d %p\n", i, &i)
	}

	// 通过反射修改String
	var s string = "hello"
	fmt.Printf("address of i is %p \n", &s)
	sValue := reflect.ValueOf(&s)
	if sValue.CanAddr() {
		sValue.SetString("world")
		fmt.Printf("3: s=%s\n", s)
	}
	sValue2 := sValue.Elem()
	if sValue2.CanAddr() {
		sValue2.SetString("world")
		fmt.Printf("4: s=%s %p\n", s, &s)
	}

	// 通过反射修改struct
	user := User{
		Name: "Zdq",
		Age:  18,
	}

	userValue := reflect.ValueOf(&user) // 必须传入&user，才为可寻址
	userValue.Elem().FieldByName("Age").SetInt(19)
	fmt.Println(user.Age)

}

// 修改切片数据
func changeSlice() {
	users := make([]*User, 3, 5)
	users[0] = &User{
		Name: "hs",
		Age:  18,
	}
	sliceValue := reflect.ValueOf(users)
	if sliceValue.Len() > 0 {
		sliceValue.Index(0).Elem().FieldByName("Age").SetInt(20)
		fmt.Println(users[0].Age)
	}

	// 切片整组元素替换
	sliceValue.Index(1).Set(reflect.ValueOf(&User{
		Name: "aaa",
		Age:  21,
	}))
	fmt.Println(users[1].Age)

}

func changeMap() {
	u1 := &User{
		Name: "aaa",
		Age:  15,
	}
	u2 := &User{
		Name: "bbb",
		Age:  16,
	}
	userMap := make(map[int]*User, 5)
	userMap[1] = u1
	mapValue := reflect.ValueOf(userMap)
	mapType := reflect.TypeOf(userMap)
	keyType := mapType.Key()
	valueType := mapType.Elem() // 获得map中value的type
	fmt.Printf("type of key %v, type of value %v\n", keyType, valueType)

	// 添加map中的一对k-v
	mapValue.SetMapIndex(reflect.ValueOf(2), reflect.ValueOf(u2))

	// 修改map中的一对k-v
	mapValue.MapIndex(reflect.ValueOf(1)).
		Elem().
		FieldByName("Age").
		SetInt(21)
	for _, v := range userMap {
		fmt.Printf("%s  %d\n", v.Name, v.Age)
	}

}

//func callFunction() {
//	valueFunc := reflect.ValueOf(Add)
//	typeFunc := valueFunc.Type()
//}

func main() {
	//get_field()
	//memoly_aligh()
	//getMethod()
	//getStructrelation()
	//valueOtherconversion()
	//valueIsempty()
	//addressable()
	//changeValue()
	//changeSlice()
	changeMap()
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
