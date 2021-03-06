package kubetmuxp

import (
	"fmt"
	"os"
	"path"

	kubeconfig "github.com/karuppiah7890/kube-tmuxp-config-gen/kubeconfig"
	kubecontext "github.com/karuppiah7890/kube-tmuxp-config-gen/kubecontext"
	tmuxpconfig "github.com/karuppiah7890/kube-tmuxp-config-gen/tmuxpconfig"
	utils "github.com/karuppiah7890/kube-tmuxp-config-gen/utils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// CreateKubeTmuxpConfig creates Kubernetes Contexts and Generates Tmuxp configurations
func CreateKubeTmuxpConfig(configFile string) {
	kubeConfigsDir := path.Join(utils.GetHomeDir(), ".kube", "configs")
	tmuxpConfigDir := path.Join(utils.GetHomeDir(), ".tmuxp")

	err := os.MkdirAll(kubeConfigsDir, 0755)
	check(err)
	err = os.MkdirAll(tmuxpConfigDir, 0755)
	check(err)

	configs, err := kubeconfig.GetConfigFrom(configFile)
	check(err)

	for _, projectConfig := range configs {
		project := projectConfig.Project
		clusters := projectConfig.Clusters

		for _, cluster := range clusters {
			name := cluster.Name
			context := cluster.Context
			extraEnvs := cluster.ExtraEnvs

			fmt.Printf(">>> Running for cluster %v\n\n", name)

			kubecontext.DeleteContext(context, kubeConfigsDir)
			kubecontext.AddContext(context, cluster, project, kubeConfigsDir)
			kubecontext.RenameContext(context, cluster, project, kubeConfigsDir)
			tmuxpconfig.CreateTmuxpConfig(context, context, extraEnvs)
		}
	}
}
