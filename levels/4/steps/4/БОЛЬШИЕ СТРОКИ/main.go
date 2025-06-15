package main

import (
	"strings"
)

type Reader interface {
	Read() []byte
}

type Writer interface {
	Write(p []byte) int
}

type ReaderWriter interface {
	Reader
	Writer
}

type UpperReaderWriter struct {
	UpperString string
}

func (up *UpperReaderWriter) Read() []byte {
	return []byte(up.UpperString)
}

func (up *UpperReaderWriter) Write(bytes []byte) int {
	str := string(bytes)
	up.UpperString = strings.ToUpper(str)

	return len(bytes)
}
