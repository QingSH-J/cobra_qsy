/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var isloud bool

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet [name]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := "卿诗雨"
		if len(args) > 0 {
			name = args[0]
		}
		message := fmt.Sprintf("Hello, %s!", name)
		if isloud {
			message = strings.ToUpper(message) + "!!!"
		}
		fmt.Println(message)
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// greetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// greetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	greetCmd.Flags().BoolVarP(&isloud, "loud", "l", false, "Make the greeting loud")
}
