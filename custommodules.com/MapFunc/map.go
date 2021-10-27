package mapfunc

import "reflect"

//TODO: Interface slice discussion https://github.com/golang/go/wiki/InterfaceSlice
type modifierFunc func(interface{}) interface{}

func Map(items interface{}, modifier modifierFunc) []interface{} {
	// Checkout what reflect package does
	switch reflect.TypeOf(items).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(items)
		arr := make([]interface{}, s.Len())
		for i := 0; i < s.Len(); i++ {
			arr[i] = modifier(s.Index(i).Interface())
		}
		return arr
	}
	return nil
}
