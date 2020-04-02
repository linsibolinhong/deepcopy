package deepcopy

import (
	"reflect"
)

func SimilarCopy(dst, src interface{}) {
	dstVal := reflect.ValueOf(dst)
	srcVal := reflect.ValueOf(src)

	if !dstVal.IsValid() || !srcVal.IsValid() {
		return
	}

	if dstVal.Kind() != reflect.Ptr || srcVal.Kind() != reflect.Ptr {
		return
	}

	if dstVal.IsNil() {
		return
	}

	similarCopyResource(dstVal.Elem(), srcVal.Elem())
}

func similarCopyResource(dst, src reflect.Value) {
	if dst.Kind() != src.Kind() {
		return
	}
	switch src.Kind() {
	case reflect.Ptr:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.New(dst.Elem().Type()))
		similarCopyResource(dst.Elem(), src.Elem())
	case reflect.Slice:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeSlice(dst.Type(), src.Len(), src.Cap()))
		for i := 0; i < src.Len(); i++ {
			similarCopyResource(dst.Index(i), src.Index(i))
		}
	case reflect.Map:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeMap(dst.Type()))
		for _, key := range src.MapKeys() {
			originalValue := src.MapIndex(key)
			copyValue := reflect.New(dst.Type().Elem()).Elem()
			similarCopyResource(copyValue, originalValue)
			copyKey := reflect.New(dst.Type().Key()).Elem()
			similarCopyResource(copyKey, key)
			dst.SetMapIndex(copyKey, copyValue)
		}
	case reflect.Struct:
		for i := 0; i < src.NumField(); i++ {
			// private field cannot be settable
			dstV := dst.FieldByName(src.Type().Field(i).Name)
			if dstV.IsValid() && dstV.CanSet() {
				similarCopyResource(dstV, src.Field(i))
			}
		}
	case reflect.Interface:
		if src.IsNil() {
			return
		}
		originalValue := src.Elem()
		copyValue := reflect.New(dst.Type()).Elem()
		similarCopyResource(copyValue, originalValue)
		dst.Set(copyValue)
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16,reflect.Int8:
		dst.SetInt(src.Int())
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		dst.SetUint(src.Uint())
	case reflect.Float32, reflect.Float64:
		dst.SetFloat(src.Float())
	case reflect.Bool:
		dst.SetBool(src.Bool())
	case reflect.String:
		dst.SetString(src.String())
	default:
		if dst.Type() != src.Type() {
			return
		}
		dst.Set(src)
	}
}
