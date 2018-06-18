package cmd

import (
	"fmt"

	"github.com/ipochi/px-test-scripts/src"
	"github.com/spf13/cobra"
)

var number int
var wpnumber int

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: createAction,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().IntVarP(&number, "number", "n", 5, "Number of namespaces")
	createCmd.MarkFlagRequired("number")

	createCmd.PersistentFlags().IntVarP(&wpnumber, "wordpress", "w", 2, "Number of wordpress installations per namespace")
	createCmd.MarkFlagRequired("wpnumber")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createAction(cmd *cobra.Command, args []string) {
	fmt.Println("create called with number -- ", number, " and wpnumber -- ", wpnumber)
	src.Create(number, wpnumber)
}
