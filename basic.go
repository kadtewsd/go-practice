package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func main() {
	// Println is exported because it starts with Upper case "Println"
	fmt.Println("Hello, world ")
	fmt.Printf("My favorite number is %d \n", rand.Intn(10))
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("1 + 2 = %d\n", add(1, 2))
	// 型推論は : で代入
	x, y := swap("x", "y")
	fmt.Printf("x is swapped by %s, y is swapped by %s\n", x, y)
	first, second := split(17)
	fmt.Printf("first is %d second is %d\n", first, second)

	// int は 32 bit マシンなら int32, 64 bit ならば int64
	// 初期値なしは 0
	var i int
	var j, k = 1, 2
	var c, python, java bool
	fmt.Println(i, j, k, c, python, java)

	fmt.Printf("Type: %T ToBe Value : %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T MaxInt Value : %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T z Value : %v\n", z, z)

	var f float64 = math.Sqrt(float64(3*4 + 4*4))
	// CAST
	var z = uint(f)
	fmt.Printf("float value is %f, uint value is %d\n", f, z)

	v := 42
	fmt.Printf("%d value Type is infererred [%T]", v, v)

	// compile error
	// const Truth := true
	const Truth = true
	fmt.Println(Truth)
	fmt.Printf("Small value is calculted %d\n", needInt(Small))
	fmt.Printf("Small value is calculated %f\n", needFloat(Small))
	fmt.Printf("Big value is calculated %f\n", needFloat(Big))

}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i) // 複素数
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x + 0.1
}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
