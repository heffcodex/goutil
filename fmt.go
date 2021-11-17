package goutil

import (
	"math"
	"strconv"
	"strings"
)

func DelimThousands(v int) string {
	if v == math.MinInt64 {
		return "-9 223 372 036 854 775 808"
	}

	sign := ""

	if v < 0 {
		sign = "-"
		v = 0 - v
	}

	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for v > 999 {
		parts[j] = strconv.FormatInt(int64(v%1000), 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		v = v / 1000
		j--
	}
	parts[j] = strconv.Itoa(v)
	return sign + strings.Join(parts[j:], " ")
}
