package test_dir

import (
	"fmt"
	integers "github.com/shashankdev81/demo/mathlib"
	new_integers "github.com/shashankdev81/demo/mathlib_new"
)

func TestFunc() {
	ans := integers.Divide(5, 1)
	new_integers.NewDivide()
	fmt.Println(ans)
}
