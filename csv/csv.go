package csv

import (
	"encoding/csv"
	"io"
)

type JSON map[string]string

type Reader struct {
	*csv.Reader
}

type Writer struct {
	*csv.Writer
}

func NewCsvReader(r io.Reader) Reader {
	return Reader{csv.NewReader(r)}
}

func NewCsvWriter(w io.Writer) Writer {
	return Writer{csv.NewWriter(w)}
}

func (reader *Reader) ReadCSV() (JSON, error) {
	row := JSON{}
	header := []string{"id", "name", "age", "phoneNumber"}

	record, err := reader.Read()
	if err == io.EOF || err != nil {
		return row, err
	} else {
		for i := range header {
			row[header[i]] = record[i]
		}
	}

	return row, nil
}

func (writer *Writer) WriteCSV(input JSON) error {
	defer writer.Flush()
	var res []string
	header := []string{"id", "name", "age", "phoneNumber", "Signature"}
	for _, v := range header {
		res = append(res, input[v])
	}

	return writer.Write(res)
}
