package linSearch

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	init        [20]int
	key         int
	expectedOut int
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{10, 20, 30, 40, 50},
			30,
			2,
		},
		{
			[20]int{100, 200, 300},
			100,
			0,
		},
		{
			[20]int{300, 600, 900, 1200},
			1200,
			3,
		},
		{
			[20]int{10, 43, 87, 99, 346, 222},
			512,
			-1,
		},
		{
			[20]int{76, 23, 99, 35, 105},
			35,
			3,
		},
	}
}

var (
	impls []ArrayLinSearcher
)

func TestArrayLinSearcher(t *testing.T) {
	impls = append(impls,
		&Impl1SA{},
	)

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	assert.NotEmpty(t, impls, "no implementations present!!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualOpt := impl.ArrayLinSearcher(tt.init, tt.key)
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
