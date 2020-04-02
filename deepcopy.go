package deepcopy

import (
	"reflect"
)

func DeepCopy(v interface{}) interface{} {
	val := reflect.ValueOf(v)
	if !val.IsValid() {
		return v
	}
	ret := reflect.New(val.Type()).Elem()

	copyResource(ret, val)
	return ret.Interface()
}


func copyResource(dst, src reflect.Value) {
	switch src.Kind() {
	case reflect.Ptr:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.New(src.Elem().Type()))
		copyResource(dst.Elem(), src.Elem())
	case reflect.Slice:
		if src.IsNil() {
			return
		}
		if dst.Kind() != src.Kind() {
			return
		}
		dst.Set(reflect.MakeSlice(src.Type(), src.Len(), src.Cap()))
		for i := 0; i < src.Len(); i++ {
			copyResource(dst.Index(i), src.Index(i))
		}
	case reflect.Map:
		if src.IsNil() {
			return
		}
		if dst.Kind() != src.Kind() {
			return
		}
		dst.Set(reflect.MakeMap(src.Type()))
		for _, key := range src.MapKeys() {
			originalValue := src.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()
			copyResource(copyValue, originalValue)
			copyKey := DeepCopy(key.Interface())
			dst.SetMapIndex(reflect.ValueOf(copyKey), copyValue)
		}
	case reflect.Struct:
		for i := 0; i < src.NumField(); i++ {
			// private field cannot be settable
			if dst.Field(i).CanSet() {
				copyResource(dst.Field(i), src.Field(i))
			}
		}
	case reflect.Interface:
		if src.IsNil() {
			return
		}
		originalValue := src.Elem()
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyResource(copyValue, originalValue)
		dst.Set(copyValue)
	default:
		dst.Set(src)
	}
}