package methods

import "reflect"

func CallMethodByName(model interface{}, name string, inputs map[reflect.Type]reflect.Value) (results []reflect.Value) {
	funcValue := reflect.ValueOf(model).MethodByName(name)
	numberOfInputs := funcValue.Type().NumIn()

	funcInputs := make([]reflect.Value, numberOfInputs)

	for index := 0; index < numberOfInputs; index++ {
		inputType := funcValue.Type().In(index)
		funcInputs[index] = inputs[inputType]
	}
	results = funcValue.Call(funcInputs)
	return
}
