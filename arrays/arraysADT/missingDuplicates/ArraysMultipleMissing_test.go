package missingDuplicates

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	arr             [20]int
	expectedMissing []int
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{3, 6, 8, 8, 10, 12, 15, 15, 15, 20},
			[]int{8, 15},
		},
		{
			[20]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			[]int{9},
		},
		{
			[20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int(nil),
		},
		{
			[20]int{1, 1, 3, 4, 5, 10, 10, 15, 15},
			[]int{1, 10, 15},
		},
		{
			[20]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
			[]int{1, 2, 3, 4, 5},
		},
	}
}

var (
	impls []ArraysMissingDuplicates
)

func TestMissingDuplicates(t *testing.T) {
	impls = append(impls,
		&ImplS1A{},
		&ImplS2H{})

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualOpt := impl.arraysMissingDuplicates(tt.arr)
			var pOrF string
			if assert.Equal(t, tt.expectedMissing, actualOpt) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s, expected opt = %v", i+1, pOrF, tt.expectedMissing)
		}
	}
}
