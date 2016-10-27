package gollection_test

import (
	"testing"

	"github.com/azihsoyn/gollection"
	"github.com/stretchr/testify/assert"
)

func TestFold(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := 155

	res, err := gollection.New(arr).Fold(100, func(v1, v2 interface{}) interface{} {
		n1, ok1 := v1.(int)
		n2, ok2 := v2.(int)
		if ok1 && ok2 {
			return n1 + n2
		}
		return ""
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestFold_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").Fold(100, func(v1, v2 interface{}) interface{} {
		return ""
	}).Result()
	assert.Error(err)
}

func TestFold_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").
		Fold(100, func(v1, v2 interface{}) interface{} {
		return ""
	}).
		Fold(100, func(v1, v2 interface{}) interface{} {
		return ""
	}).
		Result()
	assert.Error(err)
}
