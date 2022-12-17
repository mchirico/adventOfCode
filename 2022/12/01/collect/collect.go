package collect

import (
	"aoc20221201/data"
	"strconv"
	"strings"
)

type Summer struct {
	count int
	sum   int
	raw   []string
}

type Stats struct {
	max   int
	id    int
	total int
	data  map[int]Summer
}

func (s *Stats) ID() int {
	return s.id
}
func (s *Stats) Max() int {
	return s.max
}

func (s *Stats) Sum(i int) int {
	return s.data[i].sum
}
func (s *Stats) NewSum(f func() ([][]string, error)) error {
	if s.data == nil {
		s.data = map[int]Summer{}
	}

	data, err := f()
	if err != nil {
		return err
	}
	s.total = len(data)
	for i, v := range data {
		sum := 0
		for _, vv := range v {
			value, err := strconv.Atoi(vv)
			if err != nil {
				return err
			}
			sum += value
		}
		if s.max < sum {
			s.max = sum
		}
		s.id = i
		s.data[i] = Summer{count: len(v), raw: v, sum: sum}

	}
	return nil

}

func Read() ([]string, error) {
	raw, err := data.Read(data.Path("./input0"))
	if err != nil {
		return nil, err
	}
	return strings.Split(string(raw), "\n"), nil
}

func Package() ([][]string, error) {
	out := [][]string{}
	data, err := Read()
	if err != nil {
		return nil, err
	}
	sub := []string{}
	for _, line := range data {
		if line == "" {
			out = append(out, sub)
			sub = []string{}
		} else {
			sub = append(sub, line)
		}
	}
	out = append(out, sub)
	return out, nil
}
