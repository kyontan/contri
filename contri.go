package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type ContainerConfig struct {
	DockerHost string
	Options    string
	Image      string
	Cmd        string
}

func (cc *ContainerConfig) GetDockerArgs() []string {
	var args []string

	if cc.DockerHost != "" {
		args = append(args, "-H", cc.DockerHost)
	}

	args = append(args, "run", "--rm")

	if cc.Options != "" {
		args = append(args, cc.Options)
	}

	args = append(args, cc.Image)

	if cc.Cmd != "" {
		args = append(args, cc.Cmd)
	}

	return args
}

func (cc *ContainerConfig) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("docker", cc.GetDockerArgs()...)

	cmd.Stdout = w
	cmd.Stderr = w

	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}

func main() {

	if len(os.Args) != 5 {
		fmt.Println(os.Args[0], "host path image cmd")
		os.Exit(1)
	}

	cc := &ContainerConfig{
		DockerHost: fmt.Sprintf("tcp://%s:2375", os.Args[1]),
		Options:    "",
		Image:      os.Args[3],
		Cmd:        os.Args[4],
	}

	path := os.Args[2]
	bind := "127.0.0.1:8080"
	fmt.Println("path:", path)
	fmt.Println("bind:", bind)

	fmt.Printf("%+v\n", cc)


	http.Handle(path, cc)
	http.ListenAndServe(bind, nil)
}
