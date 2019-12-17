package output

import "flexrea/pkg/model"

type Output interface {
	GenerateOutput(model.Result) ([]byte, error)
}
