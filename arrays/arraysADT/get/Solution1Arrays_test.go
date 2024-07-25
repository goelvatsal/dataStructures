package get

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	base        [20]int
	index       int
	expectedOut int
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{20, 40, 60, 80},
			3,
			80,
		},
		{
			[20]int{48, 26, 55, 99, 20},
			0,
			48,
		},
		{
			[20]int{100, 50, 25, 12, 6, 3},
			5,
			3,
		},
		{
			[20]int{99, 108, 43, 1005, 25},
			2,
			43,
		},
	}
}

var (
	impls []Getter
)

func TestArrayGetter(t *testing.T) {
	impls = append(impls,
		&ImplS1A{},
	)

	tests := testData()
	t.Logf("[#tests=%d]\n", len(tests))

	assert.NotEmpty(t, impls, "no implementations present!")

	for _, impl := range impls {
		t.Logf("Using: %T\n", impl)

		for i, tt := range tests {
			actualOpt := impl.get(tt.base, tt.index)
			var pOrF string
			if assert.Equal(t, tt.expectedOut, actualOpt) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test %d, %s, expected output=%v", i+1, pOrF, tt.expectedOut)
		}
	}
}
