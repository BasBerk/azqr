// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azqr

import (
	"github.com/Azure/azqr/internal/azqr"
	"github.com/Azure/azqr/internal/scanners/iot"
	"github.com/spf13/cobra"
)

func init() {
	scanCmd.AddCommand(iotCmd)
}

var iotCmd = &cobra.Command{
	Use:   "iot",
	Short: "Scan Azure IoT Hub",
	Long:  "Scan Azure IoT Hub",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		serviceScanners := []azqr.IAzureScanner{
			&iot.IoTHubScanner{},
		}

		scan(cmd, serviceScanners)
	},
}
