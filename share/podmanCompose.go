package share

import (
	"flag"
	"os"
)

type args struct {
	Version             bool
	In_pod              string
	Pod_args            string
	Env_file            string
	File                []string
	Profile             []string
	Project_name        string
	Podman_path         string
	Podman_args         []string
	Podman_pull_args    []string
	Podman_push_args    []string
	Podman_build_args   []string
	Podman_inspect_args []string
	Podman_run_args     []string
	Podman_start_args   []string
	Podman_stop_args    []string
	Podman_rm_args      []string
	Podman_volume_args  []string
	No_ansi             bool
	No_cleanup          bool
	Dry_run             bool
	Parallel            int
	Verbose             bool
	Command             string
	Format              string
	Short               bool
}

// TODO: not sure what kind of variables
type podmanCompose struct {
	Podman                  *podman
	podmanVersion           any
	environ                 any
	exitCode                any
	commands                map[string]func(*podmanCompose, *args)
	GlobalArgs              args
	ProjectName             string
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

var (
	PodmanCompose = &podmanCompose{}
)

func init() {
	PodmanCompose.Podman = &podman{compose: PodmanCompose, podmanPath: "podman", dryRun: false}
	PodmanCompose.defaultNet = "default"
	PodmanCompose.consoleColors = []string{
		"\x1b[1;32m",
		"\x1b[1;33m",
		"\x1b[1;34m",
		"\x1b[1;35m",
		"\x1b[1;36m",
	}
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

func (pC *podmanCompose) initGlobalParser(fs *flag.FlagSet) {

	fs.BoolVar(&pC.GlobalArgs.Version, "v", false, "show version")
	fs.BoolVar(&pC.GlobalArgs.Version, "version", false, "show version")
}

// // Constructor to initialize the PodmanCompose struct
// func (pC *PodmanCompose) Init() {

// 	pC.consoleColors =
// }

// Initialize possible args and proccess given
func (pC *podmanCompose) parseArgs() *args {
	parser := flag.NewFlagSet("parser", flag.ContinueOnError)
	pC.initGlobalParser(parser)
	return &pC.GlobalArgs
}

// func (pC *PodmanCompose) Start() {

// 	InfoLogger.Println(pC.podman.Output("", "help", ""))
// }

func (pC *podmanCompose) Run() {
	InfoLogger.Printf("selkieman-compose version: %s\n", Version)

	//get state of all arguments
	args := pC.parseArgs()
	InfoLogger.Println(*args)
	// InfoLogger.Println(reflect.TypeOf(args).Elem().Name()) //print type of struct
	podmanPath := args.Podman_path
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
			if !args.Dry_run {
				ErrorLogger.Printf("Binary %s has not been found.\n", podmanPath)
				// panic("Podman Binary has not been found")
				os.Exit(1)
			}
		}
	}
	//Initialize Podman object
	pC.Podman = &podman{compose: pC, podmanPath: podmanPath, dryRun: args.Dry_run}

	if !args.Dry_run {

		//TODO: check if should return string, python returns coroutine
		pC.podmanVersion = pC.Podman.Output([]string{"--version"}, "", []string{})
		// not found podman version
		if pC.podmanVersion == "" {
			ErrorLogger.Printf("it seems that you do not have `podman` installed")
			os.Exit(1)
		}
		InfoLogger.Printf("using podman version: %s", pC.podmanVersion)
	}
	//get function name from arguments
	cmd_name := args.Command
	//get entry (function address) from map
	cmd := pC.commands[cmd_name]
	//call function with given parameters
	cmd(pC, args)

}

func (pC *podmanCompose) parse_compose_file() {
	args := pC.GlobalArgs

	// Get the environment variable
	dirname, exists := os.Getenv("COMPOSE_PROJECT_DIR")

	if dirname != "" {

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
			if !args.Dry_run {
				ErrorLogger.Printf("Binary %s has not been found.\n", podmanPath)
				// panic("Podman Binary has not been found")
				os.Exit(1)
			}
		}
	}

}
