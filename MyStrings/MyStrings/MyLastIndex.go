package MyStrings

func MyLastIndex(s, str string)int {
	if len(s) <= 0 || len(str) <= 0 {
		panic("str null")
	}

	var res string = reverse(s)
	var restr string = reverse(str)
	var pos int = MyIndex(res, restr)
	return len(s) - pos - 1
}
