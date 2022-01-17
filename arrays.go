package goutil

import (
	"math"
	"strconv"
	"strings"
)

func StrInArray(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}

	return false
}

func StrArrayUnique(arr []string) []string {
	m := make(map[string]interface{}, len(arr))

	for _, s := range arr {
		m[s] = nil
	}

	res := make([]string, 0, len(m))

	for s := range m {
		res = append(res, s)
	}

	return res
}

func StrContainsInArrayInsensitive(arr []string, value string) bool {
	if len(arr) == 0 || len(value) == 0 {
		return false
	}

	value = strings.ToLower(value)

	for _, v := range arr {
		if strings.Contains(strings.ToLower(v), value) {
			return true
		}
	}

	return false
}

func StrArrayIncludesOtherInsensitive(inclusion []string, arr []string) bool {
	if len(inclusion) == 0 || len(arr) == 0 {
		return false
	}

	m := make(map[string]interface{})

	for _, s := range arr {
		m[s] = nil
	}

	count := 0

	for _, s := range inclusion {
		if _, ok := m[s]; ok {
			count++
		}
	}

	return count == len(inclusion)
}

func StrArrayIntersect(a, b []string) []string {
	if len(a) == 0 || len(b) == 0 {
		return []string{}
	}

	res := make([]string, 0, int(math.Min(float64(len(a)), float64(len(b)))))
	m := make(map[string]interface{})

	for _, v := range a {
		m[v] = nil
	}

	for _, v := range b {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
	}

	return res
}

func StrArrayFilter(arr []string, fn func(s string) bool) (res []string, excluded uint) {
	res = make([]string, 0)

	for _, s := range arr {
		if fn(s) {
			res = append(res, s)
		} else {
			excluded++
		}
	}

	return
}

func IntArrFromStringArr(arr []string) []int {
	res := make([]int, 0, len(arr))

	for _, s := range arr {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			continue
		}

		res = append(res, int(i))
	}

	return res
}

func IntArrayIntersect(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return []int{}
	}

	res := make([]int, 0, int(math.Min(float64(len(a)), float64(len(b)))))
	m := make(map[int]interface{})

	for _, v := range a {
		m[v] = nil
	}

	for _, v := range b {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
	}

	return res
}
