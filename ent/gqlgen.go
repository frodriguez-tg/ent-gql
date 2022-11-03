//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	p := modelgen.Plugin{
		MutateHook: mutateHook,
	}

	if err := api.Generate(cfg, api.ReplacePlugin(&p)); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}

}

// Defining mutation function
func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {
			i := strings.LastIndex(field.Tag, `"`)
			nt := field.Tag[:i] + `,omitempty"`
			field.Tag = nt

			if strings.EqualFold(field.Description, "warning if empty") {
				field.Tag += ` empty:"warning"`
			}
		}
	}

	return b
}
