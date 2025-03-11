package main

import (
	"fmt"
	"os"

	"github.com/PrefectHQ/mex-admission-webhook-certificator/cmd"
)

func main() {
	if err := cmd.Execute(os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "\n%v\n", err)
		os.Exit(1)
	}
}
