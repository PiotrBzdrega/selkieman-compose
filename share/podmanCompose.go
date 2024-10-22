package share

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
)

type args struct {
	Version             bool
	In_pod              string
	Pod_args            string
	Env_file            string
	File                string
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
	Containers              []map[string]interface{}
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

func (pC *podmanCompose) Parse_compose_file() {
	args := pC.GlobalArgs

	// Get the environment variable
	dirname := os.Getenv("COMPOSE_PROJECT_DIR")

	// Directory
	if dirname != "" {

		infoFile, err := os.Stat(dirname)
		if err != nil {
			// If error, it could mean the file doesn't exist
			ErrorLogger.Printf("Directory %s doesn't exist .\n", dirname)
			panic(err)
		} else {
			if infoFile.Mode().IsDir() {
				//Valid directory -> change
				os.Chdir(dirname)
			}
		}
	}
	// Get the environment variable for path separator
	pathsep := os.Getenv("COMPOSE_PATH_SEPARATOR")
	if pathsep == "" {
		pathsep = string(os.PathSeparator)
	}

	if args.File == "" {
		default_str := os.Getenv("COMPOSE_FILE")
		default_ls := []string{}
		if default_str != "" {
			default_ls = strings.Split(default_str, string(os.PathListSeparator))
		} else {
			default_ls = []string{
				"compose.yaml",
				"compose.yml",
				"compose.override.yaml",
				"compose.override.yml",
				"podman-compose.yaml",
				"podman-compose.yml",
				"docker-compose.yml",
				"docker-compose.yaml",
				"docker-compose.override.yml",
				"docker-compose.override.yaml",
				"container-compose.yml",
				"container-compose.yaml",
				"container-compose.override.yml",
				"container-compose.override.yaml",
			}
		}
		for _, path := range default_ls {
			if _, err := os.Stat(path); err == nil {
				//Found compose file do net check further
				args.File = path
				break
			}
		}
	}
	file := args.File

	if file == "" {
		ErrorLogger.Println("no compose.yaml, docker-compose.yml or container-compose.yml file found, ")
		ErrorLogger.Println("pass files with -f")
		os.Exit(1)
	}

	//TODO:missing stage, do not understand why check many files
	// 	ex = map(lambda x: x == '-' or os.path.exists(x), files)
	// 	missing = [fn0 for ex0, fn0 in zip(ex, files) if not ex0]
	// 	if missing:
	// 		log.fatal("missing files: %s", missing)
	// 		sys.exit(1)

	// relative_files = files //TODO:missing stage, do not understand why check many files
	filename := file
	project_name := args.Project_name

	// Get the absolute path of the directory containing the file
	dirname, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		ErrorLogger.Println("Cannot get absolute path to file")
		panic(err)
	}
	// Get the base name of the directory
	dirBasename := filepath.Base(dirname)
	os.Chdir(dirname)

	if args.Env_file != "" {
		// Load .env from the Compose file's directory to preserve
		// behavior prior to 1.1.0 and to match with Docker Compose (v2).
		if ".env" == args.Env_file {
			project_dotenv_file = os.path.realpath(os.path.join(dirname, ".env"))

			if os.path.exists(project_dotenv_file) {
				dotenv_dict.update(dotenv_to_dict(project_dotenv_file))
			}

		}

		dotenv_path = os.path.realpath(args.env_file)
		dotenv_dict.update(dotenv_to_dict(dotenv_path))

	}

}
