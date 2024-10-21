package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	// "github.com/spf13/cobra"
	"github.com/PiotrBzdrega/selkieman-compose/cmd"
	"github.com/PiotrBzdrega/selkieman-compose/share"
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
		share.ErrorLogger.Println(err)
		return
	}

	share.InfoLogger.Println("Script path:", script)

	maps := []string{"apple", "raspberry", "greiphfruit"}
	do := make(chan bool, 1)
	go do_json(do, maps)

	<-do

	data := `{"name": "Rob", "age": 18}`

	var obj any
	err = json.Unmarshal([]byte(data), &obj)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// obj now contains the parsed JSON data, but its type is "any"
	// We can use a type assertion to access its properties:
	m := obj.(map[string]any)
	fmt.Println("Name:", m["name"])
	fmt.Println("Age:", m["age"])

	// podman := Podman{}
	// podman.Init()
	// InfoLogger.Println("Podman Initialized:", podman.compose, podman.podmanPath)
	// InfoLogger.Println(podman.Output("", "help", ""))
	// podman.Exec("", "--version", "")

	share.InfoLogger.Println("cmd arguments ", os.Args)

	cmd.Execute()
	os.Exit(0)
	// test_command := &cobra.Command{}
	// share.PodmanCompose.Run()
	// podmanCompose.Start()

}
