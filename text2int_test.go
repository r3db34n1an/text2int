package text2int

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	In       string
	Expected int
}

var testCases = []TestCase{
	{
		In:       "seven hundred ninety-nine",
		Expected: 799,
	},
	{
		In:       "zero",
		Expected: 0,
	},
	{
		In:       "one",
		Expected: 1,
	},
	{
		In:       "twenty one thousand three hundred forty-five",
		Expected: 21345,
	},
	{
		In:       "one thousand three hundred thirty-seven",
		Expected: 1337,
	},
	{
		In:       "one thousand three hundred thirty seven",
		Expected: 1337,
	},
	{
		In:       "one thousand three hundred thirty \nseven",
		Expected: 1337,
	},
	{
		In:       "onethousandthreehundredthirtyseven",
		Expected: 1337,
	},
}

func TestConvert(t *testing.T) {
	for _, item := range testCases {
		value, convertError := Convert(item.In)
		assert.Nil(t, convertError)
		assert.Equal(t, item.Expected, value)
	}
}
