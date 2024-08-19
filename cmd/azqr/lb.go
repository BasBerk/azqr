// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azqr

import (
	"github.com/Azure/azqr/internal/azqr"
	"github.com/Azure/azqr/internal/scanners/lb"
	"github.com/spf13/cobra"
)

func init() {
	scanCmd.AddCommand(lbCmd)
}

var lbCmd = &cobra.Command{
	Use:   "lb",
	Short: "Scan Azure Load Balancer",
	Long:  "Scan Azure Load Balancer",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		serviceScanners := []azqr.IAzureScanner{
			&lb.LoadBalancerScanner{},
		}

		scan(cmd, serviceScanners)
	},
}
