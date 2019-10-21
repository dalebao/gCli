package main

import "github.com/dalebao/gli/src"

func main() {
	for {
		r := src.Reader()


		src.Print(r)
	}
}
