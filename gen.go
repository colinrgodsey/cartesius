package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type rep map[string]string

var genList = []struct {
	from, to string
	replace  rep
}{
	{"f64/vecN", "f32/vecN", rep{"f64": "f32", "float64": "float32"}},
	{"f64/base", "f32/base", rep{"f64": "f32", "float64": "float32"}},

	{"f64/vecN", "f64/vec2", rep{"VecN": "Vec2"}},
	{"f64/vecN", "f64/vec3", rep{"VecN": "Vec3"}},
	{"f64/vecN", "f64/vec4", rep{"VecN": "Vec4"}},

	{"f32/vecN", "f32/vec2", rep{"VecN": "Vec2"}},
	{"f32/vecN", "f32/vec3", rep{"VecN": "Vec3"}},
	{"f32/vecN", "f32/vec4", rep{"VecN": "Vec4"}},
}

/* Generates the different vector classes */
func main() {
	for _, gen := range genList {
		str := readFile(gen.from)
		for from, to := range gen.replace {
			str = strings.ReplaceAll(str, from, to)
		}
		str = fmt.Sprintf("/* FILE WAS AUTO-GENERATED FROM %v */\n\n%v", gen.from, str)
		writeFile(gen.to, str)
		fmt.Println("generated " + gen.to)
	}
}

func readFile(path string) string {
	path = path + ".go"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func writeFile(path string, str string) {
	path = path + ".go"
	if err := ioutil.WriteFile(path, []byte(str), 0664); err != nil {
		panic(err)
	}
}
