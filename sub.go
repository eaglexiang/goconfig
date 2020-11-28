package goconfig

import (
	"errors"
	"fmt"
	"reflect"
)

var ErrNotSubConfig = errors.New("sub item for key is not sub config")

func errNotSubConfig(sub interface{}, key string) (err error) {
	t := reflect.TypeOf(sub)
	err = fmt.Errorf("%w: %s(%s)", ErrNotSubConfig, t.String(), key)
	return
}
