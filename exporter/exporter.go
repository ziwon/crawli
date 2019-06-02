package exporter

import (
	"encoding/csv"
	"fmt"
	"io"
)

type Writer struct {
	Writer *csv.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		Writer: csv.NewWriter(w),
	}
}

func (w *Writer) WriteHeader(columns []string) {
	fmt.Println(columns)
	w.Writer.Write(columns)
}

func (w *Writer) Write(header []string, record map[string][]string) {
	for key := range record {
		row := make([]string, 0)
		row = append(row, key)
		row = append(row, record[key]...)
		fmt.Println(row)
		w.Writer.Write(row)
	}
}
