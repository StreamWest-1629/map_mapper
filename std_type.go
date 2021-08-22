// std_type.go
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
	"fmt"
	"strconv"
)

func UnmarshalInt64(src interface{}, property string) (dest int64, err error) {
	if val, ok := src.(int64); ok {
		return val, nil
	} else if val, ok := src.(int32); ok {
		return int64(val), nil
	} else if val, ok := src.(int16); ok {
		return int64(val), nil
	} else if val, ok := src.(int8); ok {
		return int64(val), nil
	} else if val, ok := src.(uint32); ok {
		return int64(val), nil
	} else if val, ok := src.(uint16); ok {
		return int64(val), nil
	} else if val, ok := src.(uint8); ok {
		return int64(val), nil
	} else if val, ok := src.(string); ok {
		if res, err := strconv.ParseInt(val, 0, 64); err != nil {
			return 0, MakeErrInvalidType(property, res, val)
		} else {
			return res, nil
		}
	} else {
		return 0, MakeErrInvalidType(property, int64(0), src)
	}
}

func UnmarshalFloat64(src interface{}, property string) (dest float64, err error) {
	if val, ok := src.(float64); ok {
		return val, nil
	} else if val, ok := src.(float32); ok {
		return float64(val), nil
	} else if val, ok := src.(string); ok {
		if res, err := strconv.ParseFloat(val, 64); err != nil {
			return 0.0, MakeErrInvalidType(property, res, val)
		} else {
			return res, nil
		}
	} else {
		return 0.0, MakeErrInvalidType(property, float64(0.0), src)
	}
}

func UnmarshalBool(src interface{}, property string) (dest bool, err error) {
	if val, ok := src.(bool); ok {
		return val, nil
	} else if val, ok := src.(string); ok {
		if res, err := strconv.ParseBool(val); err != nil {
			return false, MakeErrInvalidType(property, res, val)
		} else {
			return res, nil
		}
	} else {
		return false, MakeErrInvalidType(property, false, src)
	}
}

func UnmarshalString(src interface{}, property string) (dest string, err error) {
	if val, ok := src.(string); ok {
		return val, nil
	} else {
		return "", MakeErrInvalidType(property, "", src)
	}
}

func MapProperty(mapped map[interface{}]interface{}, key interface{}, property string) (val interface{}, err error) {
	if res, ok := mapped[key]; !ok {
		return nil, MakeErrCannotFoundProperty(property + "'s " + fmt.Sprint(key))
	} else {
		return res, nil
	}
}

func UnmarshalStringArray(src interface{}, property string) (dest []string, err error) {
	if val, ok := src.([]string); ok {
		return val, nil
	} else if vals, ok := src.([]interface{}); ok {
		dest = make([]string, 0, len(vals))
		for i, val := range vals {
			if buf, err := UnmarshalString(val, property+"["+strconv.Itoa(i)+"]"); err != nil {
				return nil, err
			} else {
				dest = append(dest, buf)
			}
		}
		return dest, nil
	} else {
		return nil, MakeErrInvalidType(property, dest, src)
	}
}
