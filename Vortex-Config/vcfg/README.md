# Vortex-Config
## vcfg

`vcfg` is a cli written in go to handle the configuration and building of Vortex

## setup
**Go into to the vcfg directory**  
`cd <path/to/vcfg>`

**Make vcfg a recongnizable command (Create vcfg executable accessible by PATH)**  

Path Setup:
Confirm GOPATH has been set  
The following should return a path: `echo $GOPATH`  
If it does not, [Set GOPATH](https://go.dev/doc/gopath_code#GOPATH)

Confirm GOPATH has been added to your PATH  
`echo $PATH` should contain GOPATH or the path specified in GOPATH  
If not, add it to your path. I like the top answer's description [here](https://askubuntu.com/questions/1024732/when-i-do-an-export-path-using-terminal-which-file-does-it-save-to)

Executable
Create vcfg executable in GOPATH  
`Go install vortex/vcfg`

**Confirm everything worked**  
`vcfg -h` Should print a help menu

## Usage
Run a full configuration using a configuration file (`vcfg full -f "Config_file_path"`)

To update configuration details, see vcfg -h for commands





## Current state of `vcfg`

```
vcfg - configure vortex installations. Configure manually or use a full configuration

Usage:
  vcfg [flags]
  vcfg [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  dc          docker-compose
  help        Help about any command
  init        Full Vortex Configurations
  nifi        Handles Nifi

Flags:
  -a, --author string   Author name for copyright attribution (default "YOUR NAME")
  -h, --help            help for vcfg

```
