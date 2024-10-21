package share

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

func init() {
	// InfoLogger.Println("\033[95m", "Init podman.go", "\033[0m") //TODO:panic
}

///////////////////
// podman and compose classes
///////////////////

type podman struct {
	compose    *podmanCompose
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
func (p *podman) Output(podmanArgs []string, cmd string, cmdArgs []string) string {
	//TODO: verify if mutex works
	p.mutex.Lock()         // Acquire the mutex (using Mutex)
	defer p.mutex.Unlock() // Release it after function completes


	//TODO: merge LS arguments
	//TODO: create function strToArray to merge strings (not crrutial)
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
	if p == nil {
		os.Exit(0)
	}

	InfoLogger.Println(p.podmanPath, cmdList[0:]) //TODO: remove square brackets

	command := exec.Command(p.podmanPath, cmdList...)
	dateOut, err := command.CombinedOutput() //returns stdout & stderr
	if err != nil {
		panic(string(dateOut))
		// panic(err)
	}
	return string(dateOut)
}

// replace the current Go process with Podman
func (p *podman) Exec(podmanArgs []string, cmd string, cmdArgs []string) {

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

func (p *podman) Run(podmanArgs []string, cmd string, cmdArgs []string) {
	//TODO: verify if mutex works
	p.mutex.Lock()         // Acquire the semaphore (using mutex)
	defer p.mutex.Unlock() // Release it after function completes
}
