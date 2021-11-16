package dblog

import (
	"fmt"
	"testing"
)

func TestONData(t *testing.T) {
	type s struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	a := s{Name: "lucy", Age: 20}
	b := s{Name: "lili", Age: 30}

	str := NewONData().OldStruct(a).NewStruct(b).Json()
	fmt.Println(str)
}
