package cmd

import (
	"vortex/vcfg/packages/vortexcfg"

	"github.com/spf13/cobra"
)

var config_file_path string

var initf = &cobra.Command{
	Use:   "init",
	Short: "Full Vortex Configurations",
	Long:  "Configures vortex using predefined configuration or by a configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		vortexcfg.Init(&config_file_path)
	},
}

func init() {
	initf.Flags().StringVarP(&config_file_path, "file", "f", "", "Specify location for Configuration File")
	rootCmd.AddCommand(initf)
}
