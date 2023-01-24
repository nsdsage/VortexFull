package cmd

import (
	"vortex/vcfg/packages/vortexcfg"

	"github.com/spf13/cobra"
)

var dockercomposecommand string

var dockercompose = &cobra.Command{
	Use:   "dc",
	Short: "docker-compose",
	Long:  "Runs docker-compose for each submodule of the Vortex Data Pipeline",
	Run: func(cmd *cobra.Command, args []string) {
		vortexcfg.Dockercompose_submodule(&dockercomposecommand)
	},
}

func init() {
	dockercompose.Flags().StringVarP(&dockercomposecommand, "command", "c", "", "docker-compose [command]")
	rootCmd.AddCommand(dockercompose)
}
