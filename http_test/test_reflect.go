package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var n int
	n = 123
	t := interface{}(n)
	rt := reflect.TypeOf(t)

	fmt.Println("Align:", rt.Align(), "FieldAlign:", rt.FieldAlign(), "NumMethod:", rt.NumMethod(),
		"name:", rt.Name(), "PkgPath:", rt.PkgPath(), "String:", rt.String())
}
