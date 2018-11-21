package main

import (
	"fmt"
	"reflect"
)

func main() {
	type S struct {
		F string `species:"gopher" color:"blue"`
		a int    `a:"name of a"`
	}

	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))

	field2 := st.Field(1)
	fmt.Println(field2.Name)
	fmt.Println(field2.Index)
	fmt.Println(field2.Tag)
	fmt.Println(field2.Type)

}
