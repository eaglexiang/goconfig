package goconfig

import (
	"io/ioutil"
	"os"
)

// persistConfig 带持久化的配置
type persistConfig struct {
	Config
	filename string
}

func newPersistConfig(base Config, filename string) (c *persistConfig, err error) {
	c = &persistConfig{
		Config:   base,
		filename: filename,
	}

	err = c.read()

	return
}

func (c *persistConfig) Set(key string, value interface{}) (err error) {
	err = c.Config.Set(key, value)
	if err != nil {
		return
	}

	err = c.save()

	return
}

func (c *persistConfig) save() (err error) {
	b, err := c.Config.Bytes()
	if err != nil {
		return
	}

	err = ioutil.WriteFile(c.filename, b, os.FileMode(0644))

	return
}

func (c *persistConfig) read() (err error) {
	b, err := ioutil.ReadFile(c.filename)
	if err != nil {
		return
	}

	err = c.Config.Import(b)

	return
}
