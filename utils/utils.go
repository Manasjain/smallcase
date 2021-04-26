package utils

import (
	"encoding/json"
	"reflect"
)

func IsNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}

var Marshal = func(inputData interface{}) string {
	inputDataJson, _ := json.Marshal(inputData)
	return string(inputDataJson)
}

func DeepClone(obj interface{}, copy interface{}) {
	x, _ := json.Marshal(obj)
	_ = json.Unmarshal(x, &copy)
}