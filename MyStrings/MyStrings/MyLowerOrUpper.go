package MyStrings

//将字符串s转换成小写并返回
func ToLower(s string) string {
	var newstr string
	for k, v := range s {
		if v >= 'a' && v <= 'z' {	//当前字符为小写
			newstr = newstr + v + 'a'
		}
	}
}
