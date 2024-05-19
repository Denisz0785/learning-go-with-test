package injiection

import (
	"fmt"
	"io"
)

func Greet(w io.Writer, name string) {
	fmt.Fprint(w, "Hello, "+name)

}
