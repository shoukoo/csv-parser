package csv

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"

	"github.com/shoukoo/csv-parser/pkg/model"
)

type csvInput struct {
	file string
}

func New(file string) *csvInput {
	return &csvInput{
		file: file,
	}
}

// Read reads the csv file and filter out the duplicate records afterward
func (c *csvInput) Read(appId string) ([]model.Row, error) {
	file, err := os.Open(c.file)
	if err != nil {
		log.Fatal(err)
	}
	parser := csv.NewReader(file)
	records := make(map[string][]model.Row)
	for {
		// Read the csv file line by line
		v, err := parser.Read()
		if err == io.EOF {
			break
		}
		if v[2] == appId {

			row := model.Row{
				Id:            v[0],
				UserId:        v[1],
				ApplicationId: v[2],
				Type:          v[3],
				Comment:       v[4],
			}

			if _, ok := records[v[1]]; !ok {
				records[v[1]] = []model.Row{}
			}

			userMachines := records[v[1]]
			if !c.IsDuplicated(userMachines, v[0]) {
				records[v[1]] = append(userMachines, row)
			}
		}
	}

	var rows []model.Row
	for _, v := range records {
		rows = append(rows, v...)
	}

	if len(rows) == 0 {
		return nil, errors.New("No record found")
	}
	return rows, nil
}

// IsDuplicated returns true if user ID and app ID both appears in one of row in the rows array
func (c *csvInput) IsDuplicated(rows []model.Row, machineId string) bool {
	var found bool
	for _, v := range rows {
		if v.Id == machineId {
			found = true
			break
		}
	}
	return found
}
