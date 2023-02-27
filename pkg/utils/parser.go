package utils

import "encoding/json"

func ParseTo[T any](v interface{}) *T {
	x := new(T)
	model, _ := json.Marshal(v)
	json.Unmarshal(model, &x)
	return x
}
