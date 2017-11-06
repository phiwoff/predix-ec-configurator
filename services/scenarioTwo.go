package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/indaco/predix-ec-configurator/helpers"
)

// ScenarioTwoService ...
type ScenarioTwoService struct {
	appConfig  *helpers.AppSettings
	userConfig *helpers.UserConfig
}

// NewScenarioTwoService ...
func NewScenarioTwoService(appConfig *helpers.AppSettings, userConfig *helpers.UserConfig) (*ScenarioTwoService, error) {
	return &ScenarioTwoService{
		appConfig:  appConfig,
		userConfig: userConfig,
	}, nil
}

// ScenarioTwoConfigurationForm ...
type ScenarioTwoConfigurationForm struct {
	TypeOfScenario string
	// UAA
	UAAServiceURI   string
	UAAClient       string
	UAAClientSecret string
	// EC
	ECZoneID      string
	ECServiceURI  string
	ECAdminToken  string
	ECIDS         []string
	ECGatewayName string
	ECServerName  string
	// Predix Resources
	ResourceHost string
	ResourcePort string
	//
	OnPremiseOS     string
	LocalPort       string
	IsDebug         string
	HealthcheckPort string
	// Networking
	Proxy string
}

// GenerateAndSaveAllScenarioTwo ...
func (sc *ScenarioTwoService) GenerateAndSaveAllScenarioTwo(outputMarkdownFile string, configurationForm ScenarioTwoConfigurationForm) error {
	gatewayContent := generateAndSaveGatewayScriptScenarioTwo(sc.appConfig, configurationForm)
	gatewayManifestContent := generateAndSaveGatewayManifestScenarioTwo(sc.appConfig, configurationForm)
	serverContent := generateAndSaveServerScriptScenarioTwo(sc.appConfig, sc.userConfig, configurationForm)
	serverManifestContent := generateAndSaveServerManifestScenarioTwo(sc.appConfig, configurationForm)
	clientContent := generateAndSaveClientScriptScenarioTwo(sc.appConfig, sc.userConfig, configurationForm)

	fullContents := []string{gatewayContent, gatewayManifestContent, serverContent, serverManifestContent, clientContent}
	saveScenarioTwoFullDocFile(
		outputMarkdownFile,
		configurationForm.ECGatewayName,
		configurationForm.ECServerName,
		configurationForm.LocalPort,
		fullContents,
		sc.appConfig)
	return nil
}

// ExtractAndMoveECBinaryScenarioTwo ...
func (sc *ScenarioTwoService) ExtractAndMoveECBinaryScenarioTwo(configurationForm ScenarioTwoConfigurationForm) {
	gatewayFilename := filepath.Join(sc.appConfig.Output.Root, sc.appConfig.ECSDK.Folder, sc.appConfig.ECSDK.OS.Linux)
	gatewayDestFolder := filepath.Join(sc.appConfig.Output.Root, sc.appConfig.Output.Gateway.Root)
	helpers.ExtractArchiveFileToFolder(gatewayFilename, gatewayDestFolder)

	serverFilename := filepath.Join(sc.appConfig.Output.Root, sc.appConfig.ECSDK.Folder, sc.appConfig.ECSDK.OS.Linux)
	serverDestFolder := filepath.Join(sc.appConfig.Output.Root, sc.appConfig.Output.Server.Root)
	helpers.ExtractArchiveFileToFolder(serverFilename, serverDestFolder)

	clientFilename := filepath.Join(sc.appConfig.Output.Root, sc.appConfig.ECSDK.Folder)
	if strings.Contains(configurationForm.OnPremiseOS, "linux") {
		clientFilename = filepath.Join(clientFilename, sc.appConfig.ECSDK.OS.Linux)
	} else if strings.Contains(configurationForm.OnPremiseOS, "arm") {
		clientFilename = filepath.Join(clientFilename, sc.appConfig.ECSDK.OS.ARM)
	} else if strings.Contains(configurationForm.OnPremiseOS, "darwin") {
		clientFilename = filepath.Join(clientFilename, sc.appConfig.ECSDK.OS.Darwin)
	} else if strings.Contains(configurationForm.OnPremiseOS, "windows") {
		clientFilename = filepath.Join(clientFilename, sc.appConfig.ECSDK.OS.Windows)
	}
	clientDestFolder := filepath.Join(sc.appConfig.Output.Root, sc.appConfig.Output.Client.Root)
	helpers.ExtractArchiveFileToFolder(clientFilename, clientDestFolder)
}

