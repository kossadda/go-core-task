package intersection

func Find(first, second []int) (bool, []int) {
	var s []int
	m := make(map[int]struct{})
	for _, key := range second {
		m[key] = struct{}{}
	}

	for _, key := range first {
		if _, exist := m[key]; exist {
			s = append(s, key)
		}
	}

	return len(s) > 0, s
}
