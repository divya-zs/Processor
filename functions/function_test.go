package functions

import (
	a "awesomeProject1/csv"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"reflect"
	"testing"
)

func TestHashing(t *testing.T) {
	var input a.JSON

	input = map[string]string{"phoneNumber": "65342532672"}

	h := sha1.New()
	io.WriteString(h, input["phoneNumber"])
	hashedPhNo := hex.EncodeToString(h.Sum(nil))

	output := Hashing(input)

	input["phoneNumber"] = hashedPhNo
	if !reflect.DeepEqual(output, input) {
		t.Errorf("Expected %v Got %v", input, output)
	}
}

func TestMessageDigest(t *testing.T) {
	var input a.JSON
	input = map[string]string{"phoneNumber": "65342532672"}
	h := sha1.New()
	io.WriteString(h, input["phoneNumber"])
	hashedPhNo := hex.EncodeToString(h.Sum(nil))

	output := MessageDigest(input)

	input["phoneNumber"] = hashedPhNo
	if !reflect.DeepEqual(output, input) {
		t.Errorf("Expected %v Got %v", input, output)
	}
}
