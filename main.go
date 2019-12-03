package main

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/duanqy/tender/pkg/babel"
	"log"
)

func main()  {
	registry := require.NewRegistryWithLoader(babel.ReadFile) // this can be shared by multiple runtimes
	src,err := babel.ReadFile("tender.js")
	if err != nil {
		log.Fatalln(err)
	}

	vm := goja.New()
	_ = registry.Enable(vm)

	vm.Set("Tender", map[string]interface{}{
		"server": func(x string,fn func(map[string]interface{})) {
			fmt.Println("x",x)
			var cfg = make(map[string]interface{})
			fn(cfg)

			fmt.Println(cfg)
		},
		"name" :"test",
	})
	vm.Set("console", map[string]interface{}{
		"log" : func(data ...interface{}) {
			fmt.Println(data...)
		},
	})

	v,err := vm.RunString(string(src))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(v)
}
