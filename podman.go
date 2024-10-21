package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"syscall"
)

func Test1() {
	fmt.Println("Testing API")
}

///////////////////
// podman and compose classes
///////////////////

type Podman struct {
	compose    *PodmanCompose
	podmanPath string
	dryRun     bool
	mutex      sync.Mutex
}

// // Constructor to initialize the Podman struct
// func (p *Podman) Init() {
// 	p.podmanPath = "podman"
// 	p.dryRun = false
// }

// execute Podman request with arguments
func (p *Podman) Output(podmanArgs []string, cmd string, cmdArgs []string) string {
	//TODO: verify if mutex works
	p.mutex.Lock()         // Acquire the mutex (using Mutex)
	defer p.mutex.Unlock() // Release it after function completes
	// args := []string{podmanArgs, cmd, cmdArgs}
	// concatenated := strings.Join(args, " ")

	//TODO: create function strToArray to merge strings
	var cmdList []string
	if podmanArgs != nil {
		cmdList = append(cmdList, podmanArgs...)
	}
	if cmd != "" {
		cmdList = append(cmdList, cmd)
	}
	if cmdArgs != nil {
		cmdList = append(cmdList, cmdArgs...)
	}
	InfoLogger.Println(p.podmanPath, cmdList[0:]) //TODO: remove square brackets
	//TODO: add logger

	command := exec.Command(p.podmanPath, cmdList...)
	dateOut, err := command.CombinedOutput() //returns stdout & stderr
	if err != nil {
		panic(string(dateOut))
		// panic(err)
	}
	return string(dateOut)
}

// replace the current Go process with Podman
func (p *Podman) Exec(podmanArgs []string, cmd string, cmdArgs []string) {

	//TODO: create function strToArray to merge strings
	var cmdList []string
	if p.podmanPath != "" {
		cmdList = append(cmdList, p.podmanPath)
	}
	if podmanArgs != nil {
		cmdList = append(cmdList, podmanArgs...)
	}
	if cmd != "" {
		cmdList = append(cmdList, cmd)
	}
	if cmdArgs != nil {
		cmdList = append(cmdList, cmdArgs...)
	}
	InfoLogger.Println(p.podmanPath, cmdList[0:]) //TODO: remove square brackets

	binary, lookErr := exec.LookPath(p.podmanPath)
	if lookErr != nil {
		panic(lookErr)
	}

	InfoLogger.Println(p.podmanPath, " path:", binary)

	err := syscall.Exec(binary, cmdList, os.Environ())

	if err != nil {
		panic(err)
	}
}

func (p *Podman) Run(podmanArgs []string, cmd string, cmdArgs []string) {
	//TODO: verify if mutex works
	p.mutex.Lock()         // Acquire the semaphore (using mutex)
	defer p.mutex.Unlock() // Release it after function completes
}
