package main

import (
	"testing"

	"github.com/shoukoo/csv-parser/pkg/handler"
	"github.com/shoukoo/csv-parser/pkg/input/csv"
	"github.com/shoukoo/csv-parser/pkg/output/txt"
)

func BenchmarkSampleLargeFile(b *testing.B) {
	input := csv.New("sample-large.csv")
	output := txt.New()
	service := handler.New(input, output, "374")
	for i := 0; i < b.N; i++ {
		service.GetCopyCount()
	}
}

func BenchmarkSampleSmallFile(b *testing.B) {
	input := csv.New("sample-small.csv")
	output := txt.New()
	service := handler.New(input, output, "374")
	for i := 0; i < b.N; i++ {
		service.GetCopyCount()
	}
}
