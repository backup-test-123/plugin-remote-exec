package main

import (
	_ "../../go/serviceplugin"
	_ "github.com/lyft/flyteplugins/go/tasks/plugins/k8s/container"
	_ "github.com/lyft/flyteplugins/go/tasks/plugins/k8s/sidecar"
	"github.com/lyft/flytepropeller/cmd/controller/cmd"
)

func main() {
	cmd.Execute()
}
