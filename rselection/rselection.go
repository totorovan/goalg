package rselection

import (
	"math/rand"
	"time"
)

type EmptyErr struct {
	s string
}

func (e EmptyErr) Error() string {
	return e.s
}

func NewEmptyErr(msg string) error {
	return &EmptyErr{msg}
}

func Select(a []int, pos int) (int, error) {
	if len(a) == 0 {
		return 0, NewEmptyErr("Empty slice")
	}
	if len(a) == 1 {
		return a[0], nil
	}
	p := pivot(a)
	i := 1
	for j := 1; j < len(a); j++ {
		if a[j] < p {
			if i != j {
				swap(i, j, a)
			}
			i++
		}
	}
	swap(0, i-1, a)
	switch {
	case i-1 == pos:
		{
			return a[i-1], nil
		}
	case i-1 > pos:
		{
			return Select(a[:i-1], pos)
		}
	default:
		{
			return Select(a[i:], pos-i)
		}
	}
}

func pivot(a []int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	ind := r.Intn(len(a))
	swap(0, ind, a)
	return a[0]
}

func swap(i int, j int, a []int) {
	a[i], a[j] = a[j], a[i]
}
