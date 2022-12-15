package collect

import (
	"aoc20221201/data"
	"strings"
)

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
