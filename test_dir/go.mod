module test_dir

go 1.18

require github.com/shashankdev81/demo/mathlib v0.1.3

require github.com/shashankdev81/demo/mathlib_new v1.1.6

replace github.com/shashankdev81/demo/mathlib v0.1.3 => ../integers

replace github.com/shashankdev81/demo/mathlib_new v1.1.6 => ../new_integers
