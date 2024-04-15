package delete

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	init        [20]int
	index       int
	expectedOut string
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{1, 2, 3, 4, 5, 5, 6},
			4,
			"Elements are: 1 2 3 4 5 6 \n",
		},
		{
			[20]int{10, 10, 20, 30, 40, 50},
			0,
			"Elements are: 10 20 30 40 50 \n",
		},
		{
			[20]int{1, 3, 5, 7, 9, 10, 11},
			5,
			"Elements are: 1 3 5 7 9 11 \n",
		},
		{
			[20]int{6, 12, 18, 24, 30, 36, 40},
			6,
			"Elements are: 6 12 18 24 30 36 \n",
		},
	}
}

var (
	impls []ArrayDeleter
)

func TestArrayDeleter(t *testing.T) {
	impls = append(impls,
		&Impl1SA{},
	)

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	assert.NotEmpty(t, impls, "no implementations present!!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualOpt := impl.arrayDeleter(tt.init, tt.index)
			var pOrF string

			if assert.Equal(t, tt.expectedOut, actualOpt) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s, expected output=%v", i+1, pOrF, tt.expectedOut)
		}
	}
}
