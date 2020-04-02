package deepcopy

import (
	"fmt"
	"reflect"
	"testing"
)

type I int
type II int
type S struct {
	Z int
	V B
}

type B struct {
	K int
}
type SS struct {
	V BB
}
type BB struct {
	K int
}

func (s *SS) Print() {
	fmt.Println(s.V.K)
}
func Test_DeepCopy(t *testing.T) {
	v := map[I]S{1: S{
		V:B{2},
	}}
	r := map[II]SS{}
	SimilarCopy(&r, &v)
	fmt.Println(reflect.TypeOf(r), reflect.TypeOf(v))
	fmt.Println(v, r)

}
