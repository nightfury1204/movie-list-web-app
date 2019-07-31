package cmds

import (
	"github.com/nightfury1204/movie-search-app/pkg/server"
	"github.com/spf13/cobra"
)

func NewCmdRun() *cobra.Command {
	cfg := server.NewConfig()
	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Launch app",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cfg.Validate(); err != nil {
				return err
			}

			if err := server.Run(cfg); err != nil {
				return err
			}
			return nil
		},
	}

	cfg.AddFlags(cmd.Flags())
	return cmd
}
