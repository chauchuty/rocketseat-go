package main

import (
	"fmt"
	"sync"
)

func main() {
	var res int
	var j int
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < 10; i++ {
		res++
	}

	for j < 10 {
		res++
		j++
	}

	for {
		res++
		if res == 30 {
			break
		}
	}

	fmt.Println(res)

	for i, v := range arr { // _, v
		fmt.Println(i, v)
	}

	for range 10 {
		fmt.Println("oi")
	}

	for i, v := range arr {
		fmt.Println(&i, &v)
	}

	fmt.Println("=========")
	const n = 10
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
