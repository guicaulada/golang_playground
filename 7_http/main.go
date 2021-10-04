package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// custom type to implement Writer
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Response Struct
	// Status     string
	// StatusCode int
	// Body       io.ReadClose
	//   Reader
	//     Read([]byte) (int, error)
	//   Closer
	//     Close() (error)
	// The interface io.ReadClose inherits from the Reader and Closer interfaces
	// The Reader interface is a common interface implemented by multiple types with readable data
	// fmt.Println(resp)

	// The Reader.Read() function receives a []byte as argument that is filled with data (like a Buffer)
	// data := make([]byte, 15000)
	// size, err := resp.Body.Read(data)
	// if err != nil && size == 0 {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(string(data))
	// fmt.Println(size)

	// if you use the io package you can use a Writer interface
	// this way it's not necessary to define the []byte manually
	// the Writer interface does the opposite of a Reader
	// it Writes information to a source instead of reading from one
	// the Reader reads information from somewhere to a []byte
	// the Writer writes the information from a []byte to somewhere
	// Writer { Write([]byte) (int, error) }
	//
	// The copy function receives a (Writer, Reader)
	// it pipes the information from one ot the other
	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// implement Writer interface on logWriter
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
