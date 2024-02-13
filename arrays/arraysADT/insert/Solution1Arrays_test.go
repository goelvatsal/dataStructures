package insert

import (
	"dataStructures"
	append2 "dataStructures/arrays/arraysADT/append"
	"dataStructures/arrays/arraysADT/display"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	init        [20]int
	index       int
	x           int
	expectedOut string
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{1, 3, 4, 5, 6},
			1,
			2,
			"Elements are: 1 2 3 4 5 6 \n",
		},
		{
			[20]int{10, 30, 40, 50, 60},
			1,
			20,
			"Elements are: 10 20 30 40 50 60 \n",
		},
		{
			[20]int{100, 200, 300, 400, 500, 700},
			5,
			600,
			"Elements are: 100 200 300 400 500 600 700 \n",
		},
		{
			[20]int{50, 100, 150, 200, 250},
			5,
			300,
			"Elements are: 50 100 150 200 250 300 \n",
		},
	}
}

var (
	impls []ArrayInserter
)

func TestArrayInserter(t *testing.T) {
	impls = append(impls,
		&ImplS1A{},
	)

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	assert.NotEmpty(t, impls, "no implementations present!!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualOpt := impl.arrayInserter(tt.init, tt.index, tt.x)
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

func TestArrays(t *testing.T) {
	arr := Array{append2.Array{display.Array{A: [20]int{1, 3, 4, 5}, Size: 20, Length: 4}}}
	arr.Insert(1, 2)
	if !assert.Equal(t, "Elements are: 1 2 3 4 5 \n", arr.Display()) {
		t.Fatal("not equal")
	}
}
