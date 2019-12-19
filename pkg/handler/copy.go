package handler

import (
	"strings"

	"github.com/shoukoo/csv-parser/pkg/input"
	"github.com/shoukoo/csv-parser/pkg/model"
	"github.com/shoukoo/csv-parser/pkg/output"
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
func calculateCopyCount(rows map[string][]model.Row) (model.Result, error) {

	var total int
	for _, rows := range rows {
		co := model.Result{}

		for _, r := range rows {
			t := strings.ToLower(r.Type)
			if t == "desktop" {
				co.Desktop += 1
			}

			if t == "laptop" {
				co.Laptop += 1
			}
		}

		// Laptop has more than desktop
		if co.Laptop > co.Desktop {
			t := co.Laptop + co.Desktop
			var leftover int
			// Make it even number
			if t%2 != 0 {
				t = t - 1
				leftover = 1
			}

			total = total + (t/2 + leftover)
		}

		// Laptop is the same quantity as desktop
		if co.Laptop == co.Desktop {
			total = total + co.Laptop
		}

		// Laptop is less than desktop
		if co.Laptop < co.Desktop {
			t := co.Laptop + co.Desktop
			total = total + (t - co.Laptop)
		}
	}

	return model.Result{TotalCopy: total}, nil
}
