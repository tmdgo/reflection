package functions

import "reflect"

func CallFunc(function interface{}, inputs map[reflect.Type]reflect.Value) (results []reflect.Value) {
	funcValue := reflect.ValueOf(function)
	numberOfInputs := funcValue.Type().NumIn()

	funcInputs := make([]reflect.Value, numberOfInputs)

	for index := 0; index < numberOfInputs; index++ {
		inputType := funcValue.Type().In(index)
		funcInputs[index] = inputs[inputType]
	}
	results = funcValue.Call(funcInputs)
	return
}
