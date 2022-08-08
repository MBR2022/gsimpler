package handler_test

import (
	"testing"

	"github.com/MBR2022/gosimpler/handler"
)

func Test_add_func(t *testing.T) {
	var i, y = 1, 2
	if v := handler.Add(i, y); v != i+y {
		t.Error("Test Falied")
	}
}
