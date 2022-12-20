package data

import (
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	data, err := Read(Path("./input0"))
	if err != nil {
		t.Error(err)
	}
	fields := strings.Fields(string(data))
	lines := strings.Split(string(data), "\n")
	t.Logf("len fields: %v", len(fields))
	t.Logf("len lines: %v", len(lines))

	t.Logf("sample data:\n: %v", data[0:17])
}
