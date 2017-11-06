package services

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/indaco/predix-ec-configurator/helpers"
)

// ScenarioOneService ...
type ScenarioOneService struct {
	appConfig  *helpers.AppSettings
	userConfig *helpers.UserConfig
}

// NewScenarioOneService ...
func NewScenarioOneService(appConfig *helpers.AppSettings, userConfig *helpers.UserConfig) (*ScenarioOneService, error) {
	return &ScenarioOneService{
		appConfig:  appConfig,
		userConfig: userConfig,
	}, nil
}

// ScenarioOneConfigurationForm ...
type ScenarioOneConfigurationForm struct {
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
	// Resources
	OnPremiseOS     string
	ResourceHost    string
	ResourcePort    string
	IsDebug         string
	HealthcheckPort string
	// Networking
	Proxy string
}

// GenerateAndSaveAllScenarioOne ...
func (sc *ScenarioOneService) GenerateAndSaveAllScenarioOne(outputMarkdownFile string, configurationForm ScenarioOneConfigurationForm) error {
	gatewayContent := generateAndSaveGatewayScriptScenarioOne(sc.appConfig, configurationForm)
	gatewayManifestContent := generateAndSaveGatewayManifestScenarioOne(sc.appConfig, configurationForm)
	serverContent := generateAndSaveServerScriptScenarioOne(sc.appConfig, sc.userConfig, configurationForm)
	clientContent := generateAndSaveClientScriptScenarioOne(sc.appConfig, sc.userConfig, configurationForm)

	fullContents := []string{gatewayContent, gatewayManifestContent, serverContent, clientContent}
	saveScenarioOneFullDocFile(outputMarkdownFile, configurationForm.ECGatewayName, fullContents, sc.appConfig)
	return nil
}

// ExtractAndMoveECBinaryScenarioOne ...
func (sc *ScenarioOneService) ExtractAndMoveECBinaryScenarioOne(configurationForm ScenarioOneConfigurationForm) {
	gatewayFilename := filepath.Join(sc.appConfig.Output.Root, sc.appConfig.ECSDK.Folder, sc.appConfig.ECSDK.OS.Linux)
	gatewayDestFolder := filepath.Join(sc.appConfig.Output.Root, sc.appConfig.Output.Gateway.Root)
	helpers.ExtractArchiveFileToFolder(gatewayFilename, gatewayDestFolder)

	clientFilename := filepath.Join(sc.appConfig.Output.Root, sc.appConfig.ECSDK.Folder, sc.appConfig.ECSDK.OS.Linux)
	clientDestFolder := filepath.Join(sc.appConfig.Output.Root, sc.appConfig.Output.Client.Root)
	helpers.ExtractArchiveFileToFolder(clientFilename, clientDestFolder)

	serverFilename := filepath.Join(sc.appConfig.Output.Root, sc.appConfig.ECSDK.Folder)
	if strings.Contains(configurationForm.OnPremiseOS, "linux") {
		serverFilename = filepath.Join(serverFilename, sc.appConfig.ECSDK.OS.Linux)
	} else if strings.Contains(configurationForm.OnPremiseOS, "arm") {
		serverFilename = filepath.Join(serverFilename, sc.appConfig.ECSDK.OS.ARM)
	} else if strings.Contains(configurationForm.OnPremiseOS, "darwin") {
		serverFilename = filepath.Join(serverFilename, sc.appConfig.ECSDK.OS.Darwin)
	} else if strings.Contains(configurationForm.OnPremiseOS, "windows") {
		serverFilename = filepath.Join(serverFilename, sc.appConfig.ECSDK.OS.Windows)
	}
	serverDestFolder := filepath.Join(sc.appConfig.Output.Root, sc.appConfig.Output.Server.Root)
	helpers.ExtractArchiveFileToFolder(serverFilename, serverDestFolder)
}

func generateAndSaveGatewayScriptScenarioOne(config *helpers.AppSettings, configurationForm ScenarioOneConfigurationForm) string {
	gatewayTemplate := filepath.Join(config.Internal.Root, config.Internal.Templates.GatewayTmpl.GatewayScript)
	replaceMe := []string{
		"<ecagent-os-sys>",
		"<ec-zone-id>",
		"<ec-cf-service-url>",
		"<ec-admin-token-from-vcap>",
	}
	replaceWith := []string{
		"ecagent_linux_sys",
		configurationForm.ECZoneID,
		configurationForm.ECServiceURI,
		configurationForm.ECAdminToken,
	}

	content := helpers.ReplaceAll(gatewayTemplate, replaceMe, replaceWith)
	saveTo := filepath.Join(config.Output.Root, config.Output.Gateway.Root, config.Output.Gateway.Script)
	if err := helpers.WriteStringToFile(saveTo, content, true); err != nil {
		panic(err)
	}
	return content
}