func generateAndSaveGatewayScriptScenarioTwo(config *helpers.AppSettings, configurationForm ScenarioTwoConfigurationForm) string {
	gatewayTemplate := filepath.Join(config.Internal.Root, config.Internal.Templates.GatewayScript)
	replaceMe := []string{
		"<ec-zone-id>",
		"<ec-cf-service-url>",
		"<ec-admin-token-from-vcap>",
		"<debug>",
	}
	replaceWith := []string{
		configurationForm.ECZoneID,
		configurationForm.ECServiceURI,
		configurationForm.ECAdminToken,
		configurationForm.IsDebug,
	}

	content := helpers.ReplaceAll(gatewayTemplate, replaceMe, replaceWith)
	saveTo := filepath.Join(config.Output.Root, config.Output.Gateway.Root, config.Output.Gateway.Script)
	if err := helpers.WriteStringToFile(saveTo, content, true); err != nil {
		panic(err)
	}
	return content
}

func generateAndSaveGatewayManifestScenarioTwo(config *helpers.AppSettings, configurationForm ScenarioTwoConfigurationForm) string {
	manifestTemplate := filepath.Join(config.Internal.Root, config.Internal.Templates.GatewayTmpl.Manifest)
	replaceMe := []string{"<ecagent_gateway_name>"}
	replaceWith := []string{configurationForm.ECGatewayName}

	content := helpers.ReplaceAll(manifestTemplate, replaceMe, replaceWith)
	saveTo := filepath.Join(config.Output.Root, config.Output.Gateway.Root, config.Output.Gateway.Manifest)
	if err := helpers.WriteStringToFile(saveTo, content, false); err != nil {
		panic(err)
	}
	return content
}

func generateAndSaveServerManifestScenarioTwo(config *helpers.AppSettings, configurationForm ScenarioTwoConfigurationForm) string {
	manifestTemplate := filepath.Join(config.Internal.Root, config.Internal.Templates.ServerTmpl.Manifest)
	replaceMe := []string{"<ecagent_server_name>"}
	replaceWith := []string{configurationForm.ECServerName}

	content := helpers.ReplaceAll(manifestTemplate, replaceMe, replaceWith)
	saveTo := filepath.Join(config.Output.Root, config.Output.Server.Root, config.Output.Server.Manifest)
	if err := helpers.WriteStringToFile(saveTo, content, false); err != nil {
		panic(err)
	}
	return content
}

func generateAndSaveServerScriptScenarioTwo(appConfig *helpers.AppSettings, userConfig *helpers.UserConfig, configurationForm ScenarioTwoConfigurationForm) string {
	serverTemplate, saveTo := handleServerScriptGenerationScenarioTwo(appConfig, configurationForm)
	replaceMe := []string{
		"<ecagent-os-sys>",
		"<server-id>",
		"<gateway-url>",
		"<resource-url>",
		"<resource-port>",
		"<uaa-client-id>",
		"<uaa-client-secret>",
		"<uaa-url>",
		"<ec-zone-id>",
		"<ec-cf-service-url>",
	}
	replaceWith := []string{
		"ecagent_linux_sys",
		configurationForm.ECIDS[0],
		"wss://" + configurationForm.ECGatewayName + userConfig.Predix.Domain,
		configurationForm.ResourceHost,
		configurationForm.ResourcePort,
		configurationForm.UAAClient,
		configurationForm.UAAClientSecret,
		configurationForm.UAAServiceURI,
		configurationForm.ECZoneID,
		configurationForm.ECServiceURI,
	}

	content := helpers.ReplaceAll(serverTemplate, replaceMe, replaceWith)
	if err := helpers.WriteStringToFile(saveTo, content, true); err != nil {
		panic(err)
	}
	return content
}

