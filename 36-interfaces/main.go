package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Fprintln(os.Stdout, "Hello World")

	fw := NewFileWriter("data.txt")

	fmt.Fprintln(fw, "Hello Shell Folks")
	// io.Writer
	// io.Writer Write(p []byte) (n int, err error)

	// to create errors particularly

	var fw1 *FileWriter

	// fw2 := NewFileWriter("")

	_, err := fmt.Fprintln(fw1, "Hello Shell")
	if err != nil {
		switch o := err.(type) {
		case *FileError:
			{
				//var fileError *FileError = o
				println("Asserted to FileError poiner")
				fmt.Println(o.Code, o.Msg)

			}
		default:
			println(err.Error())
		}

	}

	// str := "hello world"
	// bytes := []byte(str)
	// fmt.Println(bytes)
	// fmt.Println(&bytes[0])
	// fmt.Println(&str)

	// runes := []rune(str)
	// fmt.Println(string(runes))
	// if it is an interface , you can type assert it
}

type FileWriter struct {
	FileName string
}

type FileError struct {
	Code string
	Msg  string
}

func NewFileError(code, msg string) *FileError {
	return &FileError{Code: code, Msg: msg}
}

func (fe *FileError) Error() string {
	return fmt.Sprintf("Code:%v Msg:%s", fe.Code, fe.Msg)
}

func NewFileWriter(fileName string) *FileWriter { // Kind of a constructor
	return &FileWriter{fileName}
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	if fw == nil {
		return 0, NewFileError("101", "nil FileWriter")
	}
	if fw.FileName == "" {
		return 0, NewFileError("101", "no file name in FileWriter")
	}

	file, err := os.OpenFile(fw.FileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return 0, NewFileError("102", err.Error())
	}
	defer file.Close()

	n, err = file.Write(p)

	if err != nil {
		return 0, NewFileError("103", err.Error())
	}
	return n, nil

}
