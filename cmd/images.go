package cmd

import (
	"fmt"

	"github.com/PiotrBzdrega/selkieman-compose/share"
	"github.com/spf13/cobra"
)

type imagesOptionsCLI struct {
	Quiet bool
}

var (
	imagesCmd = &cobra.Command{
		Use:   "images [options] [IMAGE]",
		Short: "List images used by the created containers",
		RunE:  compose_images,
	}
	imagesOptions imagesOptionsCLI
)

func init() {
	rootCmd.AddCommand(psCmd)
	listFlagSet(psCmd)
}

func listFlagSet(cmd *cobra.Command) {

	flags := cmd.Flags()

	flags.BoolVarP(&psOptions.Quiet, "quiet", "q", false, "Only display container IDs")
}

func compose_images(cmd *cobra.Command, _ []string) error {

	img_containers = [cnt for cnt in compose.containers if "image" in cnt]
	if psOptions.Quiet {
		ps_args = append(ps_args, []string{"--format", "{{.ID}}"}...)
	} else {
		ps_args = append(ps_args, []string{"--format", share.PodmanCompose.GlobalArgs.Format}...)
	}
	print(share.PodmanCompose.Podman.Output([]string{}, "ps", ps_args)) //TODO: must be run
	return nil
}
