package babel

import (
	babel "github.com/jvatic/goja-babel"
	"io/ioutil"
	"strings"
)

func ReadFile(path string) (bytes []byte, err error)   {
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
}
