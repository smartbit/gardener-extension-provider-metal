package main

import (
	"github.com/metal-stack/gardener-extension-provider-metal/cmd/gardener-extension-admission-metal/app"

	controllercmd "github.com/gardener/gardener/extensions/pkg/controller/cmd"
	log "github.com/gardener/gardener/pkg/logger"
	runtimelog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

func main() {
	runtimelog.SetLogger(log.ZapLogger(false))
	cmd := app.NewAdmissionCommand(signals.SetupSignalHandler())

	if err := cmd.Execute(); err != nil {
		controllercmd.LogErrAndExit(err, "error executing the main command")
	}
}
