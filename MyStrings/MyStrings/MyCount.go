package MyStrings

//统计str在s中出现的次数
func MyCount(s, str string) int {
	var startpos int
	var count int = 0
	for startpos = 0; startpos < len(s); startpos = startpos + len(str) {
		startpos = index(s, str, startpos)	//该函数定于于本包中的某位置
		if startpos == -1 {	//未找到对应字符串
			break
		}

		count++
	}

	return count
}
