package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func usage(cmd *cobra.Command, xit int, msgs ...string) {
	var w io.Writer = os.Stdout
	if xit > 0 {
		w = os.Stderr
	}
	for _, msg := range msgs {
		fmt.Fprintln(w, msg)
	}
	fmt.Fprintln(w, cmd.Usage())
	os.Exit(xit)
}
