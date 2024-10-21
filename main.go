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

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	version       = "0.1.0"
)

func init() {
	//TODO: activate when release
	// file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	//     log.Fatal(err)
	// }

	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	script, err := filepath.Abs(os.Args[0])
	if err != nil {
		ErrorLogger.Println(err)
		return
	}

	InfoLogger.Println("Script path:", script)

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

	InfoLogger.Println("cmd arguments ", os.Args)

	podmanCompose := PodmanCompose{defaultNet: "default", yamlHash: "",
		consoleColors: []string{
			"\x1b[1;32m",
			"\x1b[1;33m",
			"\x1b[1;34m",
			"\x1b[1;35m",
			"\x1b[1;36m",
		}}

	podmanCompose.Run()
	// podmanCompose.Start()

}
