package main

import (
	"fmt"
	"reflect"
)

/**
基础知识：
1.获取值使用 ValueOf
2.获取元素使用 Elem
3.判断类型使用 Kind 或者 Type，如果类型使用 type 定义，Kind 返回的是底层类型
4.设置值使用 Set
5.如果想修改值，必须为指针类型
6.如果想要获取结构体信息或者类型，可以使用 TypeOf
*/

type CustomInt int

type User struct {
	Name interface{} `json:"name"`
	Age  CustomInt   `json:"age"`
}

func main() {
	user := User{Name: "张三", Age: 20}
	// 获取的是指针地址的反射对象
	uPtr := reflect.ValueOf(&user)
	fmt.Println(uPtr.Type())
	// Elem 获取地址（指针）所指向的元素
	u := uPtr.Elem()
	fmt.Println(u.Type())
	// 获取字段数量
	numField := u.NumField()
	fmt.Println(numField)
	// 获取字段
	name := u.FieldByName("Name")
	age := u.FieldByName("Age")
	fmt.Println(name, age)
	// 判断字段类型
	if name.Kind() == reflect.Interface {
		// name是interface类型，因此还需要 Elem 才能获取值的反射对象
		if name.Elem().Kind() == reflect.String {
			name.Set(reflect.ValueOf("李四"))
		}
		if name.Elem().Kind() == reflect.Int {
			name.Set(reflect.ValueOf(88))
		}
	}
	fmt.Println(user)
	fmt.Println("------------------")
	// 获取字段相关的信息:名字、类型、tag
	uType := reflect.TypeOf(user)
	for i := 0; i < uType.NumField(); i++ {
		field := uType.Field(i)
		fmt.Println("名字", field.Name)
		fmt.Println("类型", field.Type)
		fmt.Println("底层类型", field.Type.Kind())
		fmt.Println("tag值", field.Tag)
		fmt.Println("json-tag的值", field.Tag.Get("json"))
		fmt.Println("----------------")
	}
}
