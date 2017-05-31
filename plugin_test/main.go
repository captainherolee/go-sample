package main

import "fmt"
import "plugin"

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. ./lib/libcal.a
// #include "./include/cal.h"
import "C"

func main() {
  	fmt.Println("StartMain")
  	fmt.Println("C library ", C.add(10,20) )

  	p, err := plugin.Open("./lib/helloplugin.so")
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("Helloworld")
	if err != nil {
		panic(err)
	}

	f.(func())()

	f1, err := p.Lookup("Add")
	if err != nil {
		panic(err)
	}

	sum := f1.(func(int, int) int)(1, 2)

	fmt.Println(sum)
}


