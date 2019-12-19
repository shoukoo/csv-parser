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
		param1   map[string][]model.Row
		expected model.Result
	}{
		{
			desc: "Laptop has the same quantity as desktop",
			param1: map[string][]model.Row{

				"1091": []model.Row{
					{
						Id:            "1",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "LAPTOP",
						Comment:       "Exported from System A",
					},
					{
						Id:            "2",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "DESKTOP",
						Comment:       "Exported from System B",
					},
				},
			},
			expected: model.Result{
				Laptop:    0,
				Desktop:   0,
				TotalCopy: 1,
			},
		},
		{
			desc: "Laptop has more than desktop",
			param1: map[string][]model.Row{

				"1091": []model.Row{
					{
						Id:            "15",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "LAPTOP",
						Comment:       "Exported from System A",
					},
					{
						Id:            "1",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "LAPTOP",
						Comment:       "Exported from System A",
					},
					{
						Id:            "12",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "LAPTOP",
						Comment:       "Exported from System A",
					},

					{
						Id:            "13",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "LAPTOP",
						Comment:       "Exported from System A",
					},
					{
						Id:            "2",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "DESKTOP",
						Comment:       "Exported from System B",
					},
				},
			},
			expected: model.Result{
				Laptop:    0,
				Desktop:   0,
				TotalCopy: 3,
			},
		},
		{
			desc: "Desktop has more than laptop",
			param1: map[string][]model.Row{

				"1091": []model.Row{
					{
						Id:            "15",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "DESKTOP",
						Comment:       "Exported from System A",
					},
					{
						Id:            "1",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "DESKTOP",
						Comment:       "Exported from System A",
					},
					{
						Id:            "12",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "DESKTOP",
						Comment:       "Exported from System A",
					},

					{
						Id:            "13",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "DESKTOP",
						Comment:       "Exported from System A",
					},
					{
						Id:            "2",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "laptop",
						Comment:       "Exported from System B",
					},
				},
			},
			expected: model.Result{
				Laptop:    0,
				Desktop:   0,
				TotalCopy: 4,
			},
		},
		{
			desc: "One laptop",
			param1: map[string][]model.Row{

				"1091": []model.Row{
					{
						Id:            "2",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "laptop",
						Comment:       "Exported from System B",
					},
				},
			},
			expected: model.Result{
				Laptop:    0,
				Desktop:   0,
				TotalCopy: 1,
			},
		},
		{
			desc: "One desktop",
			param1: map[string][]model.Row{

				"1091": []model.Row{
					{
						Id:            "2",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "desktop",
						Comment:       "Exported from System B",
					},
				},
			},
			expected: model.Result{
				Laptop:    0,
				Desktop:   0,
				TotalCopy: 1,
			},
		},
		{
			desc: "Test with multiple users",
			param1: map[string][]model.Row{

				"1091": []model.Row{
					{
						Id:            "15",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "LAPTOP",
						Comment:       "Exported from System A",
					},
					{
						Id:            "1",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "LAPTOP",
						Comment:       "Exported from System A",
					},
					{
						Id:            "2",
						UserId:        "1091",
						ApplicationId: "1",
						Type:          "DESKTOP",
						Comment:       "Exported from System B",
					},
				},
				"2003": []model.Row{
					{
						Id:            "15",
						UserId:        "2003",
						ApplicationId: "1",
						Type:          "LAPTOP",
						Comment:       "Exported from System A",
					},
					{
						Id:            "2",
						UserId:        "2003",
						ApplicationId: "1",
						Type:          "DESKTOP",
						Comment:       "Exported from System B",
					},
				},
				"31": []model.Row{
					{
						Id:            "15",
						UserId:        "31",
						ApplicationId: "1",
						Type:          "LAPTOP",
						Comment:       "Exported from System A",
					},
					{
						Id:            "1",
						UserId:        "31",
						ApplicationId: "1",
						Type:          "DESKTOP",
						Comment:       "Exported from System A",
					},
					{
						Id:            "2",
						UserId:        "31",
						ApplicationId: "1",
						Type:          "DESKTOP",
						Comment:       "Exported from System B",
					},
				},
			},
			expected: model.Result{
				Laptop:    0,
				Desktop:   0,
				TotalCopy: 5,
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
		paraddm1 map[string][]model.Row
		expected []byte
	}{
		{
			desc:     "return total count is 10",
			input:    &inputMock{},
			output:   txt.New(),
			expected: []byte("10"),
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
