package algorithm

func GetPrefix(a, b string) string {
	var i int
	size1 := len(a)
	size2 := len(b)
	for i < size1 && i < size2 && a[i] == b[i] {
		i++
	}
	return a[:i]
}
