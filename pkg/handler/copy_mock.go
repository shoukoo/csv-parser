package handler

import "github.com/shoukoo/csv-parser/pkg/model"

type inputMock struct {
}

func (i *inputMock) Read(appId string) ([]model.Row, error) {

	return []model.Row{

		{
			Id:            "11",
			ApplicationId: "101",
			UserId:        "41",
			Type:          "Laptop",
			Comment:       "Imported from system X",
		},
		{
			Id:            "12",
			ApplicationId: "101",
			UserId:        "24",
			Type:          "Desktop",
			Comment:       "Imported from system X",
		},
		{
			Id:            "13",
			ApplicationId: "101",
			UserId:        "24",
			Type:          "Laptop",
			Comment:       "Imported from system X",
		},
		{
			Id:            "14",
			ApplicationId: "101",
			UserId:        "12",
			Type:          "Desktop",
			Comment:       "Imported from system X",
		},
	}, nil

}

func (s *inputMock) IsDuplicated(rows []model.Row, machineId string) bool {
	return false
}
