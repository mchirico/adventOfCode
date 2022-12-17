package collect

import (
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	b, err := Read()
	if err != nil {
		t.Error(err)
	}
	data := strings.Split(string(b), "\n")
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

func mockRead() ([]byte, error) {
	return []byte(`0
1
2

1`), nil
}

func mockModify(b []byte, ops ...string) ([]string, error) {
	return strings.Split(string(b), ops[0]), nil
}

func TestPackage(t *testing.T) {
	p := &Packer{Read: mockRead,
		Modify: mockModify,
		sep:    "",
		ops:    []string{"\n"},
	}
	err := p.Package()
	if err != nil {
		t.Error(err)
	}
	if len(p.Data()) != 2 {
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
