package deepcopy

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_DeepCopy(t *testing.T) {
	v1 := 1
	v := &v1
	v = nil
	r := DeepCopy(v)
	fmt.Println(reflect.TypeOf(r), reflect.TypeOf(v))
	fmt.Println(*v, *r.(*int))
}
