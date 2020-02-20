package main

/*

int ret_value = 666;

int single_ret(int *ret) {

	*ret = ret_value;

	// error code
	return 0;
}

int regular_single_ret() {
	return ret_value;
}

*/
import "C"

import "github.com/LaevusDexter/asmcgocall"

var singleRetExpected = C.ret_value

var singleRetAsmcgocall = func() (result func() C.int) {
	asmcgocall.Register(C.single_ret, &result)

	return
}()

func singleRetCgocall() C.int {
	return C.regular_single_ret()
}
