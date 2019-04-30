package MyStrings

//判断字符串的包含关系，判断substr
func MyContains(s, substr string)bool {
	if len(s) <= 0 || len(substr) <= 0 {
		panic("str null")
	}

	if i:= MyIndex(s, substr); i == -1 {
		return false
	}
	return true
}
