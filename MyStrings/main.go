package main
import "fmt"
import ."go_code/MyStrings/MyStrings"

func main() {
	str := "helddddlo worddddddddddldddddddddddddd"
	prefix := "llo"
	fmt.Println(MyHasPrefix(str, prefix))
	fmt.Println(MyHasSuffix(str, "ld"))
	fmt.Println(MyContains(str, "lilo"))
	fmt.Println(MyIndex(str, "ld"))
	fmt.Println(MyLastIndex(str, "d"))
	fmt.Println(MyReplace(str, "d", "l", -1))
	var str1 string = MyReplace(str, "d", "l", -1)
	fmt.Println(MyCount(str1, "l"))
	fmt.Println(MyRepeat("ahaoo", 10))

}
