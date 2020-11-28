package goconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_split(t *testing.T) {
	key := "test1.test2.test3"

	path, key := splitPath(key)

	if !assert.Equal(t, 2, len(path)) {
		return
	}
	if !assert.Equal(t, "test1", path[0]) {
		return
	}
	if !assert.Equal(t, "test2", path[1]) {
		return
	}

	assert.Equal(t, "test3", key)
}
