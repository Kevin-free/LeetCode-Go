package _380_insert_delete_getrandom_o1

import (
	"fmt"
	"testing"
)

func Test0380(t *testing.T) {
	obj := Constructor()
	param_1 := obj.Insert(1)
	param_2 := obj.Remove(2)
	param_3 := obj.Insert(2)
	param_4 := obj.GetRandom()
	param_5 := obj.Remove(1)
	param_6 := obj.Insert(2)
	param_7 := obj.GetRandom()
	res := []interface{}{param_1, param_2, param_3, param_4, param_5, param_6, param_7}
	fmt.Print(res)
}
