package handler

import "github.com/shoukoo/csv-parser/pkg/model"

type inputMock struct {
}

func (i *inputMock) Read(appId string) (map[string][]model.Row, error) {

	return map[string][]model.Row{

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
		"5": []model.Row{
			{
				Id:            "155",
				UserId:        "5",
				ApplicationId: "1",
				Type:          "LAPTOP",
				Comment:       "Exported from System A",
			},
			{
				Id:            "152",
				UserId:        "5",
				ApplicationId: "1",
				Type:          "LAPTOP",
				Comment:       "Exported from System A",
			},
			{
				Id:            "153",
				UserId:        "5",
				ApplicationId: "1",
				Type:          "LAPTOP",
				Comment:       "Exported from System A",
			},
			{
				Id:            "15",
				UserId:        "5",
				ApplicationId: "1",
				Type:          "LAPTOP",
				Comment:       "Exported from System A",
			},
			{
				Id:            "1",
				UserId:        "5",
				ApplicationId: "1",
				Type:          "DESKTOP",
				Comment:       "Exported from System A",
			},
			{
				Id:            "2",
				UserId:        "5",
				ApplicationId: "1",
				Type:          "DESKTOP",
				Comment:       "Exported from System B",
			},
		},
		"2215": []model.Row{
			{
				Id:            "2",
				UserId:        "2215",
				ApplicationId: "1",
				Type:          "DESKTOP",
				Comment:       "Exported from System B",
			},
		},
		"111": []model.Row{
			{
				Id:            "2",
				UserId:        "111",
				ApplicationId: "1",
				Type:          "Laptop",
				Comment:       "Exported from System B",
			},
		},
	}, nil

}

func (s *inputMock) IsDuplicated(rows []model.Row, machineId string) bool {
	return false
}
