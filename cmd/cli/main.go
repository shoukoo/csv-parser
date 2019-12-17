package main

import (
	"flag"
	"flexrea/pkg/handler"
	"flexrea/pkg/input/csv"
	"flexrea/pkg/output/txt"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {

	// Flag allow the user to pass in application ID
	// and custom csv file into the program.
	// default application ID 374
	// default CSV file is sample-large.csv
	desc := `
Usage of flexera:
  flexera [flags] [app_id]
`
	var file = flag.String("file", "sample-large.csv", "Load a custom csv file")
	var help = flag.Bool("help", false, "Print help and exit")
	flag.Parse()

	if *help {
		fmt.Printf("%v", desc)
		fmt.Println("Flags:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	appId := flag.Arg(0)
	if flag.Arg(0) == "" {
		appId = "374"
	}

	// Use CSV as input
	input := csv.New(*file)
	// Output the result as text
	output := txt.New()
	service := handler.New(input, output, appId)

	b, err := service.GetCopyCount()
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", string(b))
	return nil
}
