package del

import (
	"time"

	"github.com/lunarway/dextre/pkg/drain"
	"github.com/lunarway/dextre/pkg/kubernetes"
	"github.com/spf13/cobra"
)

func nodeCommand(kubectl *kubernetes.Client, verbose *bool) *cobra.Command {
	var (
		nodeName        string
		gracePeriod     time.Duration
		skipValidation  bool
		nodeTermination bool
		awsRegion       string
	)

	c := &cobra.Command{
		Use:   "node",
		Short: "",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return drain.RunWithDelete(kubectl, nodeName,
				gracePeriod, skipValidation, nodeTermination, awsRegion, *verbose)
		},
	}

	c.Flags().StringVar(&nodeName, "node", "", "The node that dextre should drain in a safe manner (required)")
	c.MarkFlagRequired("node")
	c.Flags().BoolVar(&skipValidation, "skip-validation", false, "Don't ask for validations")
	c.Flags().BoolVar(&nodeTermination, "terminate-node", false, "Terminate the AWS instance in the autoscaling group")
	c.Flags().DurationVar(&gracePeriod, "grace-period", (30 * time.Second), "pod grace-period")
	c.Flags().StringVar(&awsRegion, "aws-region", "ap-northeast-2", "The region to use for node")

	return c
}
