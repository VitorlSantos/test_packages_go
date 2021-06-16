package main

import (
	"github.com/VitorlSantos/test/src/samplepkg"
	"github.com/VitorlSantos/test/src/samplepkg/subpkg"
)

func main() {
	sample := samplepkg.New("Testing Sample Package")
	sample.Print()
	sub := subpkg.New("Testing sub package")
	sub.Print()
}
