package utils

import (
	"fmt"
	"testing"
)

func Test_func(t *testing.T) {
	for k := range [100]int{} {
		fmt.Println(k)
	}
}
