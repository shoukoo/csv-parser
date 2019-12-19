package input

import "github.com/shoukoo/csv-parser/pkg/model"

type Input interface {
	Read(appId string) (map[string][]model.Row, error)
	IsDuplicated(rows []model.Row, machineId string) bool
}
