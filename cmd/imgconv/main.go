package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tenntenn/imgconv"
)

var (
	flagTo      = imgconv.PNG
	flagFrom    = imgconv.TIFF
	flagVersion bool
)

func init() {
	flag.Var(&flagTo, "to", "after format")
	flag.Var(&flagFrom, "from", "before format")
	flag.BoolVar(&flagVersion, "v", false, "print version")
}

func main() {
	flag.Parse()

	if flagVersion {
		fmt.Println("v0.2.0")
		os.Exit(0)
	}

	if err := imgconv.ConvertAll(os.Args[1], flagFrom, flagTo); err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
		os.Exit(1)
	}
}
