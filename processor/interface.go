package processor

import a "awesomeProject1/csv"

type Reader interface {
	ReadCSV() (a.JSON, error)
}

type Writer interface {
	WriteCSV(json a.JSON) error
}
