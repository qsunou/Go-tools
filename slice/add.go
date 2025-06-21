package slice

import (
	"github.com/qsunou/Go-tools/internal/slice"
)

func Add[Src any](src []Src, element Src, index int) ([]Src, error) {
	res, err := slice.Add[Src](src, element, index)
	return res, err
}
