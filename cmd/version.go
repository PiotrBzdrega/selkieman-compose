package cmd

import (
	"github.com/PiotrBzdrega/selkieman-compose/share"
	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "show version",
		RunE:  compose_version,
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func compose_version(cmd *cobra.Command, arg []string) error {

	// if args.short {
	// 	fmt.Println(share.Version)
	// 	return
	// } else if args.format == "json" {
	// 	res := map[string]string{
	// 		"version": share.Version,
	// 	}
	// 	fmt.Println(json.Marshal(res))
	// 	return
	// }
	// fmt.Println("podman-compose version", share.Version)
	print(share.PodmanCompose.Podman.Output([]string{"--version"}, "", []string{}))
	return nil
}
