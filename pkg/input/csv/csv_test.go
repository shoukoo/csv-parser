package csv

import (
	"testing"

	"github.com/shoukoo/csv-parser/pkg/model"

	"github.com/google/go-cmp/cmp"
)

func TestRead(t *testing.T) {
	testCases := []struct {
		desc        string
		csvReader   *csvInput
		searchValue string
		expected    []model.Row
	}{
		{
			desc:        "return Rows when searching application id 1",
			csvReader:   New("test/sample.csv"),
			searchValue: "1",
			expected:    getCSVMockSearchApp1(),
		},
		{
			desc:        "return Rows without duplicate entries when searching application id 3",
			csvReader:   New("test/sample.csv"),
			searchValue: "3",
			expected:    getCSVMockSearchApp3(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			result, err := testCase.csvReader.Read(testCase.searchValue)
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

func TestIsDuplicated(t *testing.T) {
	testCases := []struct {
		desc      string
		csvReader *csvInput
		param1    []model.Row
		param2    string
		expected  bool
	}{
		{
			desc:      "return false, Id 103 doesn't exist in param1",
			csvReader: New("test/sample.csv"),
			param1: []model.Row{
				{
					Id:            "101",
					UserId:        "2",
					ApplicationId: "3",
					Type:          "DESKTOP",
					Comment:       "Test",
				},
				{
					Id:            "102",
					UserId:        "2",
					ApplicationId: "3",
					Type:          "DESKTOP",
					Comment:       "Test",
				},
			},
			param2:   "103",
			expected: false,
		},
		{
			desc:      "return true machine Id 101 already exists in param1",
			csvReader: New("test/sample.csv"),
			param1: []model.Row{
				{
					Id:            "101",
					UserId:        "2",
					ApplicationId: "3",
					Type:          "DESKTOP",
					Comment:       "Test",
				},
				{
					Id:            "102",
					UserId:        "2",
					ApplicationId: "3",
					Type:          "DESKTOP",
					Comment:       "Test",
				},
			},
			param2:   "101",
			expected: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			bo := testCase.csvReader.IsDuplicated(
				testCase.param1,
				testCase.param2,
			)
			valid := cmp.Equal(testCase.expected, bo)
			if !valid {
				t.Error(cmp.Diff(testCase.expected, bo))
			}
		})
	}
}
