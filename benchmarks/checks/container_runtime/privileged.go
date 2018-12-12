package containerruntime

import (
	"github.com/stackrox/rox/benchmarks/checks/utils"
	"github.com/stackrox/rox/generated/storage"
)

type privilegedBenchmark struct{}

func (c *privilegedBenchmark) Definition() utils.Definition {
	return utils.Definition{
		BenchmarkCheckDefinition: storage.BenchmarkCheckDefinition{
			Name:        "CIS Docker v1.1.0 - 5.4",
			Description: "Ensure privileged containers are not used",
		}, Dependencies: []utils.Dependency{utils.InitContainers},
	}
}

func (c *privilegedBenchmark) Run() (result storage.BenchmarkCheckResult) {
	utils.Pass(&result)
	for _, container := range utils.ContainersRunning {
		if container.HostConfig.Privileged {
			utils.Warn(&result)
			utils.AddNotef(&result, "Container '%v' (%v) is running as privileged", container.ID, container.Name)
		}
	}
	return
}

// NewPrivilegedBenchmark implements CIS-5.4
func NewPrivilegedBenchmark() utils.Check {
	return &privilegedBenchmark{}
}
