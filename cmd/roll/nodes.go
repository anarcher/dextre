package roll

import (
	"github.com/lunarway/dextre/pkg/kubernetes"
	"github.com/lunarway/dextre/pkg/roll"
	"github.com/spf13/cobra"
)

// NewCommand sets up the move command
func nodesCommand(kubectl *kubernetes.Client, verbose *bool) *cobra.Command {
	var instanceGroupNodeLabel string
	var instanceGroupAsgLabel string
	var instanceGroup string
	var cluster string

	c := &cobra.Command{
		Use:   "nodes",
		Short: "",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			ig := roll.InstanceGroup{
				NodeLabel:             instanceGroupNodeLabel,
				AutoScalingGroupLabel: instanceGroupAsgLabel,
				Name:                  instanceGroup,
			}
			return roll.Nodes(kubectl, ig, cluster, *verbose)
		},
	}
	c.Flags().StringVar(&instanceGroupNodeLabel, "instance-group-node-label", "kops.k8s.io/instancegroup", "node label for instance group")
	c.Flags().StringVar(&instanceGroupAsgLabel, "instance-group-asg-label", "k8s.io/cluster-autoscaler/node-template/label/kops.k8s.io/instancegroup", "autoscaling label for instance group")
	c.Flags().StringVar(&instanceGroup, "instance-group", "", "kops instance group to perfrom the rolling on")
	c.MarkFlagRequired("instance-group-node-label")
	c.MarkFlagRequired("instance-group-asg-label")
	c.MarkFlagRequired("instance-group")
	c.Flags().StringVar(&label, "label", "", "label of the nodes to be rolled")
	c.Flags().StringVar(&cluster, "cluster", "", "the name of the kops cluster")
	c.MarkFlagRequired("cluster")

	return c
}
