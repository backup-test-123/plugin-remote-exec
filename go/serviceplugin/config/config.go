package config

import pluginsConfig "github.com/lyft/flyteplugins/go/tasks/config"

//go:generate pflags Config --default-var=defaultConfig

// Per Plugin optional configuration
type Config struct {
	// TODO Add any configuration parameters here. For example, if region should be preset? or if there are some common parameters that should be passed
	// for every execution. If this is not the case, just delete the config folder
	MyConfigParam string `json:"myCfgParam" pflag:",Some configuration that is used by the backend plugin "`
}

var (
	defaultConfig = &Config{MyConfigParam: "default-value"}

	// Default values can be provided directly when registering the section
	section = pluginsConfig.MustRegisterSubSection("servicePlugin", defaultConfig)
)

func Get() *Config {
	return section.GetConfig().(*Config)
}
