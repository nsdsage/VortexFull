package vortexcfg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Nifi(image_loc *string, template_loc *string, template_dest *string) error {
	//vcfg nifi

	if *template_loc != "" && *template_dest != "" {
		//-t, -d

		fmt.Println("Importing Nifi Templates")
		var command string = "cp -r " + filepath.Dir(*template_loc) + "/* " + filepath.Dir(*template_dest)
		cmd := exec.Command("/bin/sh", "-c", command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Printf("error retrieving templates from %s\n", *template_loc)
			return err
		}
		fmt.Println("Done.")
	}

	if *image_loc != "" {
		//-l

		fmt.Println("Creating Nifi Image")
		cmd := exec.Command("docker", "build", "-t", "apache/nifi:jre-11", *image_loc)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Printf("\ncould not create nifi image in %s\n", *image_loc)
			return err
		}
		fmt.Println("Done.")
	}

	return nil
}
