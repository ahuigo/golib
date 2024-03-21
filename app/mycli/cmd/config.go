/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("config with args:%v", args)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// configCmd.PersistentFlags().String("user", "u", "username")
	configCmd.Flags().BoolP("user", "u", false, "username")
}
