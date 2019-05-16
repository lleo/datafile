package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/ugorji/go/codec"
	"github.com/vmihailenco/msgpack"
)

var mh codec.MsgpackHandle

func init() {
	mh.MapType = reflect.TypeOf(map[string]interface{}(nil))
}

type mystruct struct {
	Integer int
	String  string
	Boolean bool
}

func main() {
	fmt.Println("map-to-map")

	var s = &mystruct{42, "hello", true}
	var in = struct2map(s)
	//in["Integer"] = int(42)
	//in["String"] = "hello"
	//in["Boolean"] = true

	fmt.Println(in)

	b, err := msgpack.Marshal(in)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("len(b)=%d\n", len(b))

	var out = make(map[string]interface{})

	err = msgpack.Unmarshal(b, &out)
	if err != nil {
		log.Print(err)
	}

	fmt.Println("out =", out)

	//dec := codec.NewDecoderBytes(b, &mh)
	//err = dec.Decode(out)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("out=%v\n", out)
	//t := reflect.TypeOf(out["str"])
	//strType := reflect.TypeOf("")
	//fmt.Printf("reflect.TypeOf(out[\"str\"]) = %s\n", t)
	//fmt.Printf("%v is convertable to %v => %t\n", t, strType, t.ConvertibleTo(strType))
	//fmt.Printf("out[\"str\"] = %q\n", string(out["str"].([]uint8)))
	//fmt.Printf("out[\"str\"] = %q\n", string(out["str"].([]byte)))
	//fmt.Printf("reflect.TypeOf(out) = %s\n", reflect.TypeOf(out))
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
		//fmt.Printf("field[%d] = %q\n", i, t.Field(i).Name)
		name := t.Field(i).Name
		value := v.FieldByName(name).Interface()
		m[name] = value
	}

	return m
}
