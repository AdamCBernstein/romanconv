package main

import (
	"fmt"
	"github.com/AdamCBernstein/romanconv"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: roman 1-3999\n")
		os.Exit(1)
	}

	decimal, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	romanStr, err := romanconv.IntToRoman(decimal)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d in Roman numerals is %q\n", decimal, romanStr)
}
