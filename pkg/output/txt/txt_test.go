package txt

import (
	"flexrea/pkg/model"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsDuplicated(t *testing.T) {
	testCases := []struct {
		desc     string
		output   *txt
		param1   model.Result
		expected []byte
	}{
		{
			desc:   "returns with no error",
			output: New(),
			param1: model.Result{
				TotalCopy: 101,
			},
			expected: []byte("101"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			b, err := testCase.output.GenerateOutput(testCase.param1)
			if err != nil {
				t.Error(err)
			}
			valid := cmp.Equal(testCase.expected, b)
			if !valid {
				t.Error(cmp.Diff(testCase.expected, b))
			}
		})
	}
}