func handleServerScriptGenerationScenarioTwo(config *helpers.AppSettings, configurationForm ScenarioTwoConfigurationForm) (template, fileToSave string) {
	serverTemplate := filepath.Join(config.Internal.Root, config.Internal.Templates.ServerTmpl.ServerScriptScenarioTwo.Unix)
	saveTo := filepath.Join(config.Output.Root, config.Output.Server.Root, config.Output.Server.Script.Unix)

	fmt.Println(serverTemplate + " - " + saveTo)
	return serverTemplate, saveTo
}

func generateAndSaveClientScriptScenarioTwo(appConfig *helpers.AppSettings, userConfig *helpers.UserConfig, configurationForm ScenarioTwoConfigurationForm) string {
	clientTemplate, saveTo := handleClientScriptGenerationScenarioTwo(appConfig, configurationForm)
	replaceMe := []string{
		"<ecagent-os-sys>",
		"<client-id>",
		"<gateway-url>",
		"<server-id>",
		"<ec-zone-id>",
		"<uaa-url>",
		"<uaa-client-id>",
		"<uaa-client-secret>",
		"<local-port>",
		"<debug>",
		"<proxy>",
	}
	replaceWith := []string{
		configurationForm.OnPremiseOS,
		configurationForm.ECIDS[1],
		"wss://" + configurationForm.ECGatewayName + userConfig.Predix.Domain,
		configurationForm.ECIDS[0],
		configurationForm.ECZoneID,
		configurationForm.UAAServiceURI,
		configurationForm.UAAClient,
		configurationForm.UAAClientSecret,
		configurationForm.LocalPort,
		configurationForm.IsDebug,
		configurationForm.Proxy,
	}

	content := helpers.ReplaceAll(clientTemplate, replaceMe, replaceWith)
	if err := helpers.WriteStringToFile(saveTo, content, true); err != nil {
		panic(err)
	}
	return content
}

func handleClientScriptGenerationScenarioTwo(config *helpers.AppSettings, configurationForm ScenarioTwoConfigurationForm) (template, fileToSave string) {
	clientTemplate := "" //filepath.Join(config.Internal.Root, config.Internal.Templates.ClientTmpl.ClientScriptScenarioTwo.Unix)
	saveTo := ""
	if strings.Contains(configurationForm.OnPremiseOS, "windows") {
		clientTemplate = filepath.Join(config.Internal.Root, config.Internal.Templates.ClientTmpl.ClientScriptScenarioTwo.Windows)
		saveTo = filepath.Join(config.Output.Root, config.Output.Client.Root, config.Output.Client.Script.Windows)
	} else {
		clientTemplate = filepath.Join(config.Internal.Root, config.Internal.Templates.ClientTmpl.ClientScriptScenarioTwo.Unix)
		saveTo = filepath.Join(config.Output.Root, config.Output.Client.Root, config.Output.Client.Script.Unix)
	}
	//saveTo := filepath.Join(config.Output.Root, config.Output.Client.Root, config.Output.Client.Script.Unix)

	return clientTemplate, saveTo
}

func saveScenarioTwoFullDocFile(filename string, ecAgentGatewayName string, ecAgentServerName string, localPort string, contents []string, config *helpers.AppSettings) {
	fullDocTemplate, err := ioutil.ReadFile(filepath.Join(config.Internal.Root, config.Internal.Templates.FullDoc.ScenarioTwo))
	if err != nil {
		log.Fatal(err)
	}
	mdContent := string(fullDocTemplate)
	mdContent = strings.Replace(mdContent, "<gateway_script_content_here>", contents[0], -1)
	mdContent = strings.Replace(mdContent, "<gateway_manifest_content_here>", contents[1], -1)
	mdContent = strings.Replace(mdContent, "<server_script_content_here>", contents[2], -1)
	mdContent = strings.Replace(mdContent, "<server_manifest_content_here>", contents[3], -1)
	mdContent = strings.Replace(mdContent, "<client_script_content_here>", contents[4], -1)
	mdContent = strings.Replace(mdContent, "<ecagent_gateway_name>", ecAgentGatewayName, -1)
	mdContent = strings.Replace(mdContent, "<ecagent_server_name>", ecAgentServerName, -1)
	mdContent = strings.Replace(mdContent, "<local_port>", localPort, -1)

	if err := helpers.WriteStringToFile(filename, mdContent, true); err != nil {
		panic(err)
	}
}
