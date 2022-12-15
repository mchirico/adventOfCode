package data

import (
	"testing"
)

func TestRead(t *testing.T) {
	data, err := Read(Path("./input0"))
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
}
