package parallel

import (
	"testing"
	"github.com/mkideal/pkg/math/random"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestStreamingWorkPool(t *testing.T) {
	totalWork := 10240
	pool := MakeStreamingWorkPool(int64(totalWork),
		func(i int64) (item WorkItem, err error) {
			ns := 10000 + random.Intn(100, random.DefaultSource)
			time.Sleep(time.Duration(ns) * time.Nanosecond)
			// must cast because it doesn't cast below
			return int(i * i), nil
		})

	for i := 0; i < totalWork; i++ {
		assert.Equal(t, i*i, pool.Get(int64(i)))
	}

	// do it again, full traversal
	for i := 0; i < totalWork; i++ {
		assert.Equal(t, i*i, pool.Get(int64(i)))
	}

	// test random access (not expected to be fast)
	for i := 0; i < totalWork/10; i++ {
		idx := random.Intn(totalWork, random.DefaultSource)
		assert.Equal(t, idx*idx, pool.Get(int64(idx)))
	}
}

func TestStreamingWorkPool0(t *testing.T) {
	totalWork := 1024
	pool := MakeStreamingWorkPool0(
		int64(totalWork),
		32, 32, 8,
		func(i int64) (item WorkItem, err error) {
			ns := 100000 + random.Intn(1000, random.DefaultSource)
			time.Sleep(time.Duration(ns) * time.Nanosecond)
			// must cast because it doesn't cast below
			return int(i * i), nil
		})

	for i := 0; i < totalWork; i++ {
		assert.Equal(t, i*i, pool.Get(int64(i)))
	}

	// do it again, full traversal
	for i := 0; i < totalWork; i++ {
		assert.Equal(t, i*i, pool.Get(int64(i)))
	}

	// test random access (not expected to be fast)
	for i := 0; i < totalWork/10; i++ {
		idx := random.Intn(totalWork, random.DefaultSource)
		assert.Equal(t, idx*idx, pool.Get(int64(idx)))
	}
}
