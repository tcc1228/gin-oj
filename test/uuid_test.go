package test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestGenerateUUID(t *testing.T) {
	println(uuid.NewV4().String())
	println(uuid.NewV4().String())
	println(uuid.NewV4().String())
	println(len(uuid.NewV4().String()))
}
