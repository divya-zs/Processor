package functions

import (
	a "awesomeProject1/csv"
	"crypto/sha1"
	"encoding/hex"
	"io"
)

func hash(ip string) string {
	h := sha1.New()
	io.WriteString(h, ip)

	return hex.EncodeToString(h.Sum(nil))
}
func Hashing(input a.JSON) a.JSON {
	input["phoneNumber"] = hash(input["phoneNumber"])

	return input
}

func MessageDigest(input a.JSON) a.JSON {
	out := input
	var msgString string
	header := []string{"id", "name", "age", "phoneNumber"}
	for _, v := range header {
		msgString += input[v]
	}

	out["Signature"] = msgString

	return out
}
