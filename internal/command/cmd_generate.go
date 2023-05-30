package command

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/responserms/lls-fivem/internal/lua"
	"github.com/responserms/lls-fivem/internal/natives"
	"github.com/urfave/cli/v2"
)

func createGenerateCommand() *cli.Command {
	return &cli.Command{
		Name:      "generate",
		Usage:     "Generates FiveM native bindings for the Lua Language Server (LLS).",
		ArgsUsage: `generate [flags] <output>`,
		Action: func(ctx *cli.Context) error {
			req, err := http.NewRequestWithContext(ctx.Context, http.MethodGet, "https://runtime.fivem.net/doc/natives.json", nil)
			if err != nil {
				panic(err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				panic(err)
			}

			defer resp.Body.Close()

			bytes, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			var startAt = time.Now()

			list, err := natives.FromJSON(bytes)
			if err != nil {
				panic(err)
			}

			err = lua.Render(list, ctx.Args().First())
			if err != nil {
				return err
			}

			fmt.Fprintf(ctx.App.Writer, "Generated %d natives in %s.\n", len(list.Natives()), time.Since(startAt))

			return nil
		},
	}
}
