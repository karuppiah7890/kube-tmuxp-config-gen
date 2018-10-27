package tmuxpconfig

import (
	"fmt"
	"os"
	"path"
	"text/template"

	templates "github.com/karuppiah7890/kube-tmuxp-config-gen/templates"
	utils "github.com/karuppiah7890/kube-tmuxp-config-gen/utils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// CreateTmuxpConfig creates tmuxp config
func CreateTmuxpConfig(kubeContext string, tmuxSessionName string, extraEnvs map[string]string) {
	fmt.Println("Creating tmuxp configuration")
	templateByteData, err := templates.Asset("templates/tmuxp-config.yaml")
	check(err)
	tmuxpconfigTemplate, err := template.New("tmuxpconfig").Parse(string(templateByteData))
	check(err)

	templateValues := make(map[string]interface{})
	kubeConfigFile := path.Join(utils.GetHomeDir(), ".kube", "configs", kubeContext)
	templateValues["KubeConfig"] = kubeConfigFile
	templateValues["SessionName"] = tmuxSessionName
	templateValues["ExtraEnvs"] = extraEnvs

	tmuxpConfigFilePath := path.Join(utils.GetHomeDir(), ".tmuxp", fmt.Sprintf("%v.yaml", kubeContext))
	tmuxpConfigFile, err := os.Create(tmuxpConfigFilePath)
	check(err)
	defer tmuxpConfigFile.Close()
	err = tmuxpconfigTemplate.ExecuteTemplate(tmuxpConfigFile, "tmuxpconfig", templateValues)
	check(err)
	fmt.Println("Created tmuxp configuration")
}
