package main

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0
	fmt.Println(z)
	for i := 0; i < 10; i++ {
		//z -= (z*z - x) / (2 * z)
		zDiff := (z*z - x) / (2 * z)
		if zDiff < 0.00000000000005 && zDiff > -0.00000000000005 {
			fmt.Printf("zDiff: %g, last z was %g\n", zDiff, z)
			return z
		}
		z -= zDiff
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))

}
