package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	defer fmt.Println("flow control is ended. This is deferred statement")

	fmt.Printf("doFor result is %d\n", doFor())
	sum, sum2 := initStatement();
	fmt.Printf("sum result is %d, sum2 result is %d\n", sum, sum2)
	fmt.Printf("pow result will be pow... %f\n", pow(3, 2, 10))
	fmt.Printf("pow result will be lim... %f\n", pow(3, 3, 20))
	Sqrt(2)
	testSwitch()

	i := 0
	for i < 10 {
		i += 1
		defer fmt.Println(i)
	}
	fmt.Println("main func is done!")
}

func doFor() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func initStatement() (int, int) {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}

	fmt.Println("done")
	return sum, sum2
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("======== %g >= %g ===========\n", v, lim)
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	z1 := 0.0
	for i := 0; i < 10; i++ {
		// x*x = 2
		// x*x -2 = 0
		// f(x) = x*x -2
		z -= (z*z - x) / (2 * z)
		fmt.Printf("current z is %g\n", z)
		if math.Abs(z-z1) < 0.001 {
			fmt.Printf("i is %d\n", i)
			break
		}
		z1 = z
	}
	return z
}

func testSwitch() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("os X\n")
	case "linux":
		fmt.Println("Linux\n")
	default:
		fmt.Printf("os is %s\n", os)
	}
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today is Saturday")
	case today + 1:
		fmt.Println("Tommorow after released go")
	case today + 2:
		fmt.Println("In two days after released go")
	default:
		fmt.Println("Too far away")

	}
	t := time.Now()
	// no condition
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
}
