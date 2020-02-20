package main

/*

int empty() {
	// error code
	return 0;
}

void regular_empty() {}

*/
import "C"

import "github.com/LaevusDexter/asmcgocall"

var emptyAsmcgocall = func() (result func()) {
	asmcgocall.Register(C.empty, &result)

	return
}()

func emptyCgocall() {
	C.regular_empty()
}
