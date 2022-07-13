package processor

import (
	"encoding/csv"
	"io"
	"log"
	"sync"

	a "awesomeProject1/csv"
)

type Process func(json a.JSON) a.JSON

type R *csv.Reader

type Processor struct {
	lg *log.Logger
	r  Reader
	w  Writer
	p  []Process
}

func New(lg *log.Logger, r Reader, w Writer, p ...Process) Processor {
	return Processor{
		lg: lg, r: r, w: w, p: p,
	}
}

func (p Processor) Start() error {
	dataChannel := make([]chan a.JSON, len(p.p)+1)

	for i := 0; i < len(p.p)+1; i++ {
		dataChannel[i] = make(chan a.JSON)
	}

	// executing different processes
	for i := range p.p {
		go func(i int) {
			defer close(dataChannel[i+1])
			for v := range dataChannel[i] {
				dataChannel[i+1] <- p.p[i](v)
			}
		}(i)
	}

	wg := sync.WaitGroup{}

	// writing data to output file
	go func(wg *sync.WaitGroup) error {
		defer wg.Done()

		for v := range dataChannel[len(dataChannel)-1] {
			err := p.w.WriteCSV(v)
			if err != nil {
				return err
				p.lg.Fatal(err)
			}
		}
		return nil
	}(&wg)

	for {
		if data, err := p.r.ReadCSV(); err == io.EOF {
			break
		} else if err != nil {
			return err
		} else {
			dataChannel[0] <- data
		}
	}

	close(dataChannel[0])

	wg.Add(1)
	wg.Wait()

	return nil
}
