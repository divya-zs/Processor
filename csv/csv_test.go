package csv

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestReader_ReadCSV(t *testing.T) {
	file, err := os.Open("test.csv")
	if err != nil {
		t.Error(err)
	}

	expectedOutput := JSON{"age": "22", "id": "1", "name": "abc", "phoneNumber": "90924239139"}

	n := NewCsvReader(file)

	op, err := n.ReadCSV()

	assert.Equal(t, expectedOutput, op)
	assert.Equal(t, err, nil)
}

func TestReader_ReadCSV_EOF_Error(t *testing.T) {
	file, err := os.Open("test2.csv")
	if err != nil {
		t.Error(err)
	}

	expectedOutput := JSON{}

	n := NewCsvReader(file)

	op, err := n.ReadCSV()

	assert.Equal(t, expectedOutput, op)
	assert.Equal(t, err, io.EOF)
}
func TestWriter_WriteCSV(t *testing.T) {
	file, err := os.Open("test.csv")
	if err != nil {
		t.Error(err)
	}

	n := NewCsvWriter(file)

	input := map[string]string{"a": "abc"}
	output := n.WriteCSV(input)

	assert.Equal(t, output, nil)
}
