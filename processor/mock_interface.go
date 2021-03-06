// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package processor is a generated GoMock package.
package processor

import (
	csv "awesomeProject1/csv"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// ReadCSV mocks base method.
func (m *MockReader) ReadCSV() (csv.JSON, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCSV")
	ret0, _ := ret[0].(csv.JSON)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadCSV indicates an expected call of ReadCSV.
func (mr *MockReaderMockRecorder) ReadCSV() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCSV", reflect.TypeOf((*MockReader)(nil).ReadCSV))
}

// MockWriter is a mock of Writer interface.
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter.
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance.
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// WriteCSV mocks base method.
func (m *MockWriter) WriteCSV(json csv.JSON) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteCSV", json)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteCSV indicates an expected call of WriteCSV.
func (mr *MockWriterMockRecorder) WriteCSV(json interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteCSV", reflect.TypeOf((*MockWriter)(nil).WriteCSV), json)
}
