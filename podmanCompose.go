package main

import (
	"os"
)

// TODO: not sure what kind of variables
type PodmanCompose struct {
	podman                  *Podman
	podmanVersion           any
	environ                 any
	exitCode                any
	commands                any
	globalArgs              any
	projectName             any
	dirname                 any
	pods                    any
	containers              []string
	vols                    any
	networks                any
	defaultNet              string
	declaredSecrets         any
	containerNamesByService any
	containerByName         any
	services                any
	allServices             any
	preferVolumeOverMount   any
	xPodman                 any
	mergedYaml              any
	yamlHash                string
	consoleColors           []string
}

// // Constructor to initialize the PodmanCompose struct
// func (pC *PodmanCompose) Init() {

// 	pC.consoleColors =
// }

func (pC *PodmanCompose) parseArgs(argv any) {
	
}

// func (pC *PodmanCompose) Start() {

// 	InfoLogger.Println(pC.podman.Output("", "help", ""))
// }

func (pC *PodmanCompose) Run() {
	InfoLogger.Printf("selkieman-compose version: %s\n", version)

	// args = self._parse_args(argv)
	// podmanPath = args.PodmanPath
	podmanPath := "/usr/bin/podman" //TODO: fake it

	if podmanPath != "podman" {

		infoFile, err := os.Stat(podmanPath)
		if err != nil {
			// If error, it could mean the file doesn't exist
			ErrorLogger.Printf("file %s doesn't exist .\n", podmanPath)
			panic(err)
		}

		InfoLogger.Println(infoFile.Mode())

		//podman is a file and is executable
		if infoFile.Mode().IsRegular() && infoFile.Mode()&0111 != 0 {
			InfoLogger.Println("podman is a file and is executable")
		} else {
			ErrorLogger.Printf("Binary %s has not been found.\n", podmanPath)
		}

	}

	pC.podman = &Podman{compose: pC, podmanPath: podmanPath, dryRun: false}

}
