package main

import "fmt"

func main() {
	var data *byte
	var in interface{}

	fmt.Println(data, data == nil)    // <nil> true
	fmt.Println(in, in == nil)    // <nil> true

	in = data
	fmt.Println(in, in == nil)    // <nil> false    // data 值为 nil，但 in 值不为 nil
}