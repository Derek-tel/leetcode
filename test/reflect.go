package main

import (
	"fmt"
	"reflect"
)

type User struct {
}

func (u *User) Put(i int) {
	fmt.Println(i)
}

func main() {
	user := User{}
	value := reflect.ValueOf(&user)
	f := value.MethodByName("Put")
	f.Call([]reflect.Value{reflect.ValueOf(1)})

	tem := User{}
	val := reflect.ValueOf(&tem)
	fc := val.MethodByName("put")
	fc.Call([]reflect.Value{reflect.ValueOf(1)})

}