func generateAndSaveGatewayManifestScenarioOne(config *helpers.AppSettings, configurationForm ScenarioOneConfigurationForm) string {
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

func generateAndSaveServerScriptScenarioOne(appConfig *helpers.AppSettings, userConfig *helpers.UserConfig, configurationForm ScenarioOneConfigurationForm) string {
	serverTemplate, saveTo := handleServerScriptGenerationScenarioOne(appConfig, configurationForm)
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
		"<debug>",
		"<proxy>",
	}
	replaceWith := []string{
		configurationForm.OnPremiseOS,
		configurationForm.ECIDS[0],
		"wss://" + configurationForm.ECGatewayName + userConfig.Predix.Domain,
		configurationForm.ResourceHost,
		configurationForm.ResourcePort,
		configurationForm.UAAClient,
		configurationForm.UAAClientSecret,
		configurationForm.UAAServiceURI,
		configurationForm.ECZoneID,
		configurationForm.ECServiceURI,
		configurationForm.IsDebug,
		configurationForm.Proxy,
	}

	content := helpers.ReplaceAll(serverTemplate, replaceMe, replaceWith)
	if err := helpers.WriteStringToFile(saveTo, content, true); err != nil {
		panic(err)
	}
	return content
}

func handleServerScriptGenerationScenarioOne(config *helpers.AppSettings, configurationForm ScenarioOneConfigurationForm) (template, fileToSave string) {
	serverTemplate := ""
	saveTo := ""

	if strings.Contains(configurationForm.OnPremiseOS, "windows") {
		serverTemplate = filepath.Join(config.Internal.Root, config.Internal.Templates.ServerTmpl.ServerScriptScenarioOne.Windows)
		saveTo = filepath.Join(config.Output.Root, config.Output.Server.Root, config.Output.Server.Script.Windows)
	} else {
		serverTemplate = filepath.Join(config.Internal.Root, config.Internal.Templates.ServerTmpl.ServerScriptScenarioOne.Unix)
		saveTo = filepath.Join(config.Output.Root, config.Output.Server.Root, config.Output.Server.Script.Unix)
	}

	return serverTemplate, saveTo
}

func generateAndSaveClientScriptScenarioOne(appConfig *helpers.AppSettings, userConfig *helpers.UserConfig, configurationForm ScenarioOneConfigurationForm) string {
	clientTemplate, saveTo := handleClientScriptGenerationScenarioOne(appConfig, configurationForm)
	replaceMe := []string{
		"<ecagent-os-sys>",
		"<client-id>",
		"<gateway-url>",
		"<server-id>",
		"<uaa-url>",
		"<uaa-client-id>",
		"<uaa-client-secret>",
		"<local-port>",
	}
	replaceWith := []string{
		"ecagent_linux_sys",
		configurationForm.ECIDS[1],
		"wss://" + configurationForm.ECGatewayName + userConfig.Predix.Domain,
		configurationForm.ECIDS[0],
		configurationForm.UAAServiceURI,
		configurationForm.UAAClient,
		configurationForm.UAAClientSecret,
		configurationForm.ResourcePort,
	}

	content := helpers.ReplaceAll(clientTemplate, replaceMe, replaceWith)
	if err := helpers.WriteStringToFile(saveTo, content, true); err != nil {
		panic(err)
	}
	return content
}

func handleClientScriptGenerationScenarioOne(config *helpers.AppSettings, configurationForm ScenarioOneConfigurationForm) (template, fileToSave string) {
	clientTemplate := filepath.Join(config.Internal.Root, config.Internal.Templates.ClientTmpl.ClientScriptScenarioOne.Unix)
	saveTo := filepath.Join(config.Output.Root, config.Output.Client.Root, config.Output.Client.Script.Unix)

	return clientTemplate, saveTo
}

func saveScenarioOneFullDocFile(filename string, ecAgentGatewayName string, contents []string, config *helpers.AppSettings) {
	fullDocTemplate, err := ioutil.ReadFile(filepath.Join(config.Internal.Root, config.Internal.Templates.FullDoc.ScenarioOne))
	if err != nil {
		log.Fatal(err)
	}
	mdContent := string(fullDocTemplate)
	mdContent = strings.Replace(mdContent, "<gateway_script_content_here>", contents[0], -1)
	mdContent = strings.Replace(mdContent, "<gateway_manifest_content_here>", contents[1], -1)
	mdContent = strings.Replace(mdContent, "<server_script_content_here>", contents[2], -1)
	mdContent = strings.Replace(mdContent, "<client_script_content_here>", contents[3], -1)
	mdContent = strings.Replace(mdContent, "<ecagent_gateway_name>", ecAgentGatewayName, -1)

	if err := helpers.WriteStringToFile(filename, mdContent, true); err != nil {
		panic(err)
	}
}
