// errors.go
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

type ErrInvalidType struct {
	propName string
	wantType reflect.Type
	hasType  reflect.Type
}

type ErrCannotFoundProperty struct {
	propName string
}

func (e *ErrCannotFoundProperty) Error() string {
	return e.propName + " is required property, but cannot found it"
}

func (e *ErrInvalidType) Error() string {
	return e.propName + " has invalid type (want: " + e.wantType.String() + ", has: " + e.hasType.String() + ")"
}

func MakeErrInvalidType(propName string, want interface{}, has interface{}) error {
	return &ErrInvalidType{
		propName: propName,
		wantType: reflect.TypeOf(want),
		hasType:  reflect.TypeOf(has),
	}
}

func MakeErrCannotFoundProperty(propName string) error {
	return &ErrCannotFoundProperty{
		propName: propName,
	}
}
