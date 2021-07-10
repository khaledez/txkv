package main

import "io"

func Cli(input io.Reader, output io.Writer) error {
	output.Write([]byte("Hello world"))
	return nil
}
