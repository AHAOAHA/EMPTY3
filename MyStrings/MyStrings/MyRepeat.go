package MyStrings

//将string字符串重复count次并组织成一个新的字符串返回
func MyRepeat(s string, count int) string {
	var newstr string
	for i := 0; i < count; i++ {
		newstr = newstr + s
	}

	return newstr
}
