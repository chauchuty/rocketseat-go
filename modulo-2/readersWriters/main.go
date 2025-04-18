package main

import (
	"errors"
	"io"
	"os"
	"strings"
)

func main() {
	str := "hello, World!\n"
	reader := strings.NewReader(str)
	writer := MyWrite{os.Stdout}

	// buffer := make([]byte, 10000)
	buffer := make([]byte, 2)
	// n, err := reader.Read(buffer)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(buffer[:n])
	// fmt.Println(buffer[:n])
	// fmt.Println(string(buffer[:n]))
	// fmt.Println(buffer)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		_, _ = writer.Write(buffer[:n])

	}
	// io.ReadAll(reader)
	io.ReadFull(reader, nil)

}

type MyWrite struct {
	r io.Writer
}

func (mw MyWrite) Write(b []byte) (int, error) {
	for i, bb := range b {
		b[i] = bb + 10
	}
	return mw.r.Write(b)
}
