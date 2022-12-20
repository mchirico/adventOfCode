package collect

import (
	"aoc20221201/data"
	"strconv"
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

func Read() ([]byte, error) {
	return data.Read(data.Path("./input0"))
}

type Packer struct {
	Read   func() ([]byte, error)
	ops    []string
	sep    string // separator
	Modify func([]byte, ...string) ([]string, error)
	sdata  [][]string
}

func (p *Packer) Data() [][]string {
	return p.sdata
}

func (p *Packer) Package() error {
	if p.sdata == nil {
		p.sdata = [][]string{}
	}
	d, err := p.Read()
	if err != nil {
		return err
	}
	list, err := p.Modify(d, p.ops...)
	if err != nil {
		return err
	}
	var sub []string
	for _, line := range list {
		if line == p.sep {
			p.sdata = append(p.sdata, sub)
			sub = []string{}
		} else {
			sub = append(sub, line)
		}
	}
	p.sdata = append(p.sdata, sub)
	return nil
}
