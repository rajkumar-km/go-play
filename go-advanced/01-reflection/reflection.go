/*
Package demonostrates the use of reflection in Go

Reflection is a built-in mechanism in Go to perform the certain operations without knowing
their actual type at compile time.
  - Know their type and inspect the values at run time - For example, fmt.Println does not
    know the argument types at compile time, but inspects and prints any value at runtime.
  - Updating variables - For example, a json decoder library accepts any type and populates its
    structure at runtime from the input json.
  - To call their methods
  - To apply the operations intrinsic to their representation

Reflection also treats the types as like first-class values.

Why Reflection?

 1. To perform certain operations uniformly on various types that don't satisfy a common interface.
    Example: fmt.Println() prints value of all types. It works with the user defined type winthout
    knowing its representation. Also, a good fit for any new types that will be defined in future.

 2. To perform operation on a type without knowing the actual representation
    Example: json marshal/unmarshal that works for any structs.

 3. To perform operation on a future/unknown data structure that don't exist today.

reflect.Type

  - reflect.Type is a built-in interface that represents a Go type

  - It has several methods to identify the type, and inspecting the properties of the type such as
    fields in a struct type, and parameters used in func type.

  - reflect.TypeOf() accepts any object with interface{} and returns dynamic reflect.Type. But it
    always returns the concrete type while passing an interface variable.

  - reflect.Type.NumField() returns the number of fields for the struct type

  - reflect.Type.Field(i) returns the information about struct field such as Name and Tag.

reflect.Value

  - reflect.Value if a built-in struct that can hold a value of any Go type.

  - It is similar to an empty interface{} that can hold any values, but an empty interface{} hides
    everything about the value and does not expose outside.

  - In contrast, reflect.Value allow us to inspect its contents regardless of its type, and
    provides several other methods.

  - The reflect.ValueOf() function accepts any interface{} and returns a reflect.Value
    containing the interface’s dynamic value.

  - reflect.Value.Type() returns the associated reflect.Type

  - reflect.Value.Interface() returns the wrapped value back in an interface{}

reflect.Kind

  - reflect.Kind is an inbuilt enum consisting of core types such as bool, int, int8, ...,
    uint, uint8, ..., float32, float64, complex64, complex128, array, chan, func, interface,
    map, pointer, slice, string, struct, and unsafepointer.

  - reflect.Value.Kind() returns reflect.Kind which is an enum consisting of core types

Inspecting reflect.Value at run time

  - Note: Although reflect.Value has several methods, it is unsafe to call them for all the kinds.
    Some of methods are applicable to specific kinds. Calling them with other kinds causes panic.

  - Methods for reflect.Array or reflect.Slice kind.

  - reflect.Value.Len() returns length of array/slice

  - reflect.Value.Index(i) returns the reflect.Value at index i

  - Methods for reflect.Map kind

  - reflect.Value.MapKeys() returns the keys list as []reflect.Value

  - reflect.Value.MapIndex(reflect.Value) returns reflect.Value for the key part of MapKeys()

  - Methods for reflect.Struct kind

  - reflect.Value.NumField() - returns number of fields in the struct

  - reflect.Value.Field(i) - returns reflect.Value for the field

  - reflect.Value.FieldByName(name) - returns reflect.Value for the specific field name

  - Methods for reflect.Pointer kind

  - reflect.Value.IsNil() returns true if the value is nil, false otherwise.

  - reflect.Value.Elem() returns the actual reflect.Value referenced by pointer

Setting variables with reflect.Value

  - In Go, a variable is an addressable storage location that contains a value. It can be updated
    with their addresses. Some of the reflect.Value are addressable and some or not.

  - None of the value returned by reflect.ValueOf() is addressable. Because, reflect.ValueOf()
    copies the value. Even for pointer variable it takes a copy, but just copies the address.

  - reflect.Value.Elem() which deference the pointer actually contains the addressable variable.

  - reflect.Value.CanAddr() is anyway useful to determine if a value is addressable.

  - reflect.Value.CanSet() returns if the value is both addressable and settable. Go does not
    allow to set values on unexported field although it allows to read.

  - Use reflect.Value.Set* methods to set the values to variables. This performs all checks that
    the compiler performs when setting a variable. It can cause panic when setting a wrong type.

  - reflect.Value.Set(reflect.ValueOf(10)) // set int or interface{}

  - reflect.Value.Set(reflect.ValueOf(int64(10))) // set int64 or interface{}

  - reflect.Value.SetInt(10) // set signed int, int8, int16, int32, int64

  - reflect.Value.Set(reflect.Zero(v.Type())) // set to corresponding null type

  - Note: Methods like setInt(), setString() should not be used for interface{}

  - reflect.New(reflect.Type) - returns a new zero value for the type

  - reflect.Append(s reflect.Value, t ...reflect.Value) - append element to the slice and returns
    the resulting slice

  - reflect.MakeMap(reflect.Type) - creates a new map with the specified type

  - reflect.Value.SetMapIndex(key, value) - set key-value pair to map

Caution:

 1. For every mistake, errors are reported at run time as panic by passing compilers type checking.
    - To avoid:
    - Encapsulate the use of reflection in a separate package
    - If possible, avoid the use of reflect.Value in your packages API args.
    - Otherwise, perform additional dynamic checks for each operation. For example, fmt.Printf
    does not cause panic when using wrong types (say %d for a string), but still prints it as
    message like "%!d(string=hello)"
    - Reflection also reduces the safety and accuracy of automatic refactoring and analysis tools
    since it does not have type information.

 2. Hard to understand the reflection code. Since the type serve as documentation for regular
    functions
    - Always carefully document the expected types for an interface{} or reflect.Value arguments.

 3. Reflection may be slower than the functions specialized for specific types.
    - Avoid reflection for the functions in critical path
    - Testing isapar tic ularly good fit for reflec tion since most tests use small data sets.
*/
package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

