package main

import (
	"database/sql"
	"fmt"
	"os"
)

func main() {
	fmt.Println(doDefer())
	doDefer2()
	doDefer3()
}

var arquivos = []string{"foo.txt", "bar.txt", "baz.txt"}

func doDeferFile() {
	for _, f := range arquivos {
		func() {
			// O defer é executado no final da função anônima
			file, err := os.Open(f)
			if err != nil {
				panic(err)
			}
			defer file.Close()
		}()
	}
}

func connectToDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		return db, err
	}

	// O defer é executado no final da função connectToDb
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Println("Error closing database:", err)
		}
	}()

	return db, nil
}

func doDefer() int {
	defer fmt.Print(" world")
	fmt.Print("hello")
	return 10
}

func doDefer2() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	fmt.Println(1)
}

func doDefer3() {
	x := 10
	defer func() {
		fmt.Println(x)
	}()

	defer func(y int) {
		fmt.Println(y)
	}(x)

	x = 50
	fmt.Println(x)
}

func openFile() {
	file, err := os.Open("foo.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
}
