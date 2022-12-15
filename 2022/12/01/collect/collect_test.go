package collect

import (
	"testing"
)

func TestRead(t *testing.T) {
	data, err := Read()
	if err != nil {
		t.Error(err)
	}
	count := 0
	blanks := 0
	for _, v := range data {
		count += 1
		if v == "" {
			blanks += 1
		}
	}
	if count != 2245 && blanks != 247 {
		t.Error("unexpected data")
	}

}

func TestPackage(t *testing.T) {
	data, err := Package()
	if err != nil {
		t.Error(err)
	}
	if len(data) != 248 {
		t.Error("unexpected data")
	}
}
