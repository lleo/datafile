package main

import (
	"fmt"
	"log"
	"reflect"
)

type mystruct struct {
	Integer int
	String  string
	Boolean bool
}

func main() {
	fmt.Println("struct-to-map")
	s := &mystruct{42, "hello", true}
	var m = struct2map(s)
	fmt.Printf("m = %v\n", m)
}

func struct2map(val interface{}) map[string]interface{} {
	p := reflect.ValueOf(val)
	if p.Type().Kind() != reflect.Ptr {
		log.Fatal("struct2map val not a Ptr")
	}
	v := p.Elem() //get the Value p points to
	t := v.Type()
	if t.Kind() != reflect.Struct {
		log.Fatal("struct2map val is not a Ptr to a struct")
	}
	//log.Println("struct type:", t.Name())
	//log.Println("struct type package path:", t.PkgPath())

	if t.NumField() != v.NumField() {
		log.Fatal("type NumField != value NumField ?!?")
	}

	m := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("field[%d] = %q\n", i, t.Field(i).Name)
		name := t.Field(i).Name
		value := v.FieldByName(name).Interface()
		m[name] = value
	}

	return m
}
