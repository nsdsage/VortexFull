package main

import "vortex/vcfg/packages/cmd"

func main() {
	cmd.Execute()
}

// TODO
// Decide how to handle default values and the scope of config files
// How much data is stored in config files (i think everything)
// Config file determines how much of the datahub to build

// type Configs struct {
// 	nifi_dir string `json:"nifi_dir"`
// }

// func (cfg *Configs) fill_defaults() {
// 	if cfg.nifi_dir == "" {
// 		cfg.nifi_dir = "../nifi-jre11/nifi-docker/dockerhub"
// 	}
// }

// func getConfigs(configfile *string) *Configs {
// 	var config Configs
// 	jsonFile, _ := os.Open(*configfile)
// 	byteValue, _ := ioutil.ReadAll(jsonFile)
// 	_ = json.Unmarshal(byteValue, &config)
// 	return &config
// }
