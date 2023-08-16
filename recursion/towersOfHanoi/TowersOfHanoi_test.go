package TowersOfHanoi

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TOHTest struct {
	n           int
	a           int
	b           int
	c           int
	expectedOut string
}

func testData() []TOHTest {
	return []TOHTest{
		{
			3,
			1,
			2,
			3,
			`(1, 3)
(1, 2)
(3, 2)
(1, 3)
(2, 1)
(2, 3)
(1, 3)
`,
		},
		{
			2,
			1,
			2,
			3,
			`(1, 3)
(1, 2)
(3, 2)
(1, 3)
(2, 1)
(2, 3)
(1, 3)
(1, 2)
(1, 3)
(2, 3)
`,
		},
	}
}

var (
	impls []TOHCalculator
)

func TestTowersOfHanoi(t *testing.T) {
	impls = append(impls,
		&ImplS1R{},
	)

	tests := testData()
	t.Logf("[tests=%d]", len(tests))

	assert.NotEmpty(t, tests, "no implementations present!!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualOpt := impl.towersOfHanoi(tt.n, tt.a, tt.b, tt.c)
			var pOrF string
			if assert.Equal(t, tt.expectedOut, actualOpt) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s, input=[%d, %d, %d, %d]", i+1, pOrF, tt.n, tt.a, tt.b, tt.c)
		}
	}
}
