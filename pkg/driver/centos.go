package driver

import (
	"github.com/Azure/azure-docker-extension/pkg/executil"
)

// CentOSDriver is for CentOS-based distros.
type CentOSDriver struct {
	systemdBaseDriver
	systemdUnitOverwriteDriver
}

func (c CentOSDriver) InstallDocker() error {
	return executil.ExecPipe("/bin/sh", "-c", "curl -sSL https://get.docker.com/ | sh")
}

func (c CentOSDriver) UninstallDocker() error {
	return executil.ExecPipe("yum", "-y", "-q", "remove", "docker-engine.x86_64")
}

func (c CentOSDriver) DockerComposeDir() string { return "/usr/local/bin" }
