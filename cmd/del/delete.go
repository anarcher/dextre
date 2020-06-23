package del

import (
	"github.com/lunarway/dextre/pkg/kubernetes"
	"github.com/spf13/cobra"
)

func NewCommand(kubectl *kubernetes.Client, verbose *bool) *cobra.Command {
	c := &cobra.Command{
		Use:   "delete",
		Short: "",
		Long:  "",
	}
	c.AddCommand(
		nodesCommand(kubectl, verbose),
	)

	return c
}
