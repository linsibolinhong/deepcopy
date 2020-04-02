package deepcopy

import (
	"fmt"
	"reflect"
)

func DeepCopy(v interface{}) interface{} {
	val := reflect.ValueOf(v)
	fmt.Println(val.Kind())
	ret := reflect.New(val.Type()).Elem()

	if val.IsNil() {
		return v
	}

	copyResource(&ret, &val)
	return ret.Interface()
}


func copyResource(dst, src *reflect.Value) {
	if src.CanAddr() && src.IsNil() {
		*dst = *src
		return
	}

	switch src.Kind() {
	case reflect.Ptr:
		srcVal := src.Elem()
		fmt.Println(srcVal.CanAddr(), srcVal.CanInterface(), srcVal.IsValid(), srcVal.IsZero())
		if srcVal.CanAddr() && srcVal.IsNil() {
			dst.Set(srcVal)
			return
		}
		dstVal := reflect.New(srcVal.Type())
		copyResource(&dstVal, &srcVal)
	}
}