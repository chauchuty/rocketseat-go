package main

func main() {
	x := 10
	p := &x
	*p = 20
	println(x, p, *p)
	//
	take(x)
	println(x)
	takePtr(&x)
	println(x)
	y := create()
	println(y, *y)
	//
	var z *int // derrefêrencia
	z = nil
	take2(z)
	println(z)
}

func take(x int) { // cópia
	x = 100
}

func takePtr(x *int) { // có trỏ
	*x = 100
}

func create() *int {
	x := 10
	return &x
}

func foo(x *int) {
	*x = 100
}

func take2(x *int) {
	*x = 100
}
