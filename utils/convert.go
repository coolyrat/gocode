package utils

import (
	"errors"
	"fmt"
	"reflect"
)

func MapToStruct(src map[string]interface{}, dest interface{}) error {
	for k, v := range src {
		err := setField(k, v, dest)
		if err != nil {
			return err
		}
	}
	return nil
}

func setField(key string, value, dest interface{}) error {
	structValue := reflect.ValueOf(dest).Elem()
	structValueField := structValue.FieldByName(key)
	if !structValueField.IsValid() {
		return fmt.Errorf("No such field: %s in struct", key)
	}

	if !structValueField.CanSet() {
		return fmt.Errorf("Cannot set %s field value", key)
	}

	structValueType := structValueField.Type()
	val := reflect.ValueOf(value)
	if structValueType != val.Type() {
		return errors.New("Provided value type didn't match struct field type")
	}
	structValueField.Set(val)
	return nil
}
