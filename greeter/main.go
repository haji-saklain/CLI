package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "greeter",
		Short: "abcd",
		Long:  "ABCD",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Greeter")
		},
	}

	greetCmd := &cobra.Command{
		Use:   "greet",
		Short: "Greet the user",
		Long:  "Greet the user with a message",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %s\n", args[0])
		},
	}

	rootCmd.AddCommand(greetCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
