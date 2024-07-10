package maxMin

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	base        [20]int
	expectedMax int
	expectedMin int
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{10, 50, 100, 20},
			100,
			10,
		},
		{
			[20]int{100, 200, 300, 300, 5},
			300,
			5,
		},
		{
			[20]int{20, 20, 20, 20, 20},
			20,
			20,
		},
		{
			[20]int{999, 998, 100, 350, 850},
			999,
			100,
		},
	}
}

var (
	impls []MaxMin
)

func TestSet(t *testing.T) {
	impls = append(impls,
		&ImplS1A{})

	tests := testData()
	t.Logf("[#tests=%d] \n", len(tests))

	assert.NotEmpty(t, impls, "no implementations present!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualMax, actualMin := impl.maxMin(tt.base)
			var pOrF string
			if assert.Equal(t, tt.expectedMax, actualMax) && assert.Equal(t, tt.expectedMin, actualMin) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s, expected min and max = %v, %v", i+1, pOrF, tt.expectedMin, tt.expectedMin)
		}
	}
}