// DemoReflection demonstrates the key types used in reflection: reflect.Type, reflect.Value, and
// reflect.Kind
func DemoReflection() {
	/* reflect.Type */
	fmt.Println("1. reflect.Type holds the underlying type")
	var t reflect.Type = reflect.TypeOf(10)
	// reflect.Type is compatiple with fmt.Stringer interface
	fmt.Println("reflect.TypeOf(10) =", t)                   // int
	fmt.Println("reflect.TypeOf(10).String() =", t.String()) // int

	// reflect.TypeOf() always returns the concrete type associated with the interface
	fmt.Println("reflect.TypeOf() always returns the concrete type for interfaces")
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // *os.File
	fmt.Printf("%T\n", w)          // *os.File (note that %T internally uses reflect.TypeOf())

	/* reflect.Value */
	fmt.Println("\n2. reflect.Value holds the underlying value")
	v := reflect.ValueOf(10)
	fmt.Println("reflect.ValueOf(10) =", v)       // 10
	fmt.Printf("%%v for reflect.Value = %v\n", v) // 10
	// reflect.Value just prints the type with fmt.Stringer
	fmt.Println(".String() for reflect.Value is different =", v.String()) // <int Value>

	// reflect.Value.Interface() - Wrap the reflect.Value as interface{}
	myData := interface{}("hello")
	v = reflect.ValueOf(myData)
	myDataIsBack := v.Interface() // returns back the the interface{}(value)
	fmt.Println("Wrap reflect.Value with .Interface() =", reflect.TypeOf(myDataIsBack))

	/* reflect.Kind */
	fmt.Println("\n3. reflect.Kind is an enum for core types")
	fmt.Println(reflect.ValueOf(10.2).Kind())       // float64
	fmt.Println(reflect.ValueOf(struct{}{}).Kind()) // struct
	fmt.Println(reflect.ValueOf(w).Kind())          // ptr
}

// DemoSetValue demonstrates setting values to variables through reflection
func DemoSetValue() {
	// Note: Only addressable values can be set, otherwise it causes panic

	// pointer to a variable is addressable
	myVal := 100
	fmt.Println("myVal before set =", myVal)
	reflect.ValueOf(&myVal).Elem().SetInt(200)
	fmt.Println("myVal after SetInt(200) =", myVal)

	// pointer to an array index is addressable
	arr := [1]uint{100}
	fmt.Println("arr[0] before set =", arr[0])
	reflect.ValueOf(&arr[0]).Elem().SetUint(200)
	fmt.Println("arr[0] after SetUint(200) =", arr[0])

	slice := []uint{100, 200}
	fmt.Println("slice[0] before set =", slice[0])
	reflect.ValueOf(&slice[0]).Elem().SetUint(300)
	fmt.Println("slice[0] after SetUint(300) =", slice[0])

	// pointer to an interface variable is addressable
	// But, remeber that the concrete value stroed inside the interface is not addressable
	intf := interface{}(100)
	fmt.Println("interface before set =", intf)
	reflect.ValueOf(&intf).Elem().Set(reflect.ValueOf("hello"))
	fmt.Println("interface after Set(reflect.ValueOf(\"hello\")) =", intf)

	// index of a map[key] is not addressable, because the map is implemented using hash
	// The hash changes when the keys are inserted/deleted and memory address can change
	// myMap := make(map[string]string)
	// myMap["hello"] = "world"
	// reflect.ValueOf(&myMap["hello"])
}
