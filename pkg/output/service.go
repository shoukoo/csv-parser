package output

import "github.com/shoukoo/csv-parser/pkg/model"

type Output interface {
	GenerateOutput(model.Result) ([]byte, error)
}
