package main

import (
	"fmt"
	"io"
	"os"
)

func usage(out io.Writer) {
	fmt.Fprintf(out, "Usage: %s <command> <args>\n")
	fmt.Fprintf(out, "Commands available: \n")
	fmt.Fprintf(out, "  doc\tDocumentation of nash libraries\n")
}

func main() {
	if err := do(os.Args, os.Stdout, os.Stderr); err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
}

func do(args []string, stdout, stderr io.Writer) error {
	if len(os.Args) < 2 {
		usage(os.Stderr)
		return nil
	}

	if os.Args[1] == "doc" {
		return doc(stdout, stderr, os.Args[1:])
	}

	fmt.Printf("Invalid command: %s\n", os.Args[1])
	usage(os.Stderr)
	return nil
}
