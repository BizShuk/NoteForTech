package main

import (
	"fmt"
	"reflect"
)

func main() {
	tryTypeOf()
	tryValueOf()
	tryValueOfInstance()
}

func tryTypeOf() {
	var t func(a string)

	t = func(a string) {
		fmt.Println("test")
	}

	var i interface{}

	i = t

	tp := reflect.TypeOf(i)
	fmt.Println("TypeOf:", tp)
	fmt.Println("TypeOf.Kind():", tp.Kind())
}

func tryValueOf() {
	var t func(a string)

	t = func(a string) {
		fmt.Println("test")
	}

	var i interface{}

	i = t

	tp := reflect.ValueOf(i)
	fmt.Println("TypeOf:", tp)
	fmt.Println("TypeOf.Kind():", tp.Kind())
}

func tryValueOfInstance() {
	var t string = "original"
	var i interface{}

	i = t

	tp := reflect.ValueOf(&i).Elem()
	t = "new"
	fmt.Println(i, t, tp)
	// TODO: how to set on reflect ValueOf
}
