package MyStrings

//判断s字符串是否以prefix作为开头
func MyHasPrefix(s, prefix string) bool {
	if len(s) <= 0 || len(prefix) <= 0 {
		panic("str null")
	}

	if len(prefix) > len(s) {
		return false
	}

	for k, _ := range prefix {
		if s[k] != prefix[k] {
			return false
		}
	}
	return true
}