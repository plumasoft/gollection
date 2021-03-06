package gollection_test

import (
	"testing"

	"github.com/azihsoyn/gollection"
	"github.com/stretchr/testify/assert"
)

func TestDistinct(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 1, 2, 3}
	expect := []int{1, 2, 3}

	res, err := gollection.New(arr).Distinct().Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestDistinct_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").Distinct().Result()
	assert.Error(err)
}

func TestDistinct_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").Distinct().Distinct().Result()
	assert.Error(err)

	res := []int{}
	err = gollection.New("not slice value").Distinct().Distinct().ResultAs(&res)
	assert.Error(err)
}

func TestDistinctBy(t *testing.T) {
	assert := assert.New(t)
	arr := []string{"aaa", "bb", "c", "ddd", "ee", "f"}
	expect := []string{"aaa", "bb", "c"}

	res, err := gollection.New(arr).DistinctBy(func(v string) int {
		return len(v)
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}
func TestDistinctBy_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").DistinctBy(func(v interface{}) interface{} {
		return v
	}).Result()
	assert.Error(err)
}
func TestDistinctBy_NotFunc(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New([]int{}).DistinctBy(0).Result()
	assert.Error(err)
}

func TestDistinctBy_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").
		DistinctBy(func(v interface{}) interface{} {
			return v
		}).
		DistinctBy(func(v interface{}) interface{} {
			return v
		}).
		Result()
	assert.Error(err)
}

func TestDistinct_Stream(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 1, 2, 3}
	expect := []int{1, 2, 3}

	res, err := gollection.NewStream(arr).Distinct().Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestDistinct_Stream_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.NewStream("not slice value").Distinct().Result()
	assert.Error(err)
}

func TestDistinct_Stream_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.NewStream("not slice value").Distinct().Distinct().Result()
	assert.Error(err)
}

func TestDistinctBy_Stream(t *testing.T) {
	assert := assert.New(t)
	arr := []string{"aaa", "bb", "c", "ddd", "ee", "f"}
	expect := []string{"aaa", "bb", "c"}

	res, err := gollection.NewStream(arr).DistinctBy(func(v string) int {
		return len(v)
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}
func TestDistinctBy_Stream_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.NewStream("not slice value").DistinctBy(func(v interface{}) interface{} {
		return v
	}).Result()
	assert.Error(err)
}
func TestDistinctBy_Stream_NotFunc(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.NewStream([]int{}).DistinctBy(0).Result()
	assert.Error(err)
}

func TestDistinctBy_Stream_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.NewStream("not slice value").
		DistinctBy(func(v interface{}) interface{} {
			return v
		}).
		DistinctBy(func(v interface{}) interface{} {
			return v
		}).
		Result()
	assert.Error(err)
}
