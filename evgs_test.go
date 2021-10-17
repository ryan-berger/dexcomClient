package dexcomClient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRanges(t *testing.T) {
	tests := []struct{
		start string
		end string
		expected []queryParam
		isErr bool
	} {
		// empty
		{isErr: true},
		{start: "2020-01-02T15:04:05", isErr: true},
		{start: "2020-01-02T15:04:05", end: "2020-04-01T15:04:05"},
		{start: "2020-01-02T15:04:05", end: "2020-04-02T15:04:05"},
	}

	for _, test := range tests {
		ranges, err := getEGVRanges(test.start, test.end)
		if test.isErr {
			assert.Error(t, err)
			continue
		}
		assert.NoError(t, err)
		assert.Equal(t, test.expected, ranges)
	}
}