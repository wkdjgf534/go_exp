/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generates UUID",
	Long: `Generates UUID v7.
For example:
securerandom uuid`,
	Run: func(cmd *cobra.Command, args []string) {
		newUUID, err := uuid.NewV7()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(newUUID.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
}
