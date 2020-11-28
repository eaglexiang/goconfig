package goconfig

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SetGet(t *testing.T) {
	config := newConfig()

	_, err := config.Get("test.key.1")
	if !assert.True(t, errors.Is(err, ErrNoKey)) {
		return
	}

	err = config.Set("test.key.1", "test_value")
	if !assert.NoError(t, err) {
		return
	}

	v, err := config.Get("test.key.1")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, v, "test_value")
}

func Test_NotSubConfig(t *testing.T) {
	config := newConfig()

	err := config.Set("test.key", "test_value_0")
	if !assert.NoError(t, err) {
		return
	}

	// test.key.1 应该无法写入，因为 test.key 已经被占用
	err = config.Set("test.key.1", "test_value_1")
	if !assert.True(t, errors.Is(err, ErrNotSubConfig)) {
		return
	}

	_, err = config.Get("test.key.1")
	assert.True(t, errors.Is(err, ErrNotSubConfig))
}

func Test_EmptyKey(t *testing.T) {
	config := newConfig()

	err := config.Set("", "test_value")
	if !assert.True(t, errors.Is(err, ErrEmptyKey)) {
		return
	}
	err = config.Set(" ", "test_value")
	if !assert.True(t, errors.Is(err, ErrEmptyKey)) {
		return
	}

	_, err = config.Get("")
	if !assert.True(t, errors.Is(err, ErrEmptyKey)) {
		return
	}
	_, err = config.Get("  ")
	if !assert.True(t, errors.Is(err, ErrEmptyKey)) {
		return
	}
}
