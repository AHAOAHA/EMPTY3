package MyStrings

func index(s, str string, n int) int {	//从s的n位置开始，匹配字符串，返回字符串位置
	runes := []rune(s)	//将字符串转换成rune类型的切片
	runes = runes[n:]	//将切片从n的位置进行切片
	return MyIndex(string(runes), str)
}

//将字符串中前n个字符中匹配old的字符串替换为new字符串，并返回一个新字符串
func MyReplace(str, old, new string, n int) string {
	if str == "" || old == "" {
		panic("str || old null")
	}

	if n == -1 {
		n = len(str)
	}

	var newstr string	//保存返回的字符串
	var nowpos int	//标记实时位置
	var modpos int	//标记修改位置
	var poscount int
	runes := []rune(str)	//将原来的字符串进行切片
	for nowpos = 0; nowpos < n; nowpos = modpos + len(old) {
		poscount = modpos
		modpos = index(str, old, nowpos)
		if modpos == -1 {	//未匹配到old字符串
			var tmpstr string = string(runes[nowpos:])
			newstr = newstr + tmpstr
			return newstr
		}

		modpos = modpos + nowpos

		var tmpstr string = string(runes[nowpos:modpos])
		newstr = newstr + tmpstr	//将修改位置的前一部分添加进newstr
		newstr = newstr + new	//将new添加进newstr
		poscount = modpos - poscount
	}
	return newstr
}
