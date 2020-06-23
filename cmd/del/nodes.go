package del

import (
	"github.com/lunarway/dextre/pkg/kubernetes"
	"github.com/lunarway/dextre/pkg/roll"
	"github.com/spf13/cobra"
)

func nodesCommand(kubectl *kubernetes.Client, verbose *bool) *cobra.Command {

	var awsRegion string
	var skipDrain bool
	var dryRun bool

	c := &cobra.Command{
		Use:   "nodes",
		Short: "",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return roll.DeleteCordonNodes(kubectl, awsRegion, skipDrain, dryRun, *verbose)
		},
	}

	c.Flags().StringVar(&awsRegion, "aws-region", "ap-northeast-2", "the region to use")
	c.Flags().BoolVar(&skipDrain, "skip-drain", false, "skip drain")
	c.Flags().BoolVar(&dryRun, "dry-run", false, "dry run")

	return c
}
