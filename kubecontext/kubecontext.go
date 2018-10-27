package kubecontext

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	gcloud "github.com/karuppiah7890/kube-tmuxp-config-gen/gcloud"
	kubeconfig "github.com/karuppiah7890/kube-tmuxp-config-gen/kubeconfig"
)

// DeleteContext deletes a Kubernetes Context
func DeleteContext(context string, kubeConfigsDir string) {
	kubeConfigurationFile := path.Join(kubeConfigsDir, context)

	err := os.RemoveAll(kubeConfigurationFile)

	if err != nil {
		panic(fmt.Errorf("Error occurred while removing the file (context) %v . Error: %v", kubeConfigurationFile, err))
	}
}

// AddContext adds a Kubernetes Context
func AddContext(context string, cluster kubeconfig.Cluster, project string, kubeConfigsDir string) {
	kubeConfigurationFile := path.Join(kubeConfigsDir, context)

	gcloud.GetCredentialsFor(cluster, project, kubeConfigurationFile)
}

// RenameContext renames a Kubernetes Context
func RenameContext(context string, cluster kubeconfig.Cluster, project string, kubeConfigsDir string) {
	kubeConfigurationFile := path.Join(kubeConfigsDir, context)
	clusterName := cluster.Name
	location := gcloud.GetLocation(cluster)

	command := exec.Command("kubectl", "config", "rename-context", fmt.Sprintf("gke_%v_%v_%v", project, location, clusterName), context)
	command.Env = []string{fmt.Sprintf("KUBECONFIG=%v", kubeConfigurationFile)}

	commandOutput, err := command.Output()

	if err != nil {
		panic(fmt.Errorf("Error occurred while renaming the kubernetes context using the command : \n $ %v \n\nError: %v", command, err))
	}

	fmt.Println(string(commandOutput))
}
