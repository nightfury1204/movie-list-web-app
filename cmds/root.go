package cmds

import (
	"flag"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:               "movie-listing-app [command]",
		Short:             `A title based movie listing app using the OMDB API.`,
		DisableAutoGenTag: true,
	}

	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	flag.CommandLine.Parse([]string{})
	rootCmd.AddCommand(NewCmdRun())

	return rootCmd
}
