package algorithm

func IntersectionArrays(arrays ...[]interface{}) []interface{} {
	if len(arrays) == 0 {
		return []interface{}{}
	}
	if len(arrays) == 1 {
		return arrays[0]
	}

	firstArray := arrays[0]
	hashing := map[interface{}]bool{}
	for _, el := range firstArray {
		hashing[el] = true
	}

	for _, array := range arrays[1:] {
		if len(array) == 0 {
			return []interface{}{}
		}

		intersectionMap := map[interface{}]bool{}
		for _, el := range array {
			if _, ok := hashing[el]; ok {
				intersectionMap[el] = true
			}
		}
		hashing = intersectionMap

		if len(hashing) == 0 {
			return []interface{}{}
		}
	}

	result := []interface{}{}
	for key, _ := range hashing {
		result = append(result, key)
	}

	return result
}

func SearchIndexSliceString(s string, slice []string) int {
	for idx, el := range slice {
		if el == s {
			return idx
		}
	}

	return -1
}
