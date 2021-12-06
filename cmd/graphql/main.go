package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// flags
	PORT string

	rootCmd = &cobra.Command{
		Use: "graphsrv",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
