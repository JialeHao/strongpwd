package generator

import (
	"bytes"
	"math/rand"
	"time"
)

func GenerateComplexPassword(length int) bytes.Buffer {
	if length > 1024 {
		panic("The maximum length of the generated password is 1024")
	}
	var buf bytes.Buffer
	rand.Seed(time.Now().UnixNano())
	for buf.Len() < length {
		tmpnum := rand.Intn(93) + 33
		buf.WriteByte(byte(tmpnum))
	}
	return buf
}