package main

func intersection(sl1, sl2 []int) (bool, []int) {
	out := []int{}
	m := map[int]struct{}{}

	for _, n := range sl1 {
		m[n] = struct{}{}
	}

	for _, n := range sl2 {
		if _, ok := m[n]; ok {
			out = append(out, n)
		}
	}

	if len(out) != 0 {
		return true, out
	}
	return false, []int{}
}
