package csv

import "github.com/shoukoo/csv-parser/pkg/model"

func getCSVMockSearchApp1() map[string][]model.Row {
	return map[string][]model.Row{

		"1091": []model.Row{
			{
				Id:            "1",
				UserId:        "1091",
				ApplicationId: "1",
				Type:          "DESKTOP",
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
		"1088": []model.Row{
			{
				Id:            "6",
				UserId:        "1088",
				ApplicationId: "1",
				Type:          "DESKTOP",
				Comment:       "Exported from System B",
			},
		},
		"1081": []model.Row{
			{
				Id:            "7",
				UserId:        "1081",
				ApplicationId: "1",
				Type:          "LAPTOP",
				Comment:       "Exported from System B",
			},
		},
	}
}

func getCSVMockSearchApp3() map[string][]model.Row {
	return map[string][]model.Row{
		"1030": []model.Row{
			{
				Id:            "4",
				UserId:        "1030",
				ApplicationId: "3",
				Type:          "desktop",
				Comment:       "Exported from System A",
			},
			{
				Id:            "5",
				UserId:        "1030",
				ApplicationId: "3",
				Type:          "LAPTOP",
				Comment:       "Exported from System A",
			},
		},
		"1035": []model.Row{
			{
				Id:            "8",
				UserId:        "1035",
				ApplicationId: "3",
				Type:          "LAPTOP",
				Comment:       "Exported from System A",
			},
		},
	}
}
