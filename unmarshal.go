// unmarshal.go
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

import "reflect"

type (
	Unmarshalizer interface {
		Unmarshal(src interface{}, property string) (dest interface{}, err error)
	}
	UnmarshalFunc func(src interface{}, property string) (dest interface{}, err error)
)

func Unmarshal(u Unmarshalizer, src interface{}) (dest interface{}, err error) {
	return u.Unmarshal(src, "")
}

func (f UnmarshalFunc) Unmarshal(src interface{}, property string) (dest interface{}, err error) {
	return f(src, property)
}

func GetUnmarshal(ty reflect.Type) (unmarshal Unmarshalizer) {

	switch ty.Kind() {
	case reflect.Int64:
		unmarshal = UnmarshalFunc(func(src interface{}, property string) (dest interface{}, err error) {
			return UnmarshalInt64(src, property)
		})
		return
	case reflect.Float64:
		unmarshal = UnmarshalFunc(func(src interface{}, property string) (dest interface{}, err error) {
			return UnmarshalFloat64(src, property)
		})
		return
	case reflect.Bool:
		unmarshal = UnmarshalFunc(func(src interface{}, property string) (dest interface{}, err error) {
			return UnmarshalBool(src, property)
		})
		return
	case reflect.String:
		unmarshal = UnmarshalFunc(func(src interface{}, property string) (dest interface{}, err error) {
			return UnmarshalString(src, property)
		})
		return
	case reflect.Slice:
		switch ty.Elem().Kind() {
		case reflect.String:
			unmarshal = UnmarshalFunc(func(src interface{}, property string) (dest interface{}, err error) {
				return UnmarshalStringArray(src, property)
			})
			return
		}
	}

	panic("not supported type")
}
