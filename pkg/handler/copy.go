package handler

import (
	"flexrea/pkg/input"
	"flexrea/pkg/model"
	"flexrea/pkg/output"
	"strings"
)

type Copy interface {
	GetCopyCount() ([]byte, error)
}

type service struct {
	Input  input.Input
	Output output.Output
	AppId  string
}

func New(i input.Input, o output.Output, appId string) Copy {
	return &service{
		Input:  i,
		Output: o,
		AppId:  appId,
	}
}

// GetCopyCount glues input, output and business logic together
// then display the result
func (s *service) GetCopyCount() ([]byte, error) {
	rows, err := s.Input.Read(s.AppId)
	if err != nil {
		return nil, err
	}

	results, err := calculateCopyCount(rows)
	if err != nil {
		return nil, err
	}

	b, err := s.Output.GenerateOutput(results)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// calculateCopyCount calculate the minimum number of copies of the application the company must purchase
func calculateCopyCount(rows []model.Row) (model.Result, error) {
	var results model.Result
	for _, row := range rows {
		switch strings.ToLower(row.Type) {
		case "desktop":
			results.Desktop += 1
		case "laptop":
			results.Laptop += 1
		}
	}

	total := (results.Desktop + results.Laptop) - results.Laptop
	results.TotalCopy = total

	return results, nil
}
