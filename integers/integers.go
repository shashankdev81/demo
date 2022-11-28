package integers

import (
	"fmt"
	"io"
)

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) int {
	return a / b
}

func Print(text string, w io.Writer) {
	fmt.Fprint(w, text)
}
