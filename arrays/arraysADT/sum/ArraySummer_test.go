package sum

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	base        [20]int
	expectedSum int
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{10, 20, 30, 40, 50},
			150,
		},
		{
			[20]int{100, 200, 50, 25, 10},
			385,
		},
		{
			[20]int{50, 250, 1250, 6250, 31250},
			39050,
		},
		{
			[20]int{8, 16, 32, 64, 128},
			248,
		},
		{
			[20]int{-10, -50, -70, 80, 50},
			0,
		},
	}
}

var (
	impls []ArraySummer
)

func TestSum(t *testing.T) {
	impls = append(impls, &ImplS1A{})

	tests := testData()
	t.Logf("[#tests=%d] \n", len(tests))

	assert.NotEmpty(t, impls, "no implementations present!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualSum := impl.arraySummer(tt.base)
			var pOrF string
			if assert.Equal(t, tt.expectedSum, actualSum) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test %d, %s, expected output = %v", i+1, pOrF, tt.expectedSum)
		}
	}
}
