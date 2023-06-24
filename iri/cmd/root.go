package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "iri [json-file]",
		Short:   "A tool for generating IRIs",
		Long:    "A tool for generating IRIs",
		Example: "iri geonode.jsonld",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("not yet implemented")

			return nil
		},
	}

	return cmd
}
