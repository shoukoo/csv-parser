package handler

import (
	"testing"

	"github.com/shoukoo/csv-parser/pkg/model"
	"github.com/shoukoo/csv-parser/pkg/output"
	"github.com/shoukoo/csv-parser/pkg/output/txt"

	"github.com/google/go-cmp/cmp"
)

func TestCalculateCopyCount(t *testing.T) {
	testCases := []struct {
		desc     string
		param1   []model.Row
		expected model.Result
	}{
		{
			desc: "return with no error",
			param1: []model.Row{
				{
					Id:            "101",
					UserId:        "1",
					ApplicationId: "222",
					Type:          "Desktop",
					Comment:       "From System X",
				},
				{
					Id:            "102",
					UserId:        "2",
					ApplicationId: "222",
					Type:          "Laptop",
					Comment:       "From System x",
				},
				{
					Id:            "102",
					UserId:        "2",
					ApplicationId: "222",
					Type:          "Desktop",
					Comment:       "From System x",
				},
			},
			expected: model.Result{
				Laptop:    1,
				Desktop:   2,
				TotalCopy: 2,
			},
		},
		{
			desc: "return with no error when testing different cases",
			param1: []model.Row{
				{
					Id:            "101",
					UserId:        "1",
					ApplicationId: "222",
					Type:          "desktop",
					Comment:       "From System X",
				},
				{
					Id:            "102",
					UserId:        "2",
					ApplicationId: "222",
					Type:          "DESKTOP",
					Comment:       "From System x",
				},
				{
					Id:            "103",
					UserId:        "4",
					ApplicationId: "222",
					Type:          "desKTOP",
					Comment:       "From System x",
				},
				{
					Id:            "104",
					UserId:        "5",
					ApplicationId: "222",
					Type:          "LAPTOP",
					Comment:       "From System x",
				},
				{
					Id:            "1001",
					UserId:        "6",
					ApplicationId: "222",
					Type:          "laptop",
					Comment:       "From System x",
				},
			},
			expected: model.Result{
				Laptop:    2,
				Desktop:   3,
				TotalCopy: 3,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			result, err := calculateCopyCount(testCase.param1)
			if err != nil {
				t.Error(err)
			}
			valid := cmp.Equal(testCase.expected, result)
			if !valid {
				t.Error(cmp.Diff(testCase.expected, result))
			}
		})
	}
}

func TestGetCopyCount(t *testing.T) {
	testCases := []struct {
		desc     string
		input    *inputMock
		output   output.Output
		paraddm1 []model.Row
		expected []byte
	}{
		{
			desc:     "return total count is 2",
			input:    &inputMock{},
			output:   txt.New(),
			expected: []byte("2"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			s := New(testCase.input, testCase.output, "121")
			result, err := s.GetCopyCount()
			if err != nil {
				t.Error(err)
			}
			valid := cmp.Equal(testCase.expected, result)
			if !valid {
				t.Error(cmp.Diff(testCase.expected, result))
			}
		})
	}
}
