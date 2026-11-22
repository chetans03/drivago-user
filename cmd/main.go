package main

import (
	"fmt"

	"os"

	"github.com/chetan03/drivago-user/runner"
)

func main() {
	if err := runner.RunApplication(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
