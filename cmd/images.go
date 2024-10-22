package cmd

import (
	"strings"

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
	rootCmd.AddCommand(imagesCmd)
	imagesFlagSet(imagesCmd)
}

func imagesFlagSet(cmd *cobra.Command) {

	flags := cmd.Flags()
	flags.BoolVarP(&imagesOptions.Quiet, "quiet", "q", false, "Only display container IDs")
}

func compose_images(cmd *cobra.Command, _ []string) error {

	share.PodmanCompose.Parse_compose_file()

	// Create a slice to store the filtered results of containers
	var imgContainers []map[string]interface{}
	var imgNames []string

	// Loop through the containers and store only those with "image" key
	for _, cnt := range share.PodmanCompose.Containers {
		if _, exists := cnt["image"]; exists {
			//key exists -> store
			imgContainers = append(imgContainers, cnt)
		}
	}

	data := []string{}
	if imagesOptions.Quiet {
		// Loop through the containers and filter based on the condition
		for _, cnt := range imgContainers {
			if strings.Contains(cnt, "image") {
				imgContainers = append(imgContainers, cnt)
			}
		}

		ps_args = append(ps_args, []string{"--format", "{{.ID}}"}...)
	} else {
		ps_args = append(ps_args, []string{"--format", share.PodmanCompose.GlobalArgs.Format}...)
	}

	print(share.PodmanCompose.Podman.Output([]string{}, "images", ps_args))
	return nil
}
