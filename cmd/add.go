package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your list",
	Long:  `Adds a task to your list`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add called")
	},
}
