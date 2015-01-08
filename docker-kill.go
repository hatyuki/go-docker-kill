package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	dockerapi "github.com/fsouza/go-dockerclient"
)

func main() {
	os.Exit(run())
}

func run() int {
	var signal = flag.String("signal", "KILL", "Signal to send to the container")
	flag.Parse()
	*signal = strings.ToUpper(*signal)

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "Usage: docker-kill [OPTION] CONTAINER [CONTAINER, ...]\n\n")
		fmt.Fprintf(os.Stderr, "Kill a running container using SIGKILL or a specified signal\n\n")
		flag.PrintDefaults()
		return 1
	}

	errors := 0

	for _, container_id := range flag.Args() {
		if err := KillContainer(container_id, *signal); err == nil {
			log.Printf("Send signal %s to %s\n", *signal, container_id)
		} else {
			log.Printf("[ERROR] %s", err)
			errors++
		}
	}

	return errors
}

func KillContainer(container_id, signal string) error {
	client, err := makeDockerClient()
	if err != nil {
		return err
	}

	sig, err := string2signal(signal)
	if err != nil {
		return err
	}

	options := &dockerapi.KillContainerOptions{
		ID:     container_id,
		Signal: sig,
	}

	if err := client.KillContainer(*options); err != nil {
		return err
	}

	return nil
}

func makeDockerClient() (*dockerapi.Client, error) {
	var endpoint string

	if value := os.Getenv("DOCKER_HOST"); value != "" {
		endpoint = value
	} else {
		endpoint = "unix:///var/run/docker.sock"
	}

	return dockerapi.NewClient(endpoint)
}

func string2signal(signal string) (dockerapi.Signal, error) {
	signals := map[string]dockerapi.Signal{
		"ABRT":   dockerapi.SIGABRT,
		"ALRM":   dockerapi.SIGALRM,
		"BUS":    dockerapi.SIGBUS,
		"CHLD":   dockerapi.SIGCHLD,
		"CLD":    dockerapi.SIGCLD,
		"CONT":   dockerapi.SIGCONT,
		"FPE":    dockerapi.SIGFPE,
		"HUP":    dockerapi.SIGHUP,
		"ILL":    dockerapi.SIGILL,
		"INT":    dockerapi.SIGINT,
		"IO":     dockerapi.SIGIO,
		"IOT":    dockerapi.SIGIOT,
		"KILL":   dockerapi.SIGKILL,
		"PIPE":   dockerapi.SIGPIPE,
		"POLL":   dockerapi.SIGPOLL,
		"PROF":   dockerapi.SIGPROF,
		"PWR":    dockerapi.SIGPWR,
		"QUIT":   dockerapi.SIGQUIT,
		"SEGV":   dockerapi.SIGSEGV,
		"STKFLT": dockerapi.SIGSTKFLT,
		"STOP":   dockerapi.SIGSTOP,
		"SYS":    dockerapi.SIGSYS,
		"TERM":   dockerapi.SIGTERM,
		"TRAP":   dockerapi.SIGTRAP,
		"TSTP":   dockerapi.SIGTSTP,
		"TTIN":   dockerapi.SIGTTIN,
		"TTOU":   dockerapi.SIGTTOU,
		"UNUSED": dockerapi.SIGUNUSED,
		"URG":    dockerapi.SIGURG,
		"USR1":   dockerapi.SIGUSR1,
		"USR2":   dockerapi.SIGUSR2,
		"VTALRM": dockerapi.SIGVTALRM,
		"WINCH":  dockerapi.SIGWINCH,
		"XCPU":   dockerapi.SIGXCPU,
		"XFSZ":   dockerapi.SIGXFSZ,
	}

	if sig, ok := signals[signal]; ok {
		return sig, nil
	} else {
		return sig, errors.New(fmt.Sprintf("Invalid signal: %s", signal))
	}
}
