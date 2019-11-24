package serviceplugin

import (
	"context"
	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery"
	pluginsCore "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"
)

const (
	pluginID        = "example"
	exampleTaskType = "example"
)

// Sanity test that the plugin implements method of k8s.Plugin
var _ pluginsCore.Plugin = mySamplePlugin{}

type mySamplePlugin struct {
}

// Unique ID for the plugin, should be ideally the same the ID in PluginEntry
func (mySamplePlugin) GetID() string {
	return pluginID
}

// Properties desired by the plugin from the available set
func (mySamplePlugin) GetProperties() pluginsCore.PluginProperties {
	return pluginsCore.PluginProperties{}
}

// The actual method that is invoked for every single task execution. The method should be a non blocking method.
// It maybe invoked multiple times and hence all actions should be idempotent. If idempotency is not possible look at
// Transitions to get some system level guarantees
func (mySamplePlugin) Handle(ctx context.Context, tCtx pluginsCore.TaskExecutionContext) (pluginsCore.Transition, error) {
	// TODO: Implement statemachine logic
	return pluginsCore.DoTransition(pluginsCore.PhaseInfoSuccess(nil)), nil
}

// Called when the task is to be killed/aborted, because the top level entity was aborted or some other failure happened.
// Abort should always be idempotent
func (mySamplePlugin) Abort(ctx context.Context, tCtx pluginsCore.TaskExecutionContext) error {
	// TODO: Abort running task
	return nil
}

// Finalize is always called, after Handle or Abort. Finalize should be an idempotent operation
func (mySamplePlugin) Finalize(ctx context.Context, tCtx pluginsCore.TaskExecutionContext) error {
	// TODO: Cleanup reserved resources
	return nil
}

func pluginLoader(ctx context.Context, iCtx pluginsCore.SetupContext) (plugin pluginsCore.Plugin, e error) {
	iCtx.ResourceRegistrar()
	return mySamplePlugin{}, nil
}

// TODO we should register the plugin
func init() {
	pluginmachinery.PluginRegistry().RegisterCorePlugin(
		pluginsCore.PluginEntry{
			ID:                  pluginID,
			RegisteredTaskTypes: []pluginsCore.TaskType{exampleTaskType},
			IsDefault:           false,
			LoadPlugin:          pluginLoader,
		})
}
