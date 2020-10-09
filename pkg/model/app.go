package model

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	// HTTPProtocol indicates HTTP Dapr protocol
	HTTPProtocol = "HTTP"
	// GRPCProtocol indicates gRPC Dapr protocol
	GRPCProtocol = "gRPC"
)

// App represents app state
type App struct {
	Meta       Meta         `yaml:"Meta"`
	PubSubs    []*PubSub    `yaml:"PubSubs"`
	Bindings   []*Component `yaml:"Bindings"`
	Services   []*Service   `yaml:"Services"`
	Components []*Component `yaml:"Components"`
}

// Marshal serializes App to YAML
func (a *App) Marshal() ([]byte, error) {
	return yaml.Marshal(a)
}

// String serializes App to YAML string
func (a *App) String() string {
	b, err := a.Marshal()
	if err != nil {
		return "Unable to parse application"
	}
	return string(b)
}

// Unmarshal deserializes bytes into an App
func Unmarshal(in []byte) (*App, error) {
	var a App
	err := yaml.Unmarshal([]byte(in), &a)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling data")
	}
	return &a, nil
}
