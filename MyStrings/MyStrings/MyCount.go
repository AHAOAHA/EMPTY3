package MyStrings

//统计str在s中出现的次数
func MyCount(s, str string) int {
	var startpos int
	var count int = 0
	var prevpos int
	for startpos = 0; startpos < len(s); startpos = startpos + len(str) {
		prevpos = startpos
		startpos = index(s, str, startpos)	//该函数定于于本包中的某位置
		if startpos == -1 {	//未找到对应字符串
			break
		}
		startpos = prevpos + startpos

		count++
	}

	return count
}
