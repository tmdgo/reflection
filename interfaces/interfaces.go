package interfaces

import "reflect"

func GetTypeName(typeInterface interface{}) string {
	t := reflect.TypeOf(typeInterface)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}
	return t.Name()
}
