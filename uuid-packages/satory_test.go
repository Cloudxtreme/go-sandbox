package satory_test

import (
	"testing"

	"github.com/satori/go.uuid"
)

func TestUUID(t *testing.T) {
	for i := 0; i < 1000; i++ {
		uuid := uuid.NewV4().String()
		t.Logf("%d: uuid: %v:", i, uuid)
	}
	t.Logf("Done")
}
