package utils

import (
	"github.com/openshift/installer/pkg/types"
)

func IsSingleNode(installConfig *types.InstallConfig) bool {
	if len(installConfig.Compute) != 1 {
		return false
	}
	computeReplicaCount := *installConfig.Compute[0].Replicas
	controlPlaneReplicaCount := *installConfig.ControlPlane.Replicas
	return controlPlaneReplicaCount == 1 && computeReplicaCount == 0
}
