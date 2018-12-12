package kubelet

import (
	"github.com/stackrox/rox/benchmarks/checks"
	"github.com/stackrox/rox/benchmarks/checks/utils"
	"github.com/stackrox/rox/generated/storage"
)

type caFilePermissions struct{}

func (c *caFilePermissions) Definition() utils.Definition {
	return utils.Definition{
		BenchmarkCheckDefinition: storage.BenchmarkCheckDefinition{
			Name:        "CIS Kubernetes v1.2.0 - 2.2.7",
			Description: "Ensure that the certificate authorities file permissions are set to 644 or more restrictive",
		}, Dependencies: []utils.Dependency{utils.InitKubeletConfig},
	}
}

func (c *caFilePermissions) Run() (result storage.BenchmarkCheckResult) {
	utils.Pass(&result)
	params, ok := utils.KubeletConfig.Get("client-ca-file")
	if !ok {
		utils.Warn(&result)
		utils.AddNotes(&result, "Cannot check kubelet CA file permissions because kubelet command line does not define 'client-ca-file' parameter")
		return
	}

	result = utils.NewPermissionsCheck("", "", params.String(), 0644, true).Run()
	return
}

// NewCAFilePermissions implements CIS Kubernetes v1.2.0 2.2.7
func NewCAFilePermissions() utils.Check {
	return &caFilePermissions{}
}

func init() {
	checks.AddToRegistry(NewCAFilePermissions())
}
