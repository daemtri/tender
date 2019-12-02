package main

import (
	"fmt"
	"github.com/dop251/goja"
	"io/ioutil"
	"log"
)

func main()  {
	src,err := ioutil.ReadFile("tender.js")
	if err != nil {
		panic(err)
	}

	vm := goja.New()
	vm.Set("Tender", map[string]interface{}{
		"server": func(x string,fn func(map[string]interface{})) {
			fmt.Println("x",x)
			var cfg = make(map[string]interface{})
			fn(cfg)
			fmt.Println(cfg)
		},
	})
	v,err := vm.RunString(string(src))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(v)
}