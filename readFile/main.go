package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type ReadCloser interface {
	Read(b []byte) (n int, err error)
	Close() error
}

func main() {

	file, err := os.Open("ex.txt")
	if err != nil {
		log.Fatal("Не удалось открыть файл")
	}
	defer file.Close()

	buf := make([]byte, 1024)

	n, err := ReadAndClose(file, buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	fmt.Printf("Read %d bytes: %s\n", n, buf[:n])

}

func ReadAndClose(r ReadCloser, buf []byte) (n int, err error) {

	for len(buf) > 0 && err == nil {
		var ReadNum int

		ReadNum, err = r.Read(buf)
		n += ReadNum
		buf = buf[ReadNum:]

	}
	r.Close()
	return
}
