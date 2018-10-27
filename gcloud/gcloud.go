package gcloud

import (
	"fmt"
	"os"
	"os/exec"

	kubeconfig "github.com/karuppiah7890/kube-tmuxp-config-gen/kubeconfig"
)

// IsRegional tells if the cluster is regional or not
func IsRegional(cluster kubeconfig.Cluster) bool {
	if cluster.Region != "" {
		return true
	}
	return false
}

// GetLocation gets the location of the cluster which could be a region or a zone
func GetLocation(cluster kubeconfig.Cluster) string {
	isRegional := IsRegional(cluster)
	if isRegional {
		return cluster.Region
	}
	return cluster.Zone

}

func getCommandForGetCredentials(clusterName string, project string, location string, isRegionalCluster bool, kubeConfigurationFile string) *exec.Cmd {
	if isRegionalCluster {
		command := exec.Command("gcloud", "beta", "container", "clusters", "get-credentials", clusterName, "--region", location, "--project", project)
		command.Env = append(os.Environ(), "CLOUDSDK_CONTAINER_USE_V1_API_CLIENT=false", "CLOUDSDK_CONTAINER_USE_V1_API=false", fmt.Sprintf("KUBECONFIG=%v", kubeConfigurationFile))
		return command
	}
	command := exec.Command("gcloud", "beta", "container", "clusters", "get-credentials", clusterName, "--zone", location, "--project", project)
	command.Env = append(os.Environ(), fmt.Sprintf("KUBECONFIG=%v", kubeConfigurationFile))
	return command
}

// GetCredentialsFor gets the credentials for a given cluster of a project in the given config file
func GetCredentialsFor(cluster kubeconfig.Cluster, project string, kubeConfigurationFile string) {
	isRegionalCluster := IsRegional(cluster)
	location := GetLocation(cluster)
	name := cluster.Name
	command := getCommandForGetCredentials(name, project, location, isRegionalCluster, kubeConfigurationFile)
	commandOutput, err := command.CombinedOutput()

	if err != nil {
		panic(fmt.Errorf("Error occurred while getting the cluster credentials using command : \n $ %v \n\nError: %v", command, err))
	}

	fmt.Println(string(commandOutput))
}
