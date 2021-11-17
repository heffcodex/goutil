package goutil

import "math"

func TruePercent(conds ...bool) int {
	p1 := float64(100) / float64(len(conds))
	trues := 0

	for _, c := range conds {
		if c {
			trues++
		}
	}

	return int(math.Ceil(p1 * float64(trues)))
}

func IsAllFalse(conds ...bool) bool {
	for _, cond := range conds {
		if cond {
			return false
		}
	}

	return true
}
