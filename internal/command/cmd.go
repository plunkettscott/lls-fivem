package command

import (
	"context"
	"os"

	"github.com/urfave/cli/v2"
)

func App() *cli.App {
	return &cli.App{
		Name:        "lls-fivem",
		Description: "A tool for generating FiveM native bindings for the Lua Language Server (LLS).",
		Commands: []*cli.Command{
			createGenerateCommand(),
		},
	}
}

func Run(ctx context.Context) error {
	cmd := App()
	return cmd.RunContext(ctx, os.Args)
}
