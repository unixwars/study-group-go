package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
    z, prev := x, x
    for {
        prev = z
        z = z - ((z*z*z - x) / (3.0*z*z))
        if cmplx.Abs(z-prev) < 1e-6 {
            return z
        }
    }
}

func main() {
	fmt.Println(Cbrt(2))
	fmt.Println(cmplx.Pow(2, 1.0/3))
}
