package nativeengine

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	schematypes "github.com/taskcluster/go-schematypes"
	"github.com/taskcluster/taskcluster-worker/engines"
	"github.com/taskcluster/taskcluster-worker/engines/native/system"
	"github.com/taskcluster/taskcluster-worker/runtime"
)

type engineProvider struct {
	engines.EngineProviderBase
}

type engine struct {
	engines.EngineBase
	environment *runtime.Environment
	log         *logrus.Entry
	config      config
	groups      []*system.Group
}

func init() {
	engines.Register("native", engineProvider{})
}

func (engineProvider) ConfigSchema() schematypes.Schema {
	return configSchema
}

func (engineProvider) NewEngine(options engines.EngineOptions) (engines.Engine, error) {
	var c config
	if schematypes.MustMap(configSchema, options.Config, &c) != nil {
		return nil, engines.ErrContractViolation
	}

	// Load user-groups
	groups := []*system.Group{}
	for _, name := range c.Groups {
		group, err := system.FindGroup(name)
		if err != nil {
			errorMsg := fmt.Sprintf(
				"Unable to find system user-group: %s from engine config.",
				name,
			)
			options.Log.WithError(err).Error(errorMsg)
			return nil, engines.NewInternalError(errorMsg)
		}
		groups = append(groups, group)
	}

	return &engine{
		environment: options.Environment,
		log:         options.Log,
		config:      c,
		groups:      groups,
	}, nil
}

func (e *engine) PayloadSchema() schematypes.Object {
	return payloadSchema
}

func (e *engine) NewSandboxBuilder(options engines.SandboxOptions) (engines.SandboxBuilder, error) {
	var p payload
	if schematypes.MustMap(payloadSchema, options.Payload, &p) != nil {
		return nil, engines.ErrContractViolation
	}
	b := &sandboxBuilder{
		engine:  e,
		payload: p,
		context: options.TaskContext,
		env:     make(map[string]string),
		log: e.log.
			WithField("taskId", options.TaskContext.TaskID).
			WithField("runId", options.TaskContext.RunID),
	}
	return b, nil
}