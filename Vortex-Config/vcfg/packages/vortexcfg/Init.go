package vortexcfg

import (
	"log"

	"github.com/spf13/viper"
)

func Init(config_file_path *string) {
	//vcfg init

	if *config_file_path != "" {
		//-f

		fill_defaults()
		read_cfg(config_file_path)
		image_loc, template_loc, template_dest := viper.GetString("nifi.image_location"), viper.GetString("nifi.template_location"), viper.GetString("nifi.template_destination")
		if image_loc != "" || template_loc != "" || template_dest != "" {
			Nifi(&image_loc, &template_loc, &template_dest)
		}
	}
}

func fill_defaults() {
	viper.SetDefault("nifi.image_location", "")
	viper.SetDefault("nifi.template_location", "")
	viper.SetDefault("nifi.template_destination", "")
}

func read_cfg(config_file_path *string) {
	viper.SetConfigType("json")
	viper.SetConfigFile(*config_file_path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
