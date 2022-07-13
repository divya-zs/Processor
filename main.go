package main

import (
	a "awesomeProject1/csv"
	"awesomeProject1/functions"
	"log"
	"os"

	p "awesomeProject1/processor"
)

func main() {
	lg := log.Logger{}
	inFile, err := os.Open("in.csv")
	if err != nil {
		lg.Fatal(err)
	}

	outFile, err := os.OpenFile("out.csv", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		lg.Fatal(err)
	}

	reader := a.NewCsvReader(inFile)
	writer := a.NewCsvWriter(outFile)

	processor := p.New(&lg, &reader, &writer, functions.Hashing, functions.MessageDigest)

	err = processor.Start()
	if err != nil {
		lg.Fatal("Something went wrong: ", err)
	}

}
