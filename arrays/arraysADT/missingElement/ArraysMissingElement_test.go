package missingElement

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	arr      [20]int
	expected int
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{19, 20, 22, 23, 24, 25},
			21,
		},
		{
			[20]int{6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17},
			12,
		},
		{
			[20]int{200, 202, 203},
			201,
		},
		{
			[20]int{1005, 1007},
			1006,
		},
		{
			[20]int{40, 41, 42, 43, 44, 46, 47, 48, 49, 50},
			45,
		},
	}
}

var (
	impls []ArraysMissingElement
)

func TestMissingElements(t *testing.T) {
	impls = append(impls, &ImplS1A{})

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualOpt := impl.arraysMissingElement(tt.arr)
			var pOrF string
			if assert.Equal(t, tt.expected, actualOpt) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s, expected int = %v", i+1, pOrF, tt.expected)
		}
	}
}
