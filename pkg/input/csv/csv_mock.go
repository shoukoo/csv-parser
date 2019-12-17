package csv

import "flexrea/pkg/model"

func getCSVMockSearchApp1() []model.Row {

	return []model.Row{

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
	}
}

func getCSVMockSearchApp3() []model.Row {

	return []model.Row{

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
	}
}
