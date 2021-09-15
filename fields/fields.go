package fields

import (
	"fmt"
	"reflect"

	"github.com/tmdgo/reflection/interfaces"
)

func GetTypeAndValue(model interface{}, fieldName string) (fieldType reflect.Type, fieldValue interface{}, err error) {
	e := reflect.ValueOf(model).Elem()
	for i := 0; i < e.NumField(); i++ {
		if fieldName == e.Type().Field(i).Name {
			fieldType = e.Type().Field(i).Type
			fieldValue = e.Field(i).Interface()
			return
		}
	}
	err = fmt.Errorf(
		`field "%s" not found in model "%v"`,
		fieldName,
		interfaces.GetTypeName(model),
	)
	return
}
