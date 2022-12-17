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

func MockPackage() ([][]string, error) {
	return [][]string{
		{"0", "1", "2"},
		{"0", "1", "1"},
		{},
		{"0", "1", "2", "3"},
	}, nil
}

func TestSummerize(t *testing.T) {
	s := Stats{}
	err := s.NewSum(MockPackage)
	if err != nil {
		t.Error(err)
	}
	if s.Max() != 6 {
		t.Error("unexpected max")
	}
	if s.Sum(0) != 3 {
		t.Error("unexpected sum")
	}
	if s.Sum(1) != 2 {
		t.Error("unexpected sum")
	}
	if s.Sum(s.ID()) != 6 {
		t.Error("unexpected sum")
	}
}
