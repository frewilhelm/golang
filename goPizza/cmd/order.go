package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(orderCmd)
}

var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Order a pizza",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := order(); err != nil {
			return fmt.Errorf("cmd: got some error: %w", err)
		}
		return nil
	},
}

func order() error {
	fmt.Println("This is the order")
	return nil
}
