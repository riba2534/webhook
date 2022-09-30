package utils

import jsoniter "github.com/json-iterator/go"

func MarshalAny2String(v interface{}) string {
	s, e := jsoniter.MarshalToString(v)
	if e != nil {
		return "{}"
	}
	return s
}
