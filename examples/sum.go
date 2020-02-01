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

int empty() {
	return 0;
}

void regular_empty() {}
*/
import "C"

import (
	"fmt"
	"github.com/LaevusDexter/asmcgocall"
)

func main() {
	var sum func(C.int, C.int, C.int, C.int, C.int) C.int

	asmcgocall.Register(C.sum, &sum)

	fmt.Println("result: ", sum(1, 2, 3, 4, 5))

	fmt.Println(sumAsmcgocall(1, 2, 3, 4, 5))

	emptyAsmcgocall()
}

var sumAsmcgocall = func() (result func(C.int, C.int, C.int, C.int, C.int) C.int) {
	asmcgocall.Register(C.sum, &result)

	return
}()

var emptyAsmcgocall = func() (result func()) {
	asmcgocall.Register(C.empty, &result)

	return
}()


func sumCgocall(p0, p1, p2, p3, p4 C.int) C.int {
	return C.regular_sum(p0, p1, p2, p3, p4)
}

func emptyCgocall() {
	C.regular_empty()
}


