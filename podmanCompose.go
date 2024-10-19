package main

import (
	"flag"
	"os"
)

type args struct {
	version             bool
	in_pod              string
	pod_args            string
	env_file            string
	file                []string
	profile             []string
	project_name        string
	podman_path         string
	podman_args         []string
	podman_pull_args    []string
	podman_push_args    []string
	podman_build_args   []string
	podman_inspect_args []string
	podman_run_args     []string
	podman_start_args   []string
	podman_stop_args    []string
	podman_rm_args      []string
	podman_volume_args  []string
	no_ansi             bool
	no_cleanup          bool
	dry_run             bool
	parallel            int
	verbose             bool
	command             string
	format              string
	short               bool
}

// TODO: not sure what kind of variables
type PodmanCompose struct {
	podman                  *Podman
	podmanVersion           any
	environ                 any
	exitCode                any
	commands                any
	globalArgs              args
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

// Podman supported commands
func podmanCmds() []string {
	return []string{
		"pull",
		"push",
		"build",
		"inspect",
		"run",
		"start",
		"stop",
		"rm",
		"volume",
	}
}

func (pC *PodmanCompose) initGlobalParser(fs *flag.FlagSet) {

	fs.BoolVar(&pC.globalArgs.version, "v", false, "show version")
	fs.BoolVar(&pC.globalArgs.version, "version", false, "show version")
}

// // Constructor to initialize the PodmanCompose struct
// func (pC *PodmanCompose) Init() {

// 	pC.consoleColors =
// }

// Initialize possible args and proccess given
func (pC *PodmanCompose) parseArgs() *args {
	parser := flag.NewFlagSet("parser", flag.ContinueOnError)
	pC.initGlobalParser(parser)
	return &pC.globalArgs
}

// func (pC *PodmanCompose) Start() {

// 	InfoLogger.Println(pC.podman.Output("", "help", ""))
// }

func (pC *PodmanCompose) Run() {
	InfoLogger.Printf("selkieman-compose version: %s\n", version)

	args := pC.parseArgs()
	podmanPath := args.podman_path
	// podmanPath := "/usr/bin/podman" //TODO: fake it

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
			if args.dry_run == false {
				ErrorLogger.Printf("Binary %s has not been found.\n", podmanPath)
				// panic("Podman Binary has not been found")
				os.Exit(1)
			}
		}
	}
	//Initialize Podman object
	pC.podman = &Podman{compose: pC, podmanPath: podmanPath, dryRun: args.dry_run}

	if args.dry_run == false {

		// not found podman version
		if pC.podmanVersion == "" {
			ErrorLogger.Printf("it seems that you do not have `podman` installed")
			os.Exit(1)
		}
		InfoLogger.Printf("using podman version: %s", pC.podmanVersion)
	}

}
