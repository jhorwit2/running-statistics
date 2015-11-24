package anomaly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunningStatistic(t *testing.T) {
	assert := assert.New(t)

	stats := NewRunningStats()
	stats.Add(1)
	stats.Add(2)
	stats.Add(3)

	assert.Equal(3, stats.Len())
	assert.Equal(2.0, stats.Mean())
	assert.Equal(1.0, stats.Stddev())
	assert.Equal(1.0, stats.Min())
	assert.Equal(3.0, stats.Max())

	stats.Clear()
	assert.Equal(0, stats.Len())

	stats.AddAll([]float64{1, 3})
	assert.Equal(2, stats.Len())
	assert.Equal(2.0, stats.Mean())
}
