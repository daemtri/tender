package main

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	babel "github.com/jvatic/goja-babel"
	"io/ioutil"
	"log"
	"strings"
)

func main()  {
	registry := require.NewRegistryWithLoader(func(path string) (bytes []byte, err error) {
		if !strings.HasSuffix(path, ".js") {
			path = path + ".js"
		}
		src,err := ioutil.ReadFile(path)
		if err != nil {
			return nil,err
		}
		resp,err := babel.TransformString(string(src), map[string]interface{}{
			"presets": []string{
				"es2015",
			},
		})
		if err != nil {
			return nil,err
		}
		return []byte(resp),nil

	}) // this can be shared by multiple runtimes

	src,err := ioutil.ReadFile("tender.js")
	if err != nil {
		panic(err)
	}

	resp,err := babel.TransformString(string(src), map[string]interface{}{
		"presets": []string{
			"es2015",
		},
	})
	if err != nil {
		panic(err)
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
	})
	ioutil.WriteFile("tender.es5.js", []byte(resp), 0755)
	v,err := vm.RunString(resp)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(v)
}
