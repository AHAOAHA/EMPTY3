package main

/*
#include <stdio.h>
#include <dirent.h>
int hello(char* str) {
	printf("hello world\n");
	printf("%s\n", str);
	return 0;
}

int filepath(char* file) {
	struct dirent **p_dirent = NULL;
	int num = scandir(file, &p_dirent, 0, alphasort);
	int i = 0;
	for(i = 0; i < num; i++) {
		printf("%s\n", p_dirent[i]->d_name);
	}
	return 0;
}
*/
import "C"

func main() {
	//C.hello(C.CString("world"))
	C.filepath(C.CString("/"))

}
