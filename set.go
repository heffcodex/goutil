package goutil

func IntSet(arr []int) []int {
	m := make(map[int]interface{}, len(arr))
	set := make([]int, 0, len(arr))

	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v] = nil
			set = append(set, v)
		}
	}

	return set
}

func StringSet(arr []string) []string {
	m := make(map[string]interface{}, len(arr))
	set := make([]string, 0, len(arr))

	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v] = nil
			set = append(set, v)
		}
	}

	return set
}
