package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := make(map[string]interface{})

	a["sample"] = "sample"
	a["sample2"] = "sample2"

	fmt.Printf("a: %v\n", a)
	b := a
	fmt.Printf("b: %v\n", b)
	delete(b, "sample")
	fmt.Printf("b: %v\n", b)


	a = make(map[string]interface{})

	a["sample"] = map[string]interface{}{"deepSample": 1}
	a["sample2"] = "sample2"

	fmt.Printf("a: %v\n", a)
	trimIgnoreParamForBody(a, []string{"deepSample"})
	fmt.Printf("a: %v\n", a)


}

func trimIgnoreParamForBody(body map[string]interface{}, ignoreKeys []string) {

	for k, v := range body {
		for _, key := range ignoreKeys {
			if key == k {
				delete(body, key)
			}
		}
		if reflect.TypeOf(v).Kind() == reflect.Map {
			trimIgnoreParamForBody(v.(map[string]interface{}), ignoreKeys)
		}
	}
}