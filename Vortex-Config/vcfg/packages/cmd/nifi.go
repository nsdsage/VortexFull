package cmd

import (
	"vortex/vcfg/packages/vortexcfg"

	"github.com/spf13/cobra"
)

var image_loc string
var template_loc string
var template_dest string

var nifi = &cobra.Command{
	Use:   "nifi",
	Short: "Handles Nifi",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return vortexcfg.Nifi(&image_loc, &template_loc, &template_dest)
	},
}

func init() {
	nifi.Flags().StringVarP(&image_loc, "location", "l", "", "Specify location to create Nifi image")
	nifi.Flags().StringVarP(&template_loc, "template_location", "t", "", "Specify location of Nifi Templates")
	nifi.Flags().StringVarP(&template_dest, "template_destination", "d", "", "Specify destination of Nifi Templates")

	rootCmd.AddCommand(nifi)
	rootCmd.Flags()
}
