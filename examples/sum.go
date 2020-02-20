package main

/*
int sum(struct {
		int p0;
		int p1;
		int p2;
		int p3;
		int p4;
		int r;
	} *a) {

	//printf("p1 = %d, p2 = %d, p3 = %d, p4 = %d, p5 = %d, r = %d\n", a->p0, a->p1, a->p2, a->p3, a->p4, a->r);

	a->r = a->p0 + a->p1 + a->p2 + a->p3 + a->p4;

	// error code
	return 0;
}

int regular_sum(int p0, int p1, int p2, int p3, int p4) {
	return p0 + p1 + p2 + p3 + p4;
}


*/
import "C"

import (
	"github.com/LaevusDexter/asmcgocall"
)

var sumAsmcgocall = func() (result func(C.int, C.int, C.int, C.int, C.int) C.int) {
	asmcgocall.Register(C.sum, &result)

	return
}()

func sumCgocall(p0, p1, p2, p3, p4 C.int) C.int {
	return C.regular_sum(p0, p1, p2, p3, p4)
}
