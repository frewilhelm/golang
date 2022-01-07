package cmd

import (
	"fmt"

	"github.com/fwilhelm92/golang/goPizza/menu"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(menuCmd)
	menuCmd.Flags().StringVarP(&path, "path", "p", "menu/menu.json", "Path to the menu json-file")
}

var (
	path string

	menuCmd = &cobra.Command{
		Use:   "menu",
		Short: "Show the menu of goPizza",
		RunE: func(cmd *cobra.Command, args []string) error {
			menu, err := menu.Create(path)
			if err != nil {
				return err
			}
			fmt.Printf("Hello, my friend! This is the menu:\n%s", menu)
			return nil
		},
	}
)
