package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	do(1)
	isWeekend(time.Now())

	switch x := math.Sqrt((4)); x {
	case 2:
		fmt.Println("x is 2")
	default:
		fmt.Println("x is not 2")
	}

	day := isWeekend2(time.Now())
	fmt.Println("Is it weekend?", day)

	do2(1)
	do2("hello")
	do2(true)
	do2(float32(3.14))
}

func do(x int) {
	switch x {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	switch {
	case x == 1:
		fmt.Println("1")
	case "abc" == "foo":
	}
}

func isWeekend(x time.Time) {
	switch {
	case x.Weekday() > 0 && x.Weekday() < 6:
		fmt.Println("It's a weekday")
	default:
		fmt.Println("It's a weekend")
	}
}

func isWeekend2(x time.Time) bool {
	switch x.Weekday() {
	case time.Sunday, time.Saturday:
		return true
	default:
		return false
	}
}

func do2(x any) {

	switch x := x.(type) {
	case int:
		fmt.Println("int", x)
	case string:
		takeString(x)
		fmt.Println("string", x)
	case bool:
		fmt.Println("bool", x)
	case float32:
		fmt.Println("float32", x)
	default:
		fmt.Println("unknown type")
	}
}

func takeString(x string) {
	fmt.Println("string", x)
}
