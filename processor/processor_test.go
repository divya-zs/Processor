package processor

import (
	a "awesomeProject1/csv"
	"awesomeProject1/functions"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"reflect"
	"testing"
)

func TestProcessor_Start(t *testing.T) {
	lg := log.Logger{}

	ctrl := gomock.NewController(t)
	mockReader := NewMockReader(ctrl)
	mockWriter := NewMockWriter(ctrl)

	processor := New(&lg, mockReader, mockWriter, functions.Hashing, functions.MessageDigest)

	testCases := []struct {
		desc      string
		mockCalls []*gomock.Call
		err       error
	}{
		{"success case", []*gomock.Call{
			mockReader.EXPECT().ReadCSV().Return(a.JSON{"id": "1", "name": "abc", "age": "22", "phoneNumber": "90924239139"}, nil),
			mockReader.EXPECT().ReadCSV().Return(a.JSON{}, io.EOF),
			mockWriter.EXPECT().WriteCSV(a.JSON{"id": "1", "name": "abc", "age": "22", "phoneNumber": "aee864024cf886864065e090ebce3435ef13f06f",
				"Signature": "1abc22aee864024cf886864065e090ebce3435ef13f06f"}).Return(nil)}, nil},

		{"error from Read", []*gomock.Call{
			mockReader.EXPECT().ReadCSV().Return(a.JSON{"id": "1", "name": "abc", "age": "22", "phoneNumber": "aee864024cf886864065e090ebce3435ef13f06f",
				"Signature": "1abc22aee864024cf886864065e090ebce3435ef13f06f"}, errors.Error("error in reading file"))}, errors.Error("error in reading file"),
		},
	}

	for i, tc := range testCases {
		err := processor.Start()

		assert.Equal(t, err, tc.err, "TEST[%d] FAILED - %s", i, tc.desc)
	}

}

func TestWriteToCSV(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWriter := NewMockWriter(ctrl)
	tests := []struct {
		desc      string
		data      a.JSON
		err       error
		mockCalls []*gomock.Call
	}{
		{"success case", a.JSON{"id": "1", "name": "abc", "age": "22", "phoneNumber": "90924239139"}, nil, []*gomock.Call{
			mockWriter.EXPECT().WriteCSV(a.JSON{"id": "1", "name": "abc", "age": "22", "phoneNumber": "90924239139"}).Return(nil).AnyTimes(),
		}},
		{"error case", a.JSON{"id": "1", "name": "abc", "age": "22", "phoneNumber": "90924239139"}, errors.Error("some error"), nil},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			input := make(chan a.JSON)

			go func() {
				defer close(input)
				input <- tc.data
			}()

			errChan := Write(mockWriter, input)

			for err := range errChan {
				if !reflect.DeepEqual(err, tc.err) {
					fmt.Printf("TEST FAILED - got %v want %v\n", err, tc.err)
				}
			}
		})
	}
}
