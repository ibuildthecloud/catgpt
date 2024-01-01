package main

import (
	"bytes"
	"io"
	"os"

	"github.com/acorn-io/cmd"
	"github.com/ibuildthecloud/catgpt/pkg/catgpt"
	"github.com/moby/term"
	"github.com/spf13/cobra"
)

type Catgpt struct {
	catgpt.Options
}

func (c *Catgpt) Run(cmd *cobra.Command, args []string) error {
	var input io.Reader = os.Stdin
	if term.IsTerminal(os.Stdin.Fd()) {
		input = bytes.NewReader(nil)
	}
	return catgpt.Run(cmd.Context(), args, input, c.Options)
}

func main() {
	cmd.Main(cmd.Command(&Catgpt{}))
}
