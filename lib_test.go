// lib_test.go
// Copyright (C) 2021 Kasai Koji

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package map_mapper_test

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/streamwest-1629/map_mapper"
)

func TestExample(t *testing.T) {
	Example()
}

func Example() {
	type B struct {
		A int64    `map_mapper:"IntA"`
		B int64    `map_mapper:"IntB?"`
		C int64    `map_mapper:"IntC"`
		D int64    `map_mapper:"IntD?"`
		E bool     `map_mapper:"Bool"`
		F float64  `map_mapper:"Float"`
		G string   `map_mapper:"String"`
		H []string `map_mapper:"StringsH"`
		I []string `map_mapper:"StringsI"`
	}
	bytesA, _ := json.Marshal(B{})

	fmt.Println("un-initialized: ")
	fmt.Println(string(bytesA))
	// un-initialized:
	// {"A":0,"B":0,"C":0,"D":0,"E":false,"F":0,"G":"","H":null,"I":null}

	if compiled, err := map_mapper.CompileStruct(B{}); err != nil {
		log.Fatal(err.Error())
	} else {

		if structure, err := compiled.Unmarshal(map[interface{}]interface{}{
			"IntA":     "1",
			"IntB":     int64(2),
			"IntC":     int64(3),
			"Bool":     true,
			"Float":    "0.20210822",
			"String":   "Hello, World!",
			"StringsH": []string{"Happy", "New", "Year!"},
			"StringsI": []interface{}{"Happy", "Foo", "Bar"},
		}, ""); err != nil {
			log.Fatal(err.Error())

		} else {
			bytesB, _ := json.Marshal(structure)
			fmt.Println("map-mappered: ")
			fmt.Println(string(bytesB))
			// map-mappered:
			// {"A":1,"B":0,"C":3,"D":0,"E":true,"F":0.20210822,"G":"Hello, World!","H":["Happy","New","Year!"],"I":["Happy","Foo","Bar"]}

			fmt.Println(reflect.TypeOf(structure).String())
			// map_mapper_test.B
		}
	}
}
