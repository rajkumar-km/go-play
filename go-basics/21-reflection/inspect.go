package main

import (
	"fmt"
	"reflect"
)

// DemoInspect demonstrates how to inspect the value of any type without knowing
// their representation
func DemoInspect() {
	type Person struct {
		Name   string
		Age    int8
		Active bool
	}
	p1 := Person{Name: "Kumar", Age: 34, Active: true}
	inspect(reflect.ValueOf(p1))
	inspect(reflect.ValueOf(&p1))
}

// inspect prints the representation of any type. For now, it supports only struct
func inspect(v reflect.Value) {
	structName := v.Type().Name()
	fmt.Println(structName)

	// switch(data.(type)) may be used for only for limited and known types
	// But reflection is required for any types including user defined types.
	switch v.Kind() {
	case reflect.Struct:
		// Iterate the fields of struct
		for i := 0; i < v.NumField(); i++ {
			// Struct/Field name is available with reflect.Type
			fieldName := v.Type().Field(i).Name

			// Value is available with reflect.Value
			// Field(i) returns the field value as reflect.Value. This can be further checked recursively
			// to inspect nested values. But for now, just use fmt.Printf to print the value.
			var fieldVal reflect.Value = v.Field(i)
			fmt.Printf("\t%s=%v\n", fieldName, fieldVal)
		}
	case reflect.Ptr, reflect.Interface:
		// Dereference the Ptr and Interface with .Elem()
		fmt.Print("*")
		inspect(v.Elem())
	default:
		fmt.Println("unsupported type", v.Kind())
	}
}
