package web

import "reflect"

func allFieldsNonNil(s interface{}) bool {
	v := reflect.ValueOf(s)

	if v.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.Kind() == reflect.Ptr && field.IsNil() {
			return false
		}
	}
	return true
}
