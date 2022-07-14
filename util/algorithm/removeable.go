package algorithm

type removeableKey interface{}

type removeable interface {
	Key() removeableKey
}

//Non completed, please don't use this function
func removeSliceKey(sl []removeable, removeables []removeableKey) []removeable {
	rs := make([]removeable, 0, len(sl))
	previousIndex := 0
	for i, b := range sl {
		k := b.Key()

		exist := false
		for _, id := range removeables {
			if k == id {
				exist = true
				break
			}
		}

		if exist {
			if i > 0 && i >= previousIndex {
				rs = append(rs, sl[previousIndex:i]...)
			}
			previousIndex = i + 1
		}
		if i == len(sl)-1 {
			rs = append(rs, sl[previousIndex:]...)
		}
	}

	return rs
}
