package MyStrings

//返回str中第一个匹配s的首字符的位置
func MyIndex(s, str string) int {
	//s&str都不能为空
	if len(s) <= 0 || len(str) <= 0 {
		panic("str null")
	}

	var startpos int	//在s字符串中起始的比较位置
	var compos int	//s字符串中，开始与str比较的位置
	var strpos int	//表示str对应位置
	for startpos = 0;startpos < len(s);startpos++ {
		if s[startpos] == str[0] {	//找到第一个匹配字符，保存该位置，开始向后比较
			compos = startpos
			strpos = 0

			//开始比较s&str
			for ;strpos < len(str); compos, strpos = compos + 1, strpos + 1 {
				if compos >= len(s) {	//判断compos是否越界，越界说明无匹配字符串
					return -1
				}
				if s[compos] != str[strpos] {
					break
				}
			}

			if strpos == len(str)  {	//匹配字符串完成
				return startpos
			}
		}
	}
	return -1
}
