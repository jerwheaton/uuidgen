package cmd

import (
	"fmt"

	"github.com/jerwheaton/uuidgen/pkg/uuid"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "uuidgen",
		Short: "Generate a UUIDv4",
		Long:  `Simple program to generate a lower case UUIDv4`,
		Run: func(cmd *cobra.Command, args []string) {
			id, err := uuid.NewV4()
			if err != nil {
				panic(err)
			}

			fmt.Println(id)
		},
	}
)

// RunCommand returns a cli command for running the application.
func RunCommand() *cobra.Command {
	return rootCmd
}
