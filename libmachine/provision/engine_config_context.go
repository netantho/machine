package provision

import (
	"github.com/netantho/machine/libmachine/auth"
	"github.com/netantho/machine/libmachine/engine"
)

type EngineConfigContext struct {
	DockerPort       int
	AuthOptions      auth.AuthOptions
	EngineOptions    engine.EngineOptions
	DockerOptionsDir string
}
