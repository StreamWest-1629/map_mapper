// struct.go
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

package map_mapper

import (
	"reflect"

	"github.com/streamwest-1629/textfilter"
)

const (
	Label = "map_mapper"
)

type (
	// Struct type's member
	Member struct {
		Keyname   string
		FieldId   int
		Require   bool
		Unmarshal Unmarshalizer
	}
	Struct struct {
		__type  reflect.Type
		Members []Member
	}
	StructPtr struct {
		*Struct
	}
)

var (
	labelCheck         = textfilter.RegexpExactMatches("[a-zA-Z]+[?]?")
	labelOptionalCheck = textfilter.RegexpExactMatches("[a-zA-Z]+[?]")
)

//
func CompileStructPtr(targetStructPtr interface{}) (*StructPtr, error) {

	ty := reflect.TypeOf(targetStructPtr)
	if ty.Kind() != reflect.Ptr {
		panic("target type isn't ptr")
	} else {
		if compiled, err := compileStruct(ty); err != nil {
			return nil, err
		} else {
			return &StructPtr{compiled}, nil
		}
	}
}

func CompileStruct(targetStruct interface{}) (*Struct, error) {

	ty := reflect.TypeOf(targetStruct)
	return compileStruct(ty)

}

func compileStruct(ty reflect.Type) (*Struct, error) {

	if ty.Kind() != reflect.Struct {
		panic("target type isn't struct")
	}
	members := make([]Member, 0)

	for i, l := 0, ty.NumField(); i < l; i++ {
		field := ty.Field(i)
		label := field.Tag.Get(Label)

		if err := labelCheck(label); err != nil {
			continue
		}

		member := Member{
			Keyname:   label,
			FieldId:   i,
			Require:   labelOptionalCheck(label) != nil,
			Unmarshal: GetUnmarshal(field.Type),
		}

		members = append(members, member)
	}

	return &Struct{
		__type:  ty,
		Members: members,
	}, nil

}

func (un *StructPtr) Unmarshal(mapped interface{}, property string) (dest interface{}, err error) {
	return un.unmarshal(mapped, property)
}

func (un *Struct) Unmarshal(mapped interface{}, property string) (dest interface{}, err error) {
	if ptr, err := un.unmarshal(mapped, property); err != nil {
		return nil, err
	} else {
		return reflect.ValueOf(ptr).Elem().Interface(), nil
	}
}

func (un *Struct) unmarshal(mapped interface{}, property string) (dest interface{}, err error) {
	ptr := reflect.New(un.__type)
	val := reflect.Indirect(ptr)

	if src, ok := mapped.(map[interface{}]interface{}); ok {

		// member assign
		for _, member := range un.Members {
			if buf, exist := src[member.Keyname]; exist {
				if varible, err := member.Unmarshal.Unmarshal(buf, property+"."+member.Keyname); err != nil {
					return ptr, err
				} else {
					val.Field(member.FieldId).Set(reflect.ValueOf(varible))
				}
			} else if member.Require {
				return ptr, MakeErrCannotFoundProperty(property + "." + member.Keyname)
			}
		}

		return ptr.Interface(), nil

	} else if src, ok := mapped.(map[string]interface{}); ok {

		// member assign
		for _, member := range un.Members {
			if buf, exist := src[member.Keyname]; exist {
				if varible, err := member.Unmarshal.Unmarshal(buf, property+"."+member.Keyname); err != nil {
					return ptr, err
				} else {
					val.Field(member.FieldId).Set(reflect.ValueOf(varible))
				}
			} else if member.Require {
				return ptr, MakeErrCannotFoundProperty(property + "." + member.Keyname)
			}
		}

		return ptr.Interface(), nil

	} else {
		return ptr, MakeErrInvalidType(property, src, mapped)
	}
}
