package samples

import (
	"testing"

	"github.com/eaglexiang/goconfig"
	"github.com/stretchr/testify/assert"
)

func Test_configFile(t *testing.T) {
	c, _ := goconfig.New("./test_conf.json")

	err := c.Set("test1.test2.test3", "value3")
	if !assert.NoError(t, err) {
		return
	}

	err = c.Set("test1.test2_2", "value2_2")
	if !assert.NoError(t, err) {
		return
	}

	v, err := c.Get("test1.test2_2")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "value2_2", v)
}
