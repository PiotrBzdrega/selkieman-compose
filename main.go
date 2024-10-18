package main

import (
	"fmt"
	"os"
	"path/filepath"

	"log"

	"encoding/json"
)

///////////////////
// podman and compose classes
///////////////////

// type Podman struct{
// compose
// podman_path
// dry_run
// semaphore

// }

func do_json(done chan bool, strings []string) {
	//json
	done <- true
	str, _ := json.Marshal(strings)
	fmt.Println(string(str))
	// time.Sleep(5 * time.Second)
	// done <- true
}

func main() {
	script, err := filepath.Abs(os.Args[0])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Script path:", script)

	//log
	log.Println("standard logger1")

	maps := []string{"apple", "raspberry", "greiphfruit"}
	do := make(chan bool, 1)
	go do_json(do, maps)

	<-do
	//log
	log.Println("standard logger2")

	Test1()

	podman := Podman{}
	podman.Init("selkieman-compose", "podman", false)
	fmt.Println("Podman Initialized:", podman.compose, podman.podmanPath)
	fmt.Println(podman.Output("", "help", ""))
	podman.Exec("", "--version", "")
	log.Println("standard logger3")

}
