package utils

import (
	"bytes"
	"github.com/satori/go.uuid"
	"strings"
)

func GetUuid() (IdString string) {
	Id := uuid.NewV4().String()
	d := strings.Split(Id, "-")
	var buffer = bytes.Buffer{}
	for i := 0; i < len(d); i++ {
		buffer.WriteString(d[i])
	}
	IdString = buffer.String()
	return
}
