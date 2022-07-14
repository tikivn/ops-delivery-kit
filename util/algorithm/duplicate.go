package algorithm

func RetreiveDuplicateStrings(a, b []string) (duplicate []string) {
	al := len(a)
	bl := len(b)
	if al == 0 || bl == 0 {
		return []string{}
	}

	capacity := al
	if bl > capacity {
		capacity = bl
	}

	duplicate = make([]string, 0, capacity)
	abMap := make(map[string]struct{})
	for _, s := range a {
		abMap[s] = struct{}{}
	}

	for _, s := range b {
		if _, exist := abMap[s]; exist {
			duplicate = append(duplicate, s)
		}
	}
	return duplicate
}

func DifferentStrings(source, object []string) (different []string) {
	al := len(source)
	if al == 0 {
		return []string{}
	}

	capacity := al

	different = make([]string, 0, capacity)
	abMap := make(map[string]struct{})
	for _, s := range object {
		abMap[s] = struct{}{}
	}

	for _, s := range source {
		if _, exist := abMap[s]; !exist {
			different = append(different, s)
		}
	}
	return different
}
