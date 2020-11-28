package goconfig

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrNotSubConfig = errors.New("sub item for key is not sub config")
)

func errNotSubConfig(sub interface{}) (err error) {
	t := reflect.TypeOf(sub)
	err = fmt.Errorf("%w: %s", ErrNotSubConfig, t.String())
	return
}
