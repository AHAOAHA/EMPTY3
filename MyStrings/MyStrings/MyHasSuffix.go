package MyStrings


//go反转字符串
func reverse(s string)string {
	if len(s) <= 0 {
		panic("str null")
	}

	runes := []rune(s)
	for begin, end := 0, len(s) - 1; begin < end ;begin, end = begin + 1, end - 1 {
		runes[begin], runes[end] = runes[end], runes[begin]
	}
	return string(runes)
}

//同一个包中的public函数可以互相调用，不需要加上包名，也不需要import当前包
func MyHasSuffix(s, suffix string) bool {
	if len(s) <= 0 || len(suffix) <= 0 {
		panic("str null")
	}

	var restr string = reverse(s)
	var resuffix string = reverse(suffix)
	return MyHasPrefix(restr, resuffix)
}
