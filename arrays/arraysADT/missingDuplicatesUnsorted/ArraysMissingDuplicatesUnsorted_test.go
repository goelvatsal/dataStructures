package missingDuplicatesUnsorted

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	arr         [20]int
	expectedOpt []int
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{8, 3, 6, 4, 6, 5, 6, 8, 2, 7},
			[]int{8, 6},
		},
		{
			[20]int{12, 5, 18, 9, 3, 14, 7, 5, 12, 18},
			[]int{12, 5, 18},
		},
		{
			[20]int{7, 14, 2, 19, 11, 6, 8, 17, 14, 8, 6},
			[]int{14, 6, 8},
		},
		{
			[20]int{15, 8, 1, 16, 4, 10, 13, 5, 5, 1, 16},
			[]int{1, 16, 5},
		},
		{
			[20]int{10, 6, 13, 17, 0, 12, 2, 19, 4, 1},
			[]int(nil),
		},
	}
}

var (
	impls []ArraysMissingDuplicatesUnsorted
)

func TestMissingDuplicatesUnsorted(t *testing.T) {
	impls = append(impls, &ImplS1A{})

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualOpt := impl.arraysMissingDuplicatesUnsorted(tt.arr)
			var pOrF string
			if assert.Equal(t, tt.expectedOpt, actualOpt) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s, expected opt = %v", i+1, pOrF, tt.expectedOpt)
		}
	}
}
