// Package lua implements a Renderer for the Lua language server.
package lua

import (
	"embed"
	"os"
	"path"
	"text/template"

	"github.com/responserms/lls-fivem/internal/natives"
)

//go:embed templates
var templateFS embed.FS

type Graph struct {
	List natives.List
}

func Render(l natives.List, outDir string) error {
	template, err := template.ParseFS(templateFS, "templates/natives.lua.tmpl")
	if err != nil {
		return err
	}

	wr, err := os.Create(path.Join(outDir, "natives.lua"))
	if err != nil {
		return err
	}

	defer wr.Close()

	if err := template.Execute(wr, Graph{List: l}); err != nil {
		return err
	}

	return nil
}
