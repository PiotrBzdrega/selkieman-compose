package main

import (
	"encoding/json"
	"fmt"
)

func CmdCompose(pC *PodmanCompose) {
	//"show version"
	pC.commands["version"] = compose_version
	//"wait running containers to stop"
	// pC.commands["wait"]=compose_wait
}

func compose_version(pC *PodmanCompose, args *args) {

	if args.short {
		fmt.Println(version)
		return
	} else if args.format == "json" {
		res := map[string]string{
			"version": version,
		}
		fmt.Println(json.Marshal(res))
		return
	}
	fmt.Println("podman-compose version", version)

	pC.podman.Run([]string{"--version"}, "", []string{})
}
