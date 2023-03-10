package pkg

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

func fn1() {
	strings.Replace("", "", "", 1) //@ diag(`doesn't have side effects`)
	foo(1, 2)                      //@ diag(`doesn't have side effects`)
	baz(1, 2)                      //@ diag(`doesn't have side effects`)
	_, x := baz(1, 2)
	_ = x
	bar(1, 2)
}

func fn2() {
	r, _ := http.NewRequest("GET", "/", nil)
	r.WithContext(context.Background()) //@ diag(`doesn't have side effects`)
}

func foo(a, b int) int        { return a + b }
func baz(a, b int) (int, int) { return a + b, a + b }
func bar(a, b int) int {
	println(a + b)
	return a + b
}

func empty()            {}
func stubPointer() *int { return nil }
func stubInt() int      { return 0 }

func fn3() {
	empty()
	stubPointer()
	stubInt()
}

func fn4() error {
	// Test for https://github.com/dominikh/go-tools/issues/949
	if true {
		return fmt.Errorf("")
	}
	for {
	}
}
