package cq

import (
	"fmt"
	"os"
)

func cq(question string) {
	fmt.Fprintln(os.Stdout, "Q:", question, "[y/n]")

	var stdin string
	fmt.Scan(&stdin)
	if stdin != "y" {
		os.Exit(0)
	}
}
