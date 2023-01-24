package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// var requirements bool

var rootCmd = &cobra.Command{
	Use:   "vcfg",
	Short: "vcfg - configure vortex installations",
	Long:  `vcfg - configure vortex installations. Configure manually or use a full configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	// rootCmd.Flags().BoolVarP(&requirements, "requirements", "r", false, "Lists vortex configuration requirements")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}
