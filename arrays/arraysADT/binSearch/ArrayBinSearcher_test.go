package binSearch

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	init        [20]int
	key         int
	expectedOpt int
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{10, 20, 30, 40, 50},
			30,
			2,
		},
		{
			[20]int{5, 10, 15, 20, 30},
			5,
			0,
		},
		{
			[20]int{100, 200, 300, 400},
			400,
			3,
		},
		{
			[20]int{50, 100, 150, 200, 300},
			250,
			-1,
		},
	}
}

var (
	impls []ArrayBinSearcher
)

func TestArrayBinSearcher(t *testing.T) {
	impls = append(impls,
		&ImplS1A{},
	)
	assert.NotEmpty(t, impls, "no implementations present!!")

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualOpt := impl.arrayBinSearcher(tt.init, tt.key)
			var pOrF string
			if assert.Equal(t, tt.expectedOpt, actualOpt) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s, expected output=%v", i+1, pOrF, tt.expectedOpt)
		}
	}
}
