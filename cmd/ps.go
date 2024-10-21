package cmd

import (
	"fmt"

	"github.com/PiotrBzdrega/selkieman-compose/share"
	"github.com/spf13/cobra"
)

type psOptionsCLI struct {
	Quiet bool
}

var (
	psCmd = &cobra.Command{
		Use:   "ps [options]",
		Short: "show status of containers",
		RunE:  compose_ps,
	}
	psOptions psOptionsCLI
)

func init() {
	rootCmd.AddCommand(psCmd)
	psFlagSet(psCmd)
}

func psFlagSet(cmd *cobra.Command) {

	flags := cmd.Flags()

	flags.BoolVarP(&psOptions.Quiet, "quiet", "q", false, "Only display container IDs")
}

func compose_ps(cmd *cobra.Command, _ []string) error {

	ps_args := []string{"-a", "--filter", fmt.Sprintf("label=io.podman.compose.project=%s", share.PodmanCompose.ProjectName)}
	if psOptions.Quiet {
		ps_args = append(ps_args, []string{"--format", "{{.ID}}"}...)
	} else if share.PodmanCompose.GlobalArgs.Format != "" {
		ps_args = append(ps_args, []string{"--format", share.PodmanCompose.GlobalArgs.Format}...)
	}
	print(share.PodmanCompose.Podman.Output([]string{}, "ps", ps_args)) //TODO: must be run
	return nil
}
