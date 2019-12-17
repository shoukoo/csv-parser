package txt

import (
	"flexrea/pkg/model"
	"strconv"
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
