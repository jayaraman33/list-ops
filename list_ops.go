package listops

type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

//Foldl
func (list IntList) Foldl(fn binFunc, initial int) int {
	res := initial
	for _, v := range list {
		res = fn(res, v)
	}
	return res
}

//Foldr
func (list IntList) Foldr(fn binFunc, initial int) int {
	res := initial
	for i := len(list) - 1; i >= 0; i-- {
		res = fn(list[i], res)
	}
	return res
}

//Filter
func (list IntList) Filter(fn predFunc) IntList {
	res := make(IntList, 0)
	for _, v := range list {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}

//Length
func (list IntList) Length() int {
	return len(list)
}

//Map
func (list IntList) Map(fn unaryFunc) IntList {
	for i := 0; i < len(list); i++ {
		list[i] = fn(list[i])
	}
	return list
}

//Reverse
func (list IntList) Reverse() IntList {
	for i := 0; i < len(list)/2; i++ {
		list[i], list[len(list)-i-1] = list[len(list)-i-1], list[i]
	}
	return list
}

//Append
func (list IntList) Append(Append IntList) IntList {
	for _, v := range Append {
		list = append(list, v)
	}
	return list
}

//Concat
func (list IntList) Concat(Concat []IntList) IntList {
	for _, l := range Concat {
		for _, v := range l {
			list = append(list, v)
		}
	}
	return list
}
