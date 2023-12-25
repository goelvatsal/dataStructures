package Arrays2D

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	expectedOut string
}

func testData() []ArraysTest {
	return []ArraysTest{
		{"1 2 3 4 \n" + "2 4 6 8 \n" + "1 3 5 7 \n"},
	}
}

var impls []Arrays2DPrinter

func TestArrays2D(t *testing.T) {
	impls = append(impls,
		&ImplS1A{})

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	assert.NotEmpty(t, impls, "no implementations present!!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualOpt := impl.arrays2D()
			var pOrF string
			if assert.Equal(t, tt.expectedOut, actualOpt) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s and expected output=%v", i+1, pOrF, tt.expectedOut)
		}
	}
}
