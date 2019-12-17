package txt

import (
	"strconv"

	"github.com/shoukoo/csv-parser/pkg/model"
)

type txt struct {
}

func New() *txt {
	return &txt{}
}

func (t *txt) GenerateOutput(r model.Result) ([]byte, error) {
	i := strconv.Itoa(r.TotalCopy)
	return []byte(i), nil
}
