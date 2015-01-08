package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hatyuki/go-docker-kill"
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
		if err := dockerkill.KillContainer(container_id, *signal); err == nil {
			log.Printf("Send signal %s to %s\n", *signal, container_id)
		} else {
			log.Printf("[ERROR] %s", err)
			errors++
		}
	}

	return errors
}
