package config

import (
	"fmt"
	"reflect"
)

// Config specifies configurable options through env vars
// nolint:lll
type Config struct {
	Port     string `env:"SERVER_PORT" envDocs:"The port which the service will listen to" envDefault:"8080"`
	BasePath string `env:"SERVER_BASE_PATH" envDocs:"The base path of this service" envDefault:"/api"`
}

// HelpDocs returns documentation of Config based on field tags
func (envVar Config) HelpDocs() []string {
	reflectEnvVar := reflect.TypeOf(envVar)
	doc := make([]string, 1+reflectEnvVar.NumField())
	doc[0] = "Environment variables config:"
	for i := 0; i < reflectEnvVar.NumField(); i++ {
		field := reflectEnvVar.Field(i)
		envName := field.Tag.Get("env")
		envDefault := field.Tag.Get("envDefault")
		envDocs := field.Tag.Get("envDocs")
		doc[i+1] = fmt.Sprintf("  %v\t %v (default: %v)", envName, envDocs, envDefault)
	}
	return doc
}
