package arrays

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
		{"Elements are: 2 3 4 5 6 \n"},
	}
}

var (
	impls []ArrayPrinter
)

func TestArraysPrinter(t *testing.T) {
	impls = append(impls,
		&ImplS1A{})

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	assert.NotEmpty(t, impls, "no implementations present!!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualOpt := impl.arrayPrinter()
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
