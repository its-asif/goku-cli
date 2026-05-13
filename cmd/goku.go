/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var converterCmd = &cobra.Command{
	Use:   "converter",
	Short: "Use to convert json to yaml and viceversa",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("converter called")
	},
}

func init() {

}
