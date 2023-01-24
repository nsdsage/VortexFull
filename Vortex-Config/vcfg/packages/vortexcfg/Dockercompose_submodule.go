package vortexcfg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sync"
)

func Dockercompose_submodule(dockercomposearg *string) {
	//vcfg dc -c command: e.g. up/start/stop/down other
	var wg sync.WaitGroup

	if *dockercomposearg != "" {
		//-c

		var files *[]string = find("../../")

		var command string = "docker-compose"

		for _, f := range *files {
			var args []string

			dir, _ := filepath.Split(f)
			args = append(args, "-f")
			dc := filepath.Join(dir, "docker-compose.yml")
			args = append(args, dc)
			args = append(args, "--env-file")
			env := filepath.Join(dir, ".env")
			args = append(args, env)
			args = append(args, *dockercomposearg)

			cmd := exec.Command(command, args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			wg.Add(1)
			go runCmd(cmd)
			defer wg.Done()
		}
		wg.Wait()
	}
}

func runCmd(cmd *exec.Cmd) {

	fmt.Println("Run")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}

func find(dir string) *[]string {
	// "//?(?:[^//]+//?)*/docker-compose[/.]yml"
	reg, _ := regexp.Compile("docker-compose[/.]yml")

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var ret []string
	for _, file := range files {
		if file.IsDir() {
			var fp1 string = filepath.Join(dir, file.Name())
			files, err := ioutil.ReadDir(fp1)
			if err != nil {
				log.Fatal(err)
			}
			for _, file := range files {
				if reg.MatchString(file.Name()) {
					ret = append(ret, filepath.Join(fp1, file.Name()))
				}
			}

		}
	}

	return &ret
}
