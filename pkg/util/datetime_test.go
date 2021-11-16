package util

import (
	"fmt"
	"testing"
)

func TestStringToUnix(t *testing.T) {
	dt, _ := StringToUnix("2021-10-10 01:02:03")
	fmt.Println(dt)
}
