package kubeconfig

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// Cluster represents a GKE Kubernetes Cluster
type Cluster struct {
	Name    string `json:"name"`
	Context string `json:"context"`
	Region  string `json:"region"`
	Zone    string `json:"zone"`
}

// ProjectConfig represents a GCP project config
type ProjectConfig struct {
	Project  string    `json:"project"`
	Clusters []Cluster `json:"clusters"`
}

// GetConfigFrom gets the configuration from the config file
func GetConfigFrom(configFile string) ([]ProjectConfig, error) {
	config, err := ioutil.ReadFile(configFile)

	if err != nil {
		return nil, err
	}

	var projectConfigs []ProjectConfig

	err = yaml.Unmarshal(config, &projectConfigs)

	return projectConfigs, nil
}
