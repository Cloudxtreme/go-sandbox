package pborman_test

import (
	"testing"

	"github.com/pborman/uuid"
)

func TestUUID(t *testing.T) {
	for i := 0; i < 1000; i++ {
		uuid := uuid.New()
		t.Logf("%d: uuid: %v:", i, uuid)
	}
	t.Logf("Done")

}
