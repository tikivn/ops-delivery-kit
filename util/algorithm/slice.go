package algorithm

func RemoveIndexString(slice []string, index int) (result []string) {
	if index < 0 || index >= len(slice) {
		return slice
	}

	if index == 0 {
		return slice[1:]
	}

	if index == len(slice)-1 {
		return slice[0 : len(slice)-1]
	}

	return append(slice[0:index], slice[index+1:]...)
}

func ReverseStringSlice(s []string) []string {
	result := make([]string, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return result
}
