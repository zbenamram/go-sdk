package cron

import (
	"testing"
	"time"

	"github.com/blend/go-sdk/assert"
	"github.com/blend/go-sdk/exception"
	"github.com/blend/go-sdk/stringutil"
)

type stringScheduleTestCase struct {
	Input       string
	ExpectedErr error
	Expected    *time.Time
	After       *time.Time
}

func TestParseString(t *testing.T) {
	assert := assert.New(t)

	testCases := []stringScheduleTestCase{
		{Input: "", ExpectedErr: ErrStringScheduleInvalid},
		{Input: stringutil.Random(stringutil.Letters, 10), ExpectedErr: ErrStringScheduleInvalid},
	}

	for _, tc := range testCases {
		parsed, err := ParseString(tc.Input)
		if tc.ExpectedErr != nil {
			assert.NotNil(err)
			assert.True(exception.Is(err, tc.ExpectedErr))
		} else {
			assert.Equal(tc.Expected, parsed.Next(tc.After))
		}
	}
}
