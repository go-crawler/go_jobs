package uuid

import (
	"github.com/satori/go.uuid"
)

func GetUUID() string {
	uid := uuid.Must(uuid.NewV4())
	return uid.String()
}
