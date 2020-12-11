package goconfig

import (
	"errors"
	"fmt"
	"reflect"

	pkgErr "github.com/pkg/errors"
)

var ErrTypeMiss = errors.New("type miss")

func errTypeMiss(i interface{}) (err error) {
	err = fmt.Errorf("%w: %s", ErrTypeMiss, reflect.TypeOf(i).String())
	err = pkgErr.WithStack(err)
	return
}

// EasyConfig 对 config 的封装，更易使用
type EasyConfig interface {
	Config
	GetString(key string) (value string, err error)
	GetDefaultString(key string, def string) (value string)

	GetInt(key string) (value int, err error)
	GetDefaultInt(key string, def int) (value int)

	GetFloat(key string) (value float64, err error)
	GetDefaultFloat(key string, def float64) (value float64)

	GetBool(key string) (value bool, err error)
	GetDefaultBool(key string, def bool) (value bool)
}

// easyConfig EasyConfig 的实现
type easyConfig struct {
	Config
}

func newEasyConfig(base Config) (newConfig *easyConfig) {
	newConfig = &easyConfig{
		Config: base,
	}
	return newConfig
}

func (c *easyConfig) GetString(key string) (value string, err error) {
	v, err := c.Config.Get(key)
	if err != nil {
		return
	}

	value, ok := v.(string)
	if !ok {
		err = errTypeMiss(v)
		return
	}

	return
}

func (c *easyConfig) GetDefaultString(key string, def string) (value string) {
	value, err := c.GetString(key)
	if err == nil {
		return
	}

	value = def

	return
}

func (c *easyConfig) GetInt(key string) (value int, err error) {
	v, err := c.Config.Get(key)
	if err != nil {
		return
	}

	value, ok := v.(int)
	if !ok {
		vF, okF := v.(float64)
		if okF {
			ok = true
			value = int(vF)
		}
	}
	if !ok {
		err = errTypeMiss(v)
		return
	}

	return
}

func (c *easyConfig) GetDefaultInt(key string, def int) (value int) {
	value, err := c.GetInt(key)
	if err == nil {
		return
	}

	value = def

	return
}

func (c *easyConfig) GetFloat(key string) (value float64, err error) {
	v, err := c.Config.Get(key)
	if err != nil {
		return
	}

	value, ok := v.(float64)
	if !ok {
		err = errTypeMiss(v)
		return
	}

	return
}

func (c *easyConfig) GetDefaultFloat(key string, def float64) (value float64) {
	value, err := c.GetFloat(key)
	if err == nil {
		return
	}

	value = def

	return
}

func (c *easyConfig) GetBool(key string) (value bool, err error) {
	v, err := c.Config.Get(key)
	if err != nil {
		return
	}

	value, ok := v.(bool)
	if !ok {
		err = errTypeMiss(v)
		return
	}

	return
}

func (c *easyConfig) GetDefaultBool(key string, def bool) (value bool) {
	value, err := c.GetBool(key)
	if err == nil {
		return
	}

	value = def

	return
}
